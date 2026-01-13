package repo

type IBaseRepo[T any] interface {
	// GetAll retrieves all entities of type T from the data source.
	GetAll() ([]T, error)

	// SaveAll saves multiple entities of type T to the data source.
	SaveAll(entities []T) error
}
