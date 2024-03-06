package domain

type Stack struct {
	Technology []Technology
}

type Technology struct {
	ContainerName string
	Image         string
	ImageVersion  string
	Ports         []string
	Volumes       []string
	DependsOn     []string
	Environment   []string
}
