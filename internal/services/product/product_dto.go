package product

import (
	//"database/sql"
	"fmt"
	"log"

	"github.com/amorindev/backend-inventory/internal/db"
)

func GetProducts() ([]Product, error) {
	var products []Product

	query := `SELECT prod_id, prod_name, prod_desc, prod_discount, prod_price, prod_stk, cat_id FROM tb_product`
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.ProductName, &p.ProductDescription, &p.ProductDiscount, &p.ProductPrice, &p.ProductStk, &p.CatID)

		if err != nil {
			return nil, fmt.Errorf("get Product %v", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get Product Err: %v", err)
	}

	return products, nil
}

func CreateProduct(p Product) (Product, error) {
	var productID int64

	query := `INSERT INTO tb_product (prod_name, prod_desc, prod_discount, prod_price, prod_stk,cat_id)
				VALUES ($1,$2,$3,$4,$5,$6) RETURNING prod_id`

	err := db.DB.QueryRow(query, p.ProductName, p.ProductDescription, p.ProductDiscount, p.ProductPrice, p.ProductStk, p.CatID).Scan(&productID)

	if err != nil {
		return Product{}, fmt.Errorf("created Product: %v", err)
	}

	p.ID = productID

	return p, nil
}

func UpdateProduct(id int64, p Product) error {

	query := `UPDATE tb_product SET prod_name = $1, prod_desc = $2, prod_discount = $3, prod_price =$4, prod_stk = $5, cat_id = $6 WHERE prod_id = $7`

	result, err := db.DB.Exec(query, p.ProductName, p.ProductDescription, p.ProductDiscount, p.ProductPrice, p.ProductStk, p.CatID, id)

	if err != nil {
		return fmt.Errorf("update Product Err: %v", err)
	}

	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updated Product Err %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func DeleteProduct(productID int64) error {
	query := `DELETE FROM tb_product WHERE prod_id = $1`

	res, err := db.DB.Exec(query, productID)
	if err != nil {
		return fmt.Errorf("delete product: %v", err)

	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected no checking: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product not found with id %d", productID)
	}
	return nil
}
