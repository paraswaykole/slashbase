package models

import (
	"time"

	"slashbase.com/backend/models/sbsql"
)

type DBConnection struct {
	ID            string            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name          string            `gorm:"not null"`
	CreatedBy     string            `gorm:"unique;not null"`
	TeamID        string            `gorm:"not null"`
	Type          string            `gorm:"not null"`
	DBHost        sbsql.CryptedData `gorm:"type:text"`
	DBPort        sbsql.CryptedData `gorm:"type:text"`
	DBPassword    sbsql.CryptedData `gorm:"type:text"`
	DBUser        sbsql.CryptedData `gorm:"type:text"`
	DBName        sbsql.CryptedData `gorm:"type:text"`
	CreatedAt     time.Time         `gorm:"autoCreateTime"`
	UpdatedAt     time.Time         `gorm:"autoUpdateTime"`
	CreatedByUser User              `gorm:"foreignkey:created_by"`
	Team          Team              `gorm:"foreignkey:team_id"`
}

const (
	DBTYPE_POSTGRES = "POSTGRES"
)

func newDBConnection(userID string, teamID string, name string, dbtype string, dbhost, dbport, dbuser, dbpassword, databaseName string) *DBConnection {
	return &DBConnection{
		Name:       name,
		CreatedBy:  userID,
		TeamID:     teamID,
		Type:       dbtype,
		DBHost:     sbsql.CryptedData(dbhost),
		DBPort:     sbsql.CryptedData(dbport),
		DBPassword: sbsql.CryptedData(dbpassword),
		DBName:     sbsql.CryptedData(databaseName),
		DBUser:     sbsql.CryptedData(dbuser),
	}
}

func NewPostgresDBConnection(userID string, teamID string, name string, dbhost, dbport, dbuser, dbpassword, databaseName string) *DBConnection {
	return newDBConnection(userID, teamID, name, DBTYPE_POSTGRES, dbhost, dbport, dbuser, dbpassword, databaseName)
}
