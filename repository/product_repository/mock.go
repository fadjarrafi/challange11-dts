package product_repository

import (
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
)

var (
	CreateProduct     func(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById    func(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById func(payload entity.Product) errs.MessageErr
	GetProducts       func() ([]*entity.Product, errs.MessageErr)
)

type productRepoMock struct{}

func NewProductRepoMock() ProductRepository {
	return &productRepoMock{}
}

func (m *productRepoMock) GetProducts() ([]*entity.Product, errs.MessageErr) {
	return GetProducts()
}

func (m *productRepoMock) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	return CreateProduct(productPayload)
}
func (m *productRepoMock) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	return GetProductById(productId)
}
func (m *productRepoMock) UpdateProductById(payload entity.Product) errs.MessageErr {
	return UpdateProductById(payload)
}
