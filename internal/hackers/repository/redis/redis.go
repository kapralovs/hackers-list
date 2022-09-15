package redis

import "github.com/go-redis/redis/v9"

type config struct {
	client *redis.Client
}

func NewConfig(client *redis.Client) *config {
	return &config{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

return &config{
	client:redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}),