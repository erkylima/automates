package dockerizer

type DockerizerService interface {
	Generate(packages []string) (*DockerCompose, error)
}

func NewNginxService(name, version string) *Service {
	service := &Service{
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
