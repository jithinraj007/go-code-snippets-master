package service

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/dto"
)

//go:generate mockgen -destination=../mocks/service/mockProductService.go -package=service github.com/ashishjuyal/banking/service ProductService
type ProductService interface {
	GetAllProduct(string) ([]dto.ProductResponse, *errs.AppError)
	GetProduct(string) (*dto.ProductResponse, *errs.AppError)
}

type DefaultProductService struct {
	repo domain.ProductRepository
}

func (s DefaultProductService) GetAllProduct(discontinued string) ([]dto.ProductResponse, *errs.AppError) {

	queryParameter := 0

	if discontinued == "yes" {
		queryParameter = 1
	} else if discontinued == "no" {
		queryParameter = 0
	} else {
		queryParameter = 0
	}
	products, err := s.repo.FindAll(queryParameter)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range products {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultProductService) GetProduct(id string) (*dto.ProductResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewProductService(repository domain.ProductRepository) DefaultProductService {
	return DefaultProductService{repository}
}
