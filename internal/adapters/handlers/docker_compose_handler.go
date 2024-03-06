package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
	"github.com/qbem-repos/dockerizing-service/internal/core/ports"
)

type DevelopersHandler struct {
	svc ports.ConfigurationService
}

func NewDevelopersHandler(dockerComposeService ports.ConfigurationService) *DevelopersHandler {
	return &DevelopersHandler{
		svc: dockerComposeService,
	}
}

func (h *DevelopersHandler) Generate(ctx *gin.Context) {
	dockerCompose := &domain.DockerCompose{}
	dockerCompose.Version = "3.8"

	if err := ctx.ShouldBindJSON(&dockerCompose); err != nil {
		HandleError(ctx, http.StatusBadRequest, errors.Wrap(err, err.Error()))
		return
	}
	technologies := []domain.Technology{
		{
			ContainerName: "PHP",
		},
		{
			ContainerName: "NodeJS",
		},
		{
			ContainerName: "Go",
		},
		{
			ContainerName: "Java",
		},
		{
			ContainerName: "Python",
		},
		{
			ContainerName: "C++",
		},
		{
			ContainerName: "C#",
		},
		{
			ContainerName: "C",
		},
		{
			ContainerName: "C++",
		},
		{
			ContainerName: "C#",
		},
		{
			ContainerName: "C",
		},
	}
	stack := &domain.Stack{
		Technology: technologies,
	}

	encoded_docker_compose, err := h.svc.Generate(*stack)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	decoded_docker_compose, err := h.svc.Serializer().Decode(*encoded_docker_compose)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.YAML(http.StatusOK, decoded_docker_compose)
}
