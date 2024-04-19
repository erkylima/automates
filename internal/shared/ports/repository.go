package shared_ports

type Repository[T any] interface {
	Create(entity *T) (string, error)
	Read(slug string, entity *T) error
	Update(slug string, entity *T) error
	Delete(slug string) error
	List(filter interface{}) ([]T, error)
}
