package models

type DBConnection struct {
	ID          string
	Name        string
	Type        string
	DBScheme    string
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	LoginType   string
	UseSSH      string
	SSHHost     string
	SSHUser     string
	SSHPassword string
	SSHKeyFile  string
	UseSSL      bool
}

const (
	DBTYPE_POSTGRES = "POSTGRES"
	DBTYPE_MONGO    = "MONGO"
	DBTYPE_MYSQL    = "MYSQL"

	DBUSESSH_NONE        = "NONE"
	DBUSESSH_PASSWORD    = "PASSWORD"
	DBUSESSH_KEYFILE     = "KEYFILE"
	DBUSESSH_PASSKEYFILE = "PASSKEYFILE"

	DBLOGINTYPE_ROOT = "USE_ROOT"
	// DBLOGINTYPE_ROLE_ACCOUNTS = "ROLE_ACCOUNTS"
)
