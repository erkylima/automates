package ports

import "github.com/qbem-repos/dockerizing-service/internal/core/domain"

type DockerComposeSerializer interface {
	Decode(input []byte) (*domain.DockerCompose, error)
	Encode(input *domain.DockerCompose) ([]byte, error)
}
