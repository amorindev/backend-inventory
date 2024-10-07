package provider

import (
	"fmt"

	"github.com/amorindev/backend-inventory/internal/db"
)

func GetProviders() ([]ProviderEntity, error) {
	var providers []ProviderEntity
	query := `SELECT prov_id, prov_name, prov_address, prov_email, prov_phone, com_user_id
		FROM tb_provider`

	rows, err := db.DB.Query(query)
	if err != nil {
	  return nil, fmt.Errorf("get providers err - dto: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p ProviderEntity
		err = rows.Scan(&p.ProvID, &p.ProvName, &p.ProvAddress, &p.ProvEmail, &p.ProvPhone, &p.ComID)
		if err != nil {
		  return nil, fmt.Errorf("get providers scan err dto: %v", err)
		}
		providers = append(providers, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return providers, nil
}