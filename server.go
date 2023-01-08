package main

import (
	"os"
	"strconv"

	"github.com/DisposableChat/api-core"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Core core.Server
	Worker fiber.Router
}

// /etc/ssl/csr.perm
func (s *Server) Init() {
	addr := os.Getenv("API_CONTAINER_LOCAL_ADDRESS")
	portStr := os.Getenv("API_USERS_CONTAINER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}

	s.Core = core.Server{
		Address: addr,
		Port:    int16(port),
	}

	s.Core.App = s.Core.New()
	router := Router{
		router: s.Core.App.Group("/api/auth"),
	}
	router.Configure()

	s.Listen()
}

func (s *Server) Listen() {
	err := s.Core.Listen()
	if err != nil {
		panic(err)
	}
}