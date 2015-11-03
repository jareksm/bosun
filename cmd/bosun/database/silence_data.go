package database

import (
	"encoding/json"
	"time"

	"bosun.org/_third_party/github.com/garyburd/redigo/redis"
	"bosun.org/collect"
	"bosun.org/models"
	"bosun.org/opentsdb"
)

/*

Silences : hash of Id - json of silence. Id is sha of fields

SilencesByEnd : zlist of end-time to id. 

Easy to find active. Find all with end time in future, and filter to those with start time in the past.

*/

const (
	silenceHash = "Silences"
	silenceIdx  = "SilencesByEnd"
)

type SilenceDataAccess interface {
	GetActiveSilences() ([]*models.Silence, error)
	AddSilence(*models.Silence) error
	DeleteSilence(id string) error

	ListSilences(perPage, page int) (map[string]*models.Silence, error)
}

func (d *dataAccess) Silence() SilenceDataAccess {
	return d
}

func (d *dataAccess) GetActiveSilences() ([]*models.Silence, error) {
	defer collect.StartTimer("redis", opentsdb.TagSet{"op": "GetActiveSilences"})()
	conn := d.GetConnection()
	defer conn.Close()

	vals, err := redis.Strings(conn.Do("ZRANGEBYSCORE", silenceIdx, time.Now().UTC().Unix(), "+inf"))
	if err != nil {
		return nil, err
	}

	silences, err := getSilences(vals, conn)
	if err != nil {
		return nil, err
	}
	s2 := make([]*models.Silence, 0, len(silences))
	now := time.Now()
	for _, s := range s2 {
		if s.Start.After(now) {
			continue
		}
		s2 = append(s2, s)
	}
	return s2, nil
}

func getSilences(ids []string, conn redis.Conn) ([]*models.Silence, error) {
	args := make([]interface{}, len(ids)+1)
	args[0] = silenceHash
	for i := range ids {
		args[i+1] = ids[i]
	}
	jsons, err := redis.Strings(conn.Do("HMGET", args...))
	if err != nil {
		return nil, err
	}

	silences := make([]*models.Silence, len(jsons))
	for _, j := range jsons {
		s := &models.Silence{}
		if err := json.Unmarshal([]byte(j), s); err != nil {
			return nil, err
		}
		silences = append(silences, s)
	}
	return silences, nil
}

func (d *dataAccess) AddSilence(s *models.Silence) error {
	defer collect.StartTimer("redis", opentsdb.TagSet{"op": "AddSilence"})()
	conn := d.GetConnection()
	defer conn.Close()

	if _, err := conn.Do("ZADD", silenceIdx, s.Start.UTC().Unix(), s.ID()); err != nil {
		return err
	}
	dat, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSET", silenceHash, s.ID(), dat)
	return err
}

func (d *dataAccess) DeleteSilence(id string) error {
	defer collect.StartTimer("redis", opentsdb.TagSet{"op": "DeleteSilence"})()
	conn := d.GetConnection()
	defer conn.Close()

	if _, err := conn.Do("ZREM", silenceIdx, id); err != nil {
		return err
	}
	if _, err := conn.Do("HDEL", silenceHash, id); err != nil {
		return err
	}
	return nil
}

func (d *dataAccess) ListSilences(perPage, page int) (map[string]*models.Silence, error) {
	defer collect.StartTimer("redis", opentsdb.TagSet{"op": "ListSilences"})()
	conn := d.GetConnection()
	defer conn.Close()

	if page < 0 {
		page = 0
	}
	if perPage < 1 {
		perPage = 20
	}
	start := page * perPage
	end := start + perPage

	ids, err := redis.Strings(conn.Do("ZRANGE", silenceIdx, start, end))
	if err != nil {
		return nil, err
	}
	silences, err := getSilences(ids, conn)
	if err != nil {
		return nil, err
	}
	m := make(map[string]*models.Silence, len(silences))
	for _, s := range silences {
		m[s.ID()] = s
	}
	return m, nil
}
