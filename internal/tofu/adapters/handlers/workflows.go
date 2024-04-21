package handlers

import tofu_ports "github.com/erkylima/automates/internal/tofu/core/ports"

type Workflows struct {
	service *tofu_ports.TofuService
}

func NewWorkflowsService(svc *tofu_ports.TofuService) *Workflows {
	return &Workflows{
		service: svc,
	}
}
