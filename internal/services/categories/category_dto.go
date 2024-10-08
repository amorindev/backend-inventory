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

func UpdateCategoryDto(ctgId int64, c CategoryEntity) error {
	query := `UPDATE tb_category SET cat_name = $1 WHERE cat_id = $2`
	result, err := db.DB.Exec(query, c.CatName, ctgId)
	if err != nil {
	  return fmt.Errorf("update catefory Err: %v", err)
	}
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected Err %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}
	return nil
}

func Deletecategory(ctgID int64) error {
	query := `DELETE FROM tb_category WHERE cat_id = $1`

	result, err := db.DB.Exec(query, ctgID)
	if err != nil {
	  return fmt.Errorf("delete category Err: %v", err)
	}
	
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
	  return fmt.Errorf("category - Rows affected err %v",err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("category not found with id %d", ctgID)
	}
	return nil
}


