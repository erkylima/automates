package serializers

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
)

type dockerComposeJson struct{}

func NewDockerComposeJson() *dockerComposeJson {
	return &dockerComposeJson{}
}

func (d *dockerComposeJson) Decode(input []byte) (*domain.DockerCompose, error) {
	dockercompose := &domain.DockerCompose{}
	if err := json.Unmarshal(input, &dockercompose); err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.decode")
	}
	return dockercompose, nil
}

func (d *dockerComposeJson) Encode(input *domain.DockerCompose) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.encode")
	}
	return rawMsg, nil
}
