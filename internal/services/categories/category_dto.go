package categories

import (
	"fmt"

	"log"

	"github.com/amorindev/backend-inventory/internal/db"
)


func GetCategoriesDto() ([]CategoryEntity, error) {
	var categories []CategoryEntity



	rows, err := db.DB.Query("SELECT cat_id, cat_name FROM tb_category")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ctg CategoryEntity
		err := rows.Scan(&ctg.CatID, &ctg.CatName)
		if err != nil {
			return nil, fmt.Errorf("category rows: %v", err)
		}
		categories = append(categories, ctg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("category rows error: %v", err)
	}
	return categories, nil
}

func CreateCategoryDto(c CategoryEntity) (CategoryEntity, error) {
	var categoryID int64

	query := `INSERT INTO tb_category (cat_name) VALUES ($1) RETURNING cat_id`

	err := db.DB.QueryRow(query, c.CatName).Scan(&categoryID)
	if err != nil {
	  return CategoryEntity{}, fmt.Errorf("create category failed %v", err)
	}
	c.CatID = categoryID
	return c, nil
}

