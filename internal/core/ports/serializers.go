package ports

import "github.com/qbem-repos/dockerizing-service/internal/core/domain"

type DockerComposeSerializer interface {
	Encode(inpuy []byte) (*domain.DockerCompose, error)
	Decode(input *domain.DockerCompose) ([]byte, error)
}
