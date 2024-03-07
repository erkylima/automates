package domain

type DockerCompose struct {
	Version  string    `json:"version,omitempty" yaml:"version"`
	Services []Service `json:"services" yaml:"services"`
	Networks []string  `json:"networks,omitempty" yaml:"networks"`
	Volumes  []string  `json:"volumes,omitempty" yaml:"volumes"`
}

type Service struct {
	ContainerName string   `json:"container_name" yaml:"container_name"`
	Build         Build    `json:"build,omitempty" yaml:"build,omitempty"`
	Volumes       []string `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	Ports         []string `json:"ports,omitempty" yaml:"ports"`
	DependsOn     []string `json:"dependsOn,omitempty" yaml:"depends_on,omitempty"`
	Environment   []string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Image         string   `json:"image,omitempty" yaml:"image,omitempty"`
}

type Build struct {
	Context    string `json:"context" yaml:"context"`
	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
}
