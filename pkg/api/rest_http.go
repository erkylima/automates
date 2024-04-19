package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitDockerComposeRouter() {
	router := gin.Default()
	pprof.Register(router)
	// v1 := router.Group("/v1")

	// developersHandler := handlers.NewDevelopersHandler(svc)
	// v1.GET("/developers", developersHandler.Generate)

	// err := router.Run(":4242")
	// if err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
}
