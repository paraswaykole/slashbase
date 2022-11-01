package sbsql

import (
	"database/sql/driver"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/utils"
)

type CryptedData string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (cd *CryptedData) Scan(value interface{}) error {
	encryptedData := value.(string)

	decryptedData, err := utils.DecryptAES(encryptedData, config.GetConfig().CryptedDataSecret)
	if err != nil {
		*cd = CryptedData("")
		return err
	}
	*cd = CryptedData(decryptedData)
	return err
}

// Value return json value, implement driver.Valuer interface
func (cd CryptedData) Value() (driver.Value, error) {
	if len(cd) == 0 {
		return nil, nil
	}
	encryptedData, err := utils.EncryptAES(string(cd), config.GetConfig().CryptedDataSecret)
	return encryptedData, err
}
