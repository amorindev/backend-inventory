package provider

type ProviderEntity struct {
	ProvID      int64  `json:"prov_id" db:"prov_id"`
	ProvName    string `json:"prov_name" db:"prov_name"`
	ProvAddress string `json:"prov_address" db:"prov_address"`
	ProvEmail   string `json:"prov_email" db:"prov_email"`
	ProvPhone   string `json:"prov_phone" db:"prov_phone"`
	ComID       int32  `json:"com_id" db:"com_user_id"`
}
