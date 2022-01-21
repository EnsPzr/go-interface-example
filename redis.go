package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisCache() (Cache, error) {
	dbAddr := os.Getenv("DB_ADDR")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbNu, _ := strconv.Atoi(os.Getenv("DB_NUMBER"))
	if dbAddr == "" || dbPassword == "" || dbNu == 0 {
		return nil, errors.New("Configlerde eksikler mevcut")
	}
	return &RedisCache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     dbAddr,
			Password: dbPassword,
			DB:       dbNu,
		}),
		ctx: context.Background(),
	}, nil
}

func (r *RedisCache) Name() string {
	return "Redis"
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	return val, err
}

func (r *RedisCache) Set(key string, value interface{}) error {
	err := r.rdb.Set(r.ctx, key, value, time.Second*60).Err()
	return err
}
