package dockerizer

import "errors"

var (
	ErrDockerConfigNotEnough = errors.New("Docker Config Not Enough")
	ErrDockerConfigInvalid   = errors.New("Docker Config Invalid")
)

type dockerComposeService struct {
}

func NewDockerComposeService() *dockerComposeService {
	return &dockerComposeService{}
}

func (s *dockerComposeService) Generate(packages []string) (*DockerCompose, error) {
	services := []Service{}
	for _, p := range packages {
		if p == "nginx" {
			service := NewNginxService("reverseProxy", "1.25")
			services = append(services, *service)
		} else {
			services = append(services, Service{
				ContainerName: "name",
				Build:         ".",
				Ports:         []string{"80:80"},
				Volumes:       []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
				DependsOn:     []string{""},
				Environment:   []string{""},
			},
			)
		}
	}

	return &DockerCompose{
		Services: services,
		Version:  "3.7",
		Networks: []string{
			"das",
		},
		Volumes: []string{
			"das",
		},
	}, nil
}
