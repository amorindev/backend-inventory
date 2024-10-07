package company

type CompanyEntity struct {
	Companyid int64 `json:"com_id" db:"com_user_id"`
	CompanyName string `json:"com_name" db:"com_name"`
	CompanyWebsite string `json:"com_website" db:"com_website"`
	CompanyAddress  string `json:"com_address" db:"com_address"`
	CompanyPhone string `json:"com_phone" db:"com_phone"`
	CompanyEmail string `json:"com_email" db:"com_email"`
	CompanyLogo string `json:"com_logo" db:"com_logo"`
}
