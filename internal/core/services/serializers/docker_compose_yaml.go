package serializers

import (
	"github.com/pkg/errors"
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
	"gopkg.in/yaml.v3"
)

type dockerComposeYaml struct{}

func NewDockerComposeYaml() *dockerComposeJson {
	return &dockerComposeJson{}
}

func (d *dockerComposeYaml) Decode(input []byte) (*domain.DockerCompose, error) {
	dockercompose := &domain.DockerCompose{}
	if err := yaml.Unmarshal(input, &dockercompose); err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.decode")
	}
	return dockercompose, nil
}

func (d *dockerComposeYaml) Encode(input *domain.DockerCompose) ([]byte, error) {
	rawMsg, err := yaml.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.encode")
	}
	return rawMsg, nil
}
