package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
	"github.com/qbem-repos/dockerizing-service/internal/core/ports"
)

type DevelopersHandler struct {
	svc ports.DockerComposeService
}

func NewDevelopersHandler(dockerComposeService ports.DockerComposeService) *DevelopersHandler {
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
	docker_compose, err := h.svc.Generate([]string{"nginx", "web"})
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.YAML(http.StatusOK, docker_compose)
}
