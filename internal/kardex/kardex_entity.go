package kardex

import (
	"time"
)


type Kardex struct {
	ID                 int64   `json:"kardex_id" db:"kar_id"`
	Description 		string `json:"kardex_description"  validate:"required" db:"kar_desc"`
	Type        	string `json:"kardex_type"  validate:"required" db:"kar_tipo"`
	KardexCreatedAt 	time.Time `json:"kardex_created_at" db:"kar_created_at"`
	Products   			[]KardexProduct `json:"kardex_products"  validate:"required"`
}

type KardexProduct struct {
	ProductID 		int64 `json:"prod_id" db:"prod_id"`  // ID del producto
	Amount    		int   `json:"pro_kar_amount" validate:"min=0" db:"pro_kar_amount"`  // Cantidad movida en la 
	ProductName     string  `json:"prod_name" db:"prod_name"`
}

  

