package services

import (
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
)

type DockerComposeService struct {
}

func NewDockerComposeService() *DockerComposeService {
	return &DockerComposeService{}
}

func (s *DockerComposeService) Generate(packages []string) (*domain.DockerCompose, error) {
	services := []domain.Service{}
	for _, p := range packages {
		if p == "nginx" {
			service := NewNginxService("reverseProxy", "1.25")
			services = append(services, *service)
		} else {
			services = append(services, domain.Service{
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

	return &domain.DockerCompose{
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

func NewNginxService(name, version string) *domain.Service {
	service := &domain.Service{
		ContainerName: name,
		Build:         "",
		Image:         "nginx:" + version,
		Ports:         []string{"80:80"},
		Volumes:       []string{"/var/run/docker.sock:/tmp/docker.sock:ro"},
		DependsOn:     []string{""},
		Environment:   []string{""},
	}
	return service
}
