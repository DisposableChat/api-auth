package main

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/DisposableChat/api-core"
)

type RedisAPI struct {
	Core core.Redis
	Client redis.Client
}

func (r *RedisAPI) Init() (error) {
	host := os.Getenv("REDIS_HOST")
	portstr := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	port, err := strconv.Atoi(portstr)
	if err != nil {
		return err
	}

	r.Core = core.Redis{
		Host:     host,
		Port:     int16(port),
		Password: password,
	}

	r.Client = *r.Core.Init()

	return nil
}