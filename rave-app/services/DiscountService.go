package services

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
)

type DiscountService interface {
	CreateDiscount(request *request.CreateDiscountRequest) (*response.CreateDiscountResponse, error)
}

type raveDiscountService struct {
}

func NewDiscountService() DiscountService {
	return nil
}

func (raveDiscountService *raveDiscountService) CreateDiscount(request *request.CreateDiscountRequest) (*response.CreateDiscountResponse, error) {
	return nil, nil
}
