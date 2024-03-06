package api

import (
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/dockerizing-service/internal/adapters/handlers"
	"github.com/qbem-repos/dockerizing-service/internal/core/ports"
)

func InitDockerComposeRouter(svc ports.DockerComposeService) {
	router := gin.Default()
	pprof.Register(router)
	v1 := router.Group("/v1")

	developersHandler := handlers.NewDevelopersHandler(svc)
	v1.POST("/developers", developersHandler.Generate)

	err := router.Run(":4242")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
