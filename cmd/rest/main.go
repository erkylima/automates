package main

import (
	"github.com/qbem-repos/dockerizing-service/internal/core/services"
	"github.com/qbem-repos/dockerizing-service/pkg/api"
)

// repo <- service -> serializer  -> http

func main() {
	service := services.NewDockerComposeService()
	api.InitDockerComposeRouter(service)
}
