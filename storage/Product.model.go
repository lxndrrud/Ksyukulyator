package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/webviewKsyukulyator/dto"
)

type ProductModel struct {
	db *sqlx.DB
}

func NewProductModel(db1 *sqlx.DB) *ProductModel {
	return &ProductModel{
		db: db1,
	}
}

func (s *ProductModel) GetAll() ([]dto.ProductFull, error) {
	products := make([]dto.ProductFull, 0)

	err := s.db.Select(&products, `
		SELECT 
			p.*, 
			CASE 
				WHEN c.category_title IS NULL THEN "Без категории"  
				WHEN c.category_title IS NOT NULL THEN c.category_title
			END as category_title 
		FROM products p
		LEFT JOIN categories AS c ON c.category_id = p.product_id_category
	`)
	if err != nil {
		return []dto.ProductFull{}, nil
	}
	return products, nil
}

func (s *ProductModel) AddProduct(product dto.Product) error {
	fmt.Println(product)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		INSERT INTO products(product_title, product_cost, product_id_category) VALUES(?, ?, ?)
	`, product.Title, product.Cost, product.IdCategory)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductModel) DeleteProduct(id int64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		DELETE FROM products WHERE products.product_id = ?
	`, id)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductModel) SearchForProducts(productTitle string, categoryTitle string) ([]dto.Product, error) {
	if categoryTitle != "" {
		products := make([]dto.Product, 0)

		err := s.db.Select(&products,
			`
			SELECT p.* 
			FROM products p
			JOIN categories AS c ON c.category_id = p.product_id_category
			WHERE p.products_title = ? AND c.category_title = ?
		`, productTitle, categoryTitle)
		if err != nil {
			return []dto.Product{}, err
		}
		return products, nil
	}
	products := make([]dto.Product, 0)

	err := s.db.Select(&products,
		`
		SELECT * FROM products WHERE products.products_title = ?
	`, productTitle)
	if err != nil {
		return []dto.Product{}, err
	}
	return products, nil
}
