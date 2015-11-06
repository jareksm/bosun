package collect

import (
	"bytes"
	"encoding/binary"
	"net"
	"syscall"
	"time"

	"bosun.org/_third_party/github.com/garyburd/redigo/redis"
	"bosun.org/slog"
)

/*
 Listen on the specified udp port for events.
 This provides long term aggrigation for sparse events.
 wire format: opcode(1 byte) | data

Opcodes:
 1: increment - increments a redis counter for the specified metric/tag set
     data: count(4 bytes signed int) | metric:tag1=foo,tag2=bar
*/
func ListenUdp(port int, redisHost string, redisBucket int) error {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}
	pool := newRedisPool(redisHost, redisBucket)
	for {
		buff := make([]byte, 1024)
		oob := make([]byte, 1024)
		n, _, flags, addr, err := conn.ReadMsgUDP(buff, oob)
		if err != nil {
			slog.Error(err)
			continue
		}
		if flags&syscall.MSG_TRUNC != 0 {
			slog.Errorf("Too large a udp packet received from: %s. Skipping.", addr.String())
			continue
		}
		go func(data []byte, addr string) {
			c := pool.Get()
			defer c.Close()
			if len(data) == 0 {
				slog.Errorf("Empty packet received from %s.", addr)
			}
			switch data[0] {
			case 1:
				incrementRedisCounter(data[1:], addr, c)
			default:
				slog.Errorf("Unknown opcode %d from %s.", data[0], addr)
			}
		}(buff[:n], addr.String())
	}
}

func incrementRedisCounter(data []byte, addr string, conn redis.Conn) {
	if len(data) < 5 {
		slog.Errorf("Insufficient data for increment from %s.", addr)
		return
	}
	r := bytes.NewReader(data)
	var i int32
	err := binary.Read(r, binary.BigEndian, &i)
	if err != nil {
		slog.Error(err)
		return
	}
	mts := string(data[4:])
	if _, err = conn.Do("HINCRBY", RedisCountersKey, mts, i); err != nil {
		slog.Errorf("Error incrementing counter %s by %d. From %s. %s", mts, i, addr, err)
	}
}

const RedisCountersKey = "scollectorCounters"

func newRedisPool(server string, database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(database))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("CLIENT", "SETNAME", metricRoot+"UDP"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
