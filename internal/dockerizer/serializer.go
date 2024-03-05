package dockerizer

type DockerizerSerializer interface {
	Encode(inpuy []byte) (*DockerCompose, error)
	Decode(input *DockerCompose) ([]byte, error)
}
