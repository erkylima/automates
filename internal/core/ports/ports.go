package ports

import "github.com/qbem-repos/dockerizing-service/internal/core/domain"

type ConfigurationService interface {
	Generate(stack domain.Stack) (*[]byte, error)
	Serializer() DockerComposeSerializer
}
