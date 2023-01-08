package main

import (
	"flag"

	"github.com/DisposableChat/api-core"
	"github.com/joho/godotenv"
)

var (
	Production *bool
	Redis 	= RedisAPI{}
)

func main() {
	Production = flag.Bool("production", false, "Production mode")
	devEnvPath := flag.String("devenv", "", "Development environment path")
	flag.Parse()

	if !*Production && *devEnvPath != "" {
		err := godotenv.Load(*devEnvPath)
		if err != nil {
			panic(err)
		}
	} else if !*Production && *devEnvPath == "" {
		panic("No development environment path provided")
	} else if *Production && *devEnvPath != "" {
		panic("Production mode and development environment path provided")
	} 
	
	err := core.GenerateLocalHashPassword()
	if err != nil {
		panic(err)
	}

	Redis.Init()
	server := Server{}
	server.Init()
}