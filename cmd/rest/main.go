package main

import (
	"github.com/qbem-repos/dockerizing-service/internal/core/services"
	"github.com/qbem-repos/dockerizing-service/internal/core/services/serializers"
	"github.com/qbem-repos/dockerizing-service/pkg/api"
)

// repo <- service -> serializer  -> http

func main() {
	serializer := serializers.NewDockerComposeYaml()
	service := services.NewDockerComposeService(serializer)
	api.InitDockerComposeRouter(service)
}
