package collectors

import (
	"fmt"
	"strconv"
	"strings"

	"bosun.org/_third_party/github.com/garyburd/redigo/redis"
	"bosun.org/collect"
	"bosun.org/metadata"
	"bosun.org/opentsdb"
	"bosun.org/slog"
)

func init() {
	collectors = append(collectors,
		&IntervalCollector{
			name: "redisCounters",
			F: func() (opentsdb.MultiDataPoint, error) {
				return c_redis_counters("localhost:6379", 2)
			},
		})
}

func c_redis_counters(server string, bucket int) (opentsdb.MultiDataPoint, error) {
	var md opentsdb.MultiDataPoint
	conn, err := redis.Dial("tcp", server, redis.DialDatabase(bucket))
	if err != nil {
		return md, err
	}
	defer conn.Close()
	if _, err := conn.Do("CLIENT", "SETNAME", "scollector"); err != nil {
		return md, err
	}
	cursor := 0
	for {
		vals, err := redis.Values(conn.Do("HSCAN", collect.RedisCountersKey, cursor))
		if err != nil {
			return md, err
		}
		if len(vals) != 2 {
			return md, fmt.Errorf("Unexpected number of values")
		}
		cursor, err = redis.Int(vals[0], nil)
		if err != nil {
			return md, err
		}
		pairs, err := redis.StringMap(vals[1], nil)
		if err != nil {
			return md, err
		}
		for mts, val := range pairs {
			parts := strings.Split(mts, ":")
			if len(parts) != 2 {
				slog.Errorf("Invalid metric tag set counter: %s", mts)
				continue
			}
			metric := parts[0]
			tags, err := opentsdb.ParseTags(parts[1])
			if err != nil {
				slog.Errorf("Invalid tags: %s", parts[1])
				continue
			}
			v, err := strconv.Atoi(val)
			if err != nil {
				slog.Errorf("Invalid counter value: %s", val)
				continue
			}
			Add(&md, metric, v, tags, metadata.Counter, metadata.Count, "")
		}
		if cursor == 0 {
			break
		}
	}
	return md, nil
}
