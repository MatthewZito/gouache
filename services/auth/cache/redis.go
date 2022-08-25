package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/exbotanical/gouache/utils"
	"github.com/go-redis/redis/v9"
)

// RedisStore holds a redis client connection object.
type RedisStore struct {
	client *redis.Client
}

// SessionManager defines the contract for a `RedisStore` instance.
type SessionManager interface {
	Get(string) (*Session, error)
	Set(string, Session) error
	Delete(string) error
}

// Session represents session metadata as stored in the redis session cache.
type Session struct {
	Username string
	Expiry   time.Time
}

var ctx = context.Background()

// IsExpired is a getter that determines whether the `Session` is expired.
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

// NewRedisStore initializes a new `SessionManager`.
func NewRedisStore() (SessionManager, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     utils.ToEndpoint(host, port),
		Password: pass,
		DB:       0, // use default DB
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisStore{client: client}, nil
}

// Delete - when provided a session ID `sid` - deletes the corresponding `Session` from the cache.
func (r *RedisStore) Delete(sid string) error {
	if _, err := r.client.Del(ctx, sid).Result(); err != nil {
		return err
	}

	return nil
}

// Get - when provided a session ID `sid` - retrieves the corresponding `Session` from the cache.
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

// Set - when provided a session ID `sid` and `Session` `session` - sets or updates the `Session` in the cache.
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
