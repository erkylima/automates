package serializers

import (
	"github.com/pkg/errors"
	"github.com/qbem-repos/dockerizing-service/internal/dockerizer"
	"gopkg.in/yaml.v3"
)

type DockerCompose struct{}

func (d *DockerCompose) Decode(input []byte) (*dockerizer.DockerCompose, error) {
	dockercompose := &dockerizer.DockerCompose{}
	if err := yaml.Unmarshal(input, &dockercompose); err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.decode")
	}
	return dockercompose, nil
}

func (d *DockerCompose) Encode(input *dockerizer.DockerCompose) ([]byte, error) {
	rawMsg, err := yaml.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.dockercompose.encode")
	}
	return rawMsg, nil
}
