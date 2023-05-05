package product_pg

import (
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
	"challange10-dts/repository/product_repository"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	getProductByIdQuery = `
		SELECT id, title, userId, description, createdAt, updatedAt from products
		WHERE id = $1;
	`

	updateProductByIdQuery = `
		UPDATE products
		SET title = $2,
		description = $3
		WHERE id = $1;
	`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) product_repository.ProductRepository {
	return &productPG{
		db: db,
	}
}

func (m *productPG) GetProducts() ([]*entity.Product, errs.MessageErr) {
	return nil, nil
}

func (m *productPG) UpdateProductById(payload entity.Product) errs.MessageErr {
	_, err := m.db.Exec(updateProductByIdQuery, payload.Id, payload.Title, payload.Description)

	if err != nil {
		// log.Println(err)
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	row := m.db.QueryRow(getProductByIdQuery, productId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.UserId, &product.Description, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found")
		}

		log.Println(err)

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}

func (m *productPG) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	createProductQuery := `
		INSERT INTO "products"
		(
			title,
			description,
			userId
		)
		VALUES($1, $2, $3)
		RETURNING id, title, description, userId;
	`
	row := m.db.QueryRow(createProductQuery, productPayload.Title, productPayload.Description, productPayload.UserId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.Description, &product.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil

}
