package company

import (
	"database/sql"
	"fmt"

	"github.com/amorindev/backend-inventory/internal/db"
)

func GetCompanyByIdDto(companyID int64) (CompanyEntity, error) {
	var c CompanyEntity

	query := `SELECT com_user_id, com_name, com_website, com_address, com_phone, com_email, com_logo
		FROM tb_company
		WHERE com_user_id = $1`

	err := db.DB.QueryRow(query, companyID).Scan(&c.Companyid, &c.CompanyName, &c.CompanyWebsite, &c.CompanyAddress, &c.CompanyPhone, &c.CompanyEmail, &c.CompanyLogo)

	if err != nil {
	  if err == sql.ErrNoRows {
		return CompanyEntity{}, fmt.Errorf("no company found with id %d", companyID)
	  }
	  return c, err
	}
	return c, nil
}