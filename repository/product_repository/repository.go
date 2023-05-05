package product_repository

import (
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById(payload entity.Product) errs.MessageErr
	GetProducts() ([]*entity.Product, errs.MessageErr)
}
