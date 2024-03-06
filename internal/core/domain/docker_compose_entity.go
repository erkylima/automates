package domain

type DockerCompose struct {
	Version  string    `json:"version,omitempty" yaml:"version"`
	Services []Service `json:"services" yaml:"services"`
	Networks []string  `json:"networks,omitempty" yaml:"networks"`
	Volumes  []string  `json:"volumes,omitempty" yaml:"volumes"`
}

type Service struct {
	ContainerName string   `json:"container_name" yaml:"container_name"`
	Build         string   `json:"build" yaml:"build "`
	Volumes       []string `json:"volumes" yaml:"volumes"`
	Ports         []string `yaml:"ports"`
	DependsOn     []string `yaml:"depends_on"`
	Environment   []string `yaml:"environment"`
	Image         string   `yaml:"image"`
}
