package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

type SerializableStore interface {
	Get(string) (*Session, error)
	Set(string, Session) error
	Delete(string) error
}

type Session struct {
	Username string
	Expiry   time.Time
}

var ctx = context.Background()

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func NewRedisStore() (SerializableStore, error) {
	// connStr := os.Getenv("REDIS_CONN")
	pass := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     "gouache-cache:6379",
		Password: pass,
		DB:       0, // use default DB
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisStore{client: client}, nil
}

func (r *RedisStore) Delete(sid string) error {
	if _, err := r.client.Del(ctx, sid).Result(); err != nil {
		return err
	}

	return nil
}

func (r *RedisStore) Get(sid string) (*Session, error) {
	if v, err := r.client.Get(ctx, sid).Bytes(); err != nil {
		return nil, err
	} else {
		session := &Session{}

		if err := json.Unmarshal(v, session); err != nil {
			return nil, err
		}

		return session, nil
	}
}

func (r *RedisStore) Set(sid string, session Session) error {
	v, err := json.Marshal(session)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, sid, v, time.Until(session.Expiry)).Err(); err != nil {
		return err
	}

	return nil
}
