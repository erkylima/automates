package dockerizer

type DockerCompose struct {
	Version  string    `yaml:"version"`
	Services []Service `yaml:"services"`
	Networks []string  `yaml:"networks"`
	Volumes  []string  `yaml:"volumes"`
}

type Service struct {
	ContainerName string   `yaml:"container_name"`
	Build         string   `yaml:"build"`
	Volumes       []string `yaml:"volumes"`
	Ports         []string `yaml:"ports"`
	DependsOn     []string `yaml:"depends_on"`
	Environment   []string `yaml:"environment"`
	Image         string   `yaml:"image"`
}
