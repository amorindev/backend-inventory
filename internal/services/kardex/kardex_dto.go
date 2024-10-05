package kardex

import (
	"fmt"

	"github.com/amorindev/backend-inventory/internal/db"
)

func GetKardexDto() ([]Kardex, error) {
	query := `SELECT 
		k.kar_id, k.kar_desc, k.kar_type, k.kar_created_at,
		pk.pro_kar_amount, p.prod_id, p.prod_name
	FROM tb_kardex k
	JOIN tb_product_kardex pk ON k.kar_id = pk.kar_id
	JOIN tb_product p ON pk.prod_id = p.prod_id
	ORDER BY k.kar_created_at DESC`
	
	var kardexs []Kardex
	
	rows, err := db.DB.Query(query)
	if err != nil {
	  return nil, err
	}

	defer rows.Close()

	kardexMap := make(map[int64]*Kardex)
	for rows.Next() {
		var kardexID int64
		var kardex Kardex
		var kardexProduct KardexProduct

		err := rows.Scan(
			&kardexID, &kardex.Description,&kardex.Type,&kardex.KardexCreatedAt,
			&kardexProduct.Amount,&kardexProduct.ProductID, &kardexProduct.ProductName,
		)
		if err != nil {
		  return nil, err
		}
		
		if existingKardex, ok := kardexMap[kardexID]; ok {
			existingKardex.Products = append(existingKardex.Products, kardexProduct)
		}else{
			kardex.ID =kardexID
			kardex.Products = append(kardex.Products, kardexProduct)
			kardexMap[kardexID] = &kardex
		}
	}
		for _, kardex := range kardexMap {
			kardexs = append(kardexs, *kardex)
		}
	

	return kardexs, nil
}


func CreateKardexDto(k Kardex) (Kardex, []KardexProduct, error) {
	// si el kardex es de salida
	// verificar que el stock del producto que se desea salir sea mayor o igual al cantidad de salida
	var  insufficientStockProducts []KardexProduct
		
	//verificamos si el kardex es de salida o entrada
	if k.Type == "SALIDA" {
		for _, kp := range k.Products {
			var currentStock int
			err := db.DB.QueryRow(`SELECT prod_stk FROM tb_product WHERE prod_id = $1`, kp.ProductID).Scan(&currentStock)
			if err != nil {
				return Kardex{}, nil, fmt.Errorf("error fetching product stock on create kardex %v ", err)
			}	
			//verificar si hay insuficiente stock
			if currentStock < kp.Amount {
				insufficientStockProducts = append(insufficientStockProducts, kp)
				continue
			}
		}
	}

	if len(insufficientStockProducts) > 0 {
		return Kardex{},insufficientStockProducts,nil
	} 

	var kardexID int64

	query := `INSERT INTO tb_kardex (kar_desc, kar_type) VALUES ($1,$2) RETURNING kar_id`

	err := db.DB.QueryRow(query,k.Description, k.Type).Scan(&kardexID)
	if err != nil {
	  return Kardex{},nil, fmt.Errorf("created kardex failed: %v", err)
	}

	for _, kp := range k.Products {
		query = `INSERT INTO tb_product_kardex(prod_id,kar_id,pro_kar_amount) VALUES ($1,$2,$3)`
		result, err := db.DB.Exec(query, kp.ProductID, kardexID, kp.Amount)
		if err != nil {
			return Kardex{},nil, fmt.Errorf("create kardex Err %v", err)
		}
		var rowsAffected int64
		rowsAffected, err = result.RowsAffected()
		if err != nil {
			return Kardex{}, nil,fmt.Errorf("create kardex Err %v", err)
		}

		if rowsAffected == 0 {
			return Kardex{}, nil, fmt.Errorf("create kardex found")
		}
	}

	k.ID = kardexID
	return k, nil,nil
}

