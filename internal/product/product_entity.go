package product

type Product struct {
	ID                 int64   `json:"prod_id" db:"prod_id"`
	ProductName        string  `json:"prod_name" validate:"required" db:"prod_name"`
	ProductDescription string  `json:"prod_desc" validate:"required" db:"prod_desc"`
	ProductDiscount    int     `json:"prod_discount" validate:"min=0" db:"prod_discount"` // campo omitido 0
	ProductPrice       float64 `json:"prod_price" validate:"required,min=0" db:"prod_price"`
	ProductStk         int     `json:"prod_stk" validate:"required,min=0" db:"prod_stk"`
	CatID              int64   `json:"cat_id" validate:"required" db:"cat_id"`
}
