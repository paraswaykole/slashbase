package models

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/utils"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/sbsql"
)

type DBConnection struct {
	ID          string            `gorm:"type:uuid;primaryKey"`
	Name        string            `gorm:"not null;uniqueIndex"`
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

func NewDBConnection(projectID string, name string, dbtype string, dbscheme, dbhost, dbport, dbuser, dbpassword, databaseName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile string) (*DBConnection, error) {

	if !utils.ContainsString([]string{qemodels.DBUSESSH_NONE, qemodels.DBUSESSH_PASSWORD, qemodels.DBUSESSH_KEYFILE, qemodels.DBUSESSH_PASSKEYFILE}, useSSH) {
		return nil, errors.New("useSSH is not correct")
	}

	if dbtype == qemodels.DBTYPE_POSTGRES {
		dbscheme = "postgres"
	} else if dbtype == qemodels.DBTYPE_MONGO {
		if !utils.ContainsString([]string{"mongodb", "mongodb+srv"}, dbscheme) {
			return nil, errors.New("invalid dbscheme")
		}

		dbpassword = url.QueryEscape(dbpassword)

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
		LoginType:   qemodels.DBLOGINTYPE_ROOT,
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

func (dbConn *DBConnection) ToQEConnection() *qemodels.DBConnection {
	return &qemodels.DBConnection{
		ID:          dbConn.ID,
		Name:        dbConn.Name,
		Type:        dbConn.Type,
		DBScheme:    string(dbConn.DBScheme),
		DBHost:      string(dbConn.DBHost),
		DBPort:      string(dbConn.DBPort),
		DBName:      string(dbConn.DBName),
		DBUser:      string(dbConn.DBUser),
		DBPassword:  string(dbConn.DBPassword),
		LoginType:   string(dbConn.LoginType),
		UseSSH:      string(dbConn.UseSSH),
		SSHHost:     string(dbConn.SSHHost),
		SSHUser:     string(dbConn.SSHUser),
		SSHPassword: string(dbConn.SSHPassword),
		SSHKeyFile:  string(dbConn.SSHKeyFile),
	}
}
