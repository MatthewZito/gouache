package repositories

import (
	"context"
	"os"

	"github.com/exbotanical/gouache/utils"
	"github.com/go-redis/redis/v9"
)

// SessionRepository holds a redis client connection object and context.
type SessionRepository struct {
	Client *redis.Client
	Ctx    context.Context
}

// NewSessionRepository initializes a new `SessionRepository`.
func NewSessionRepository() (*SessionRepository, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	ctx := context.Background()

	opts := &redis.Options{
		Addr: utils.ToEndpoint(host, port),
		DB:   0, // use default DB
	}

	if os.Getenv("LOCAL_MODE") == "" {
		pass := os.Getenv("REDIS_PASSWORD")
		opts.Password = pass
	}

	client := redis.NewClient(opts)

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &SessionRepository{Client: client, Ctx: ctx}, nil
}
