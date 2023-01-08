package main

import (
	"github.com/DisposableChat/api-core"
	"github.com/joho/godotenv"
)

var (
	Production = true
	Redis 	= RedisAPI{}
)

func main() {
	if !Production {
		err := godotenv.Load("/home/jean/Desktop/DisposableChat/api/api.dev.env")
		if err != nil {
			panic(err)
		}
	}

	err := core.GenerateLocalHashPassword()
	if err != nil {
		panic(err)
	}

	Redis.Init()
	server := Server{}
	server.Init()
}