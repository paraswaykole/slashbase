package sbsql

import (
	"database/sql/driver"

	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/utils"
)

type CryptedData string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (cd *CryptedData) Scan(value interface{}) error {
	encryptedData := value.(string)

	decryptedData, err := utils.DecryptAES(encryptedData, config.GetCryptedDataSecretKey())
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
	encryptedData, err := utils.EncryptAES(string(cd), config.GetCryptedDataSecretKey())
	return encryptedData, err
}
