package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/kapralovs/hackers-list/internal/models"
)

type RedisStorage struct {
	redisClient *redis.Client
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{
		redisClient: client,
	}
}

func (rs *RedisStorage) GetHackersList() []*models.Hacker {
	hackers := []*models.Hacker{}
	result, err := rs.redisClient.ZRangeWithScores(context.Background(), "hackers", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result {
		hacker := &models.Hacker{}
		hacker.Score = v.Score
		hacker.Name = v.Member.(string)
		hackers = append(hackers, hacker)
	}
	return hackers
}
