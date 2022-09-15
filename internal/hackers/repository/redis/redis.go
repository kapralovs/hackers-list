package redis

import "github.com/go-redis/redis/v9"

type config struct {
	client *redis.Client
}

func NewConfig() *config {
return client:
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
