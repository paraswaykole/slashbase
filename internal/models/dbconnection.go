package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/pkg/sbsql"
)

type DBConnection struct {
	ID          string            `gorm:"type:uuid;primaryKey"`
	Name        string            `gorm:"not null"`
	ProjectID   string            `gorm:"not null"`
	Type        string            `gorm:"not null"`
	DBScheme    sbsql.CryptedData `gorm:"type:text"`
	DBHost      sbsql.CryptedData `gorm:"type:text"`
	DBPort      sbsql.CryptedData `gorm:"type:text"`
	DBName      sbsql.CryptedData `gorm:"type:text"`
	DBUser      sbsql.CryptedData `gorm:"type:text"`
	DBPassword  sbsql.CryptedData `gorm:"type:text"`
	LoginType   string            `gorm:"not null;default:USE_ROOT;"`
	UseSSH      string            `gorm:"not null"`
	SSHHost     sbsql.CryptedData `gorm:"type:text"`
	SSHUser     sbsql.CryptedData `gorm:"type:text"`
	SSHPassword sbsql.CryptedData `gorm:"type:text"`
	SSHKeyFile  sbsql.CryptedData `gorm:"type:text"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`

	Project Project `gorm:"foreignkey:ProjectID"`
}

const (
	DBTYPE_POSTGRES = "POSTGRES"
	DBTYPE_MONGO    = "MONGO"

	DBUSESSH_NONE        = "NONE"
	DBUSESSH_PASSWORD    = "PASSWORD"
	DBUSESSH_KEYFILE     = "KEYFILE"
	DBUSESSH_PASSKEYFILE = "PASSKEYFILE"

	DBLOGINTYPE_ROOT = "USE_ROOT"
	// DBLOGINTYPE_ROLE_ACCOUNTS = "ROLE_ACCOUNTS"
)

func NewDBConnection(projectID string, name string, dbtype string, dbscheme, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile string) (*DBConnection, error) {

	if !utils.ContainsString([]string{DBUSESSH_NONE, DBUSESSH_PASSWORD, DBUSESSH_KEYFILE, DBUSESSH_PASSKEYFILE}, useSSH) {
		return nil, errors.New("useSSH is not correct")
	}

	if dbtype == DBTYPE_POSTGRES {
		dbscheme = "postgres"
	} else if dbtype == DBTYPE_MONGO {
		if !utils.ContainsString([]string{"mongodb", "mongodb+srv"}, dbscheme) {
			return nil, errors.New("invalid dbscheme")
		}
	} else {
		return nil, errors.New("dbtype is not correct")
	}

	if name == "" || dbhost == "" || dbport == "" || databaseName == "" {
		return nil, errors.New("cannot be empty")
	}

	return &DBConnection{
		ID:          uuid.NewString(),
		Name:        name,
		ProjectID:   projectID,
		Type:        dbtype,
		DBScheme:    sbsql.CryptedData(dbscheme),
		DBHost:      sbsql.CryptedData(dbhost),
		DBPort:      sbsql.CryptedData(dbport),
		DBName:      sbsql.CryptedData(databaseName),
		DBUser:      sbsql.CryptedData(dbuser),
		DBPassword:  sbsql.CryptedData(dbpassword),
		LoginType:   DBLOGINTYPE_ROOT,
		UseSSH:      useSSH,
		SSHHost:     sbsql.CryptedData(sshHost),
		SSHUser:     sbsql.CryptedData(sshUser),
		SSHPassword: sbsql.CryptedData(sshPassword),
		SSHKeyFile:  sbsql.CryptedData(sshKeyFile),
	}, nil
}

func (dbConn DBConnection) Save() error {
	return db.GetDB().Save(&dbConn).Error
}
