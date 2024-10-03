package categories

import (
	"fmt"

	"log"

	"github.com/amorindev/backend-inventory/db"
)


func GetCategoriesDto() ([]Category, error) {
	var categories []Category



	rows, err := db.DB.Query("SELECT cat_id, cat_name FROM tb_category")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ctg Category
		err := rows.Scan(&ctg.CatID, &ctg.CatName)
		if err != nil {
			return nil, fmt.Errorf("Category rows: %v", err)
		}
		categories = append(categories, ctg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Category rows error: %v", err)
	}
	return categories, nil
}
