package services

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
)

type DiscountService interface {
	CreateDiscount(request *request.CreateDiscountRequest) (*response.CreateDiscountResponse, error)
}

func NewDiscountService() DiscountService {
	return nil
}
