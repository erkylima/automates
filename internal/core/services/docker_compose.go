package services

import (
	"github.com/qbem-repos/dockerizing-service/internal/core/domain"
	"github.com/qbem-repos/dockerizing-service/internal/core/ports"
)

type DockerComposeService struct {
	serializer ports.DockerComposeSerializer
}

func NewDockerComposeService(serializer ports.DockerComposeSerializer) *DockerComposeService {
	return &DockerComposeService{serializer: serializer}
}

func (s *DockerComposeService) Generate(stack domain.Stack) (*[]byte, error) {
	services := []domain.Service{}
	for _, p := range stack.Technology {
		service := NewDockerService(p, domain.Build{
			Context:    "../",
			Dockerfile: ".devops/Dockerfile",
		})
		services = append(services, *service)
	}
	dockerCompose := &domain.DockerCompose{
		Services: services,
		Version:  "3.7",
		Networks: []string{
			"das",
		},
		Volumes: []string{
			"das",
		},
	}

	encoded, err := s.Serializer().Encode(dockerCompose)
	if err != nil {
		return nil, err
	}
	return &encoded, nil
}

func (s *DockerComposeService) Serializer() ports.DockerComposeSerializer {
	return s.serializer
}

func NewDockerService(technology domain.Technology, build domain.Build) *domain.Service {
	var service *domain.Service
	if build.Context == "" {
		service = &domain.Service{
			ContainerName: technology.ContainerName,
			Image:         technology.Image + ":" + technology.ImageVersion,
			Ports:         technology.Ports,
			Volumes:       technology.Volumes,
			DependsOn:     technology.DependsOn,
			Environment:   technology.Environment,
		}
	} else {
		service = &domain.Service{
			ContainerName: technology.ContainerName,
			Build:         build,
			Ports:         technology.Ports,
			Volumes:       technology.Volumes,
			DependsOn:     technology.DependsOn,
			Environment:   technology.Environment,
		}
	}
	return service
}
