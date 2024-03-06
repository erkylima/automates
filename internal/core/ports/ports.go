package ports

import "github.com/qbem-repos/dockerizing-service/internal/core/domain"

type DockerComposeService interface {
	Generate(packages []string) (*domain.DockerCompose, error)
}
