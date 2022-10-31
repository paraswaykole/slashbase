package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models/sbsql"
	"slashbase.com/backend/internal/utils"
)

type DBConnection struct {
	ID          string            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string            `gorm:"not null"`
	CreatedBy   string            `gorm:"not null"`
	ProjectID   string            `gorm:"not null"`
	Type        string            `gorm:"not null"`
	DBHost      sbsql.CryptedData `gorm:"type:text"`
	DBPort      sbsql.CryptedData `gorm:"type:text"`
	DBName      sbsql.CryptedData `gorm:"type:text"`
	DBScheme    sbsql.CryptedData `gorm:"type:text"`
	LoginType   string            `gorm:"not null;default:USE_ROOT;"`
	UseSSH      string            `gorm:"not null"`
	SSHHost     sbsql.CryptedData `gorm:"type:text"`
	SSHUser     sbsql.CryptedData `gorm:"type:text"`
	SSHPassword sbsql.CryptedData `gorm:"type:text"`
	SSHKeyFile  sbsql.CryptedData `gorm:"type:text"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`

	CreatedByUser     User               `gorm:"foreignkey:CreatedBy"`
	Project           Project            `gorm:"foreignkey:ProjectID"`
	DBConnectionUsers []DBConnectionUser `gorm:"foreignKey:DBConnectionID;constraint:OnDelete:CASCADE;"`
	ConnectionUser    *DBConnectionUser  `gorm:"-"`
}

type DBConnectionUser struct {
	ID             string            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	DBUser         sbsql.CryptedData `gorm:"type:text"`
	DBPassword     sbsql.CryptedData `gorm:"type:text"`
	UserIDs        pq.StringArray    `gorm:"type:text[]"`
	DBConnectionID string            `gorm:"not null"`
	IsRoot         bool              `gorm:"not null;"`
	ForRole        sql.NullString
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

func NewDBConnection(userID string, projectID string, name string, dbtype string, dbscheme, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile string) (*DBConnection, error) {

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

	if name == "" || dbhost == "" || dbport == "" || dbuser == "" ||
		dbpassword == "" || databaseName == "" {
		return nil, errors.New("cannot be empty")
	}

	connUser := DBConnectionUser{
		DBUser:     sbsql.CryptedData(dbuser),
		DBPassword: sbsql.CryptedData(dbpassword),
		IsRoot:     true,
	}

	return &DBConnection{
		Name:              name,
		CreatedBy:         userID,
		ProjectID:         projectID,
		Type:              dbtype,
		DBHost:            sbsql.CryptedData(dbhost),
		DBPort:            sbsql.CryptedData(dbport),
		DBName:            sbsql.CryptedData(databaseName),
		DBScheme:          sbsql.CryptedData(dbscheme),
		LoginType:         DBLOGINTYPE_ROOT,
		UseSSH:            useSSH,
		SSHHost:           sbsql.CryptedData(sshHost),
		SSHUser:           sbsql.CryptedData(sshUser),
		SSHPassword:       sbsql.CryptedData(sshPassword),
		SSHKeyFile:        sbsql.CryptedData(sshKeyFile),
		DBConnectionUsers: []DBConnectionUser{connUser},
		ConnectionUser:    &connUser,
	}, nil
}

func (dbConn DBConnection) Save() error {
	return db.GetDB().Save(&dbConn).Error
}

func (dbConnUser DBConnectionUser) Save() error {
	return db.GetDB().Save(&dbConnUser).Error
}
