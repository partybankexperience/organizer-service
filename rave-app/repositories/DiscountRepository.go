package repositories

import "github.com/djfemz/rave/rave-app/models"

type DiscountRepository interface {
	crudRepository[models.Discount, uint64]
}

type discountRepository struct {
	*repositoryImpl[models.Discount, uint64]
}

func NewDiscountRepository() DiscountRepository {
	return &discountRepository{
		&repositoryImpl[models.Discount, uint64]{},
	}
}
