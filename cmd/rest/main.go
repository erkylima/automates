package main

import (
	tofu_domains "github.com/erkylima/automates/internal/tofu/core/domains"
	tofu_services "github.com/erkylima/automates/internal/tofu/core/services"
	"github.com/erkylima/automates/pkg/database"
)

// repo <- service -> serializer  -> http

func main() {
	base, err := database.NewMongoConnection[tofu_domains.TofuWorkflow]("mongodb://admin:password@localhost:27017/admin?retryWrites=true&loadBalanced=false&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1", "tofu", "workflows")
	if err != nil {
		panic(err)
	}

	tofu := tofu_services.NewTofuService(base)
	tofu.InitTofu()
}
