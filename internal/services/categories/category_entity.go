package categories

type CategoryEntity struct {
	CatID   int64  `json:"cat_id" db:"cat_id"`
	CatName string `json:"cat_name" validate:"required" db:"cat_name"`
}
