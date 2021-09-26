package models

import (
	"errors"
	"time"

	"slashbase.com/backend/src/models/sbsql"
	"slashbase.com/backend/src/utils"
)

type DBConnection struct {
	ID            string            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name          string            `gorm:"not null"`
	CreatedBy     string            `gorm:"not null"`
	ProjectID     string            `gorm:"not null"`
	Type          string            `gorm:"not null"`
	DBHost        sbsql.CryptedData `gorm:"type:text"`
	DBPort        sbsql.CryptedData `gorm:"type:text"`
	DBPassword    sbsql.CryptedData `gorm:"type:text"`
	DBUser        sbsql.CryptedData `gorm:"type:text"`
	DBName        sbsql.CryptedData `gorm:"type:text"`
	UseSSH        string            `gorm:"not null"`
	SSHHost       sbsql.CryptedData `gorm:"type:text"`
	SSHUser       sbsql.CryptedData `gorm:"type:text"`
	SSHPassword   sbsql.CryptedData `gorm:"type:text"`
	SSHKeyFile    sbsql.CryptedData `gorm:"type:text"`
	CreatedAt     time.Time         `gorm:"autoCreateTime"`
	UpdatedAt     time.Time         `gorm:"autoUpdateTime"`
	CreatedByUser User              `gorm:"foreignkey:created_by"`
	Project       Project           `gorm:"foreignkey:project_id"`
}

const (
	DBTYPE_POSTGRES = "POSTGRES"

	DBUSESSH_NONE        = "NONE"
	DBUSESSH_PASSWORD    = "PASSWORD"
	DBUSESSH_KEYFILE     = "KEYFILE"
	DBUSESSH_PASSKEYFILE = "PASSKEYFILE"
)

func newDBConnection(userID string, projectID string, name string, dbtype string, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile string) (*DBConnection, error) {

	if !utils.ContainsString([]string{DBUSESSH_NONE, DBUSESSH_PASSWORD, DBUSESSH_KEYFILE, DBUSESSH_PASSKEYFILE}, useSSH) {
		return nil, errors.New("UseSSH is not correct")
	}

	return &DBConnection{
		Name:        name,
		CreatedBy:   userID,
		ProjectID:   projectID,
		Type:        dbtype,
		DBHost:      sbsql.CryptedData(dbhost),
		DBPort:      sbsql.CryptedData(dbport),
		DBPassword:  sbsql.CryptedData(dbpassword),
		DBName:      sbsql.CryptedData(databaseName),
		DBUser:      sbsql.CryptedData(dbuser),
		UseSSH:      useSSH,
		SSHHost:     sbsql.CryptedData(sshHost),
		SSHUser:     sbsql.CryptedData(sshUser),
		SSHPassword: sbsql.CryptedData(sshPassword),
		SSHKeyFile:  sbsql.CryptedData(sshKeyFile),
	}, nil
}

func NewPostgresDBConnection(userID string, projectID string, name string, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile string) (*DBConnection, error) {
	return newDBConnection(userID, projectID, name, DBTYPE_POSTGRES, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile)
}
