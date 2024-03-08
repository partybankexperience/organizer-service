package repositories

type Repository[T, U any] interface {
	Save(t *T) *T
	FindById(id U) *T
	FindAll() []*T
	DeleteById(id U)
}

type RepositoryImpl[T, U any] struct {
}

func (r *RepositoryImpl[T, U]) Save(t *T) *T {

	return nil
}

func (r *RepositoryImpl[T, U]) FindById(id U) *T {
	return nil
}

func (r *RepositoryImpl[T, U]) FindAll() []*T {
	return nil
}

func (r *RepositoryImpl[T, U]) DeleteById(id U) {

}
