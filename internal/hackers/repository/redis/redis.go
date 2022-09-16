package redis

import (
	"github.com/go-redis/redis/v9"
	"github.com/kapralovs/hackers-list/internal/models"
)

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

func GetHackersList(rc *redis.Client) ([]*models.Hacker, error) {
	hackers := []*models.Hacker{}
	// rc.ZRange(context.Background(),"hackers")
	return hackers, nil
}
