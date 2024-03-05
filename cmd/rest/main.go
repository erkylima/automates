package main

import (
	"fmt"

	"github.com/qbem-repos/dockerizing-service/internal/dockerizer"
	"github.com/qbem-repos/dockerizing-service/internal/dockerizer/serializers"
)

// repo <- service -> serializer  -> http

func main() {
	service := dockerizer.NewDockerComposeService()
	dockercompose, err := service.Generate([]string{"nginx", "web"})
	if err != nil {
		fmt.Println(err)
		return
	}
	serializer := serializers.DockerCompose{}
	encoded, err := serializer.Encode(dockercompose)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(encoded))
}
