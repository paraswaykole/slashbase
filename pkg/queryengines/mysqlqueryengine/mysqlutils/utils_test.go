package mysqlutils

import "testing"

func TestReadMySQLType(t *testing.T) {
	sql := `SELECT * FROM users WHERE bio = 'UPDATE'`
	stype, _ := GetMySQLQueryType(sql)
	if stype != QUERY_READ {
		t.Error("stype:", stype)
	}
}

func TestWriteMySQLType(t *testing.T) {
	sql := `UPDATE users SET bio = 'SELECT' WHERE name = 'ALTER'`
	stype, _ := GetMySQLQueryType(sql)
	if stype != QUERY_WRITE {
		t.Error("stype:", stype)
	}
}

func TestModifySchemaMySQLType(t *testing.T) {
	sql := `CREATE TABLE users(
		id serial PRIMARY KEY,
		name VARCHAR (255) UNIQUE NOT NULL
	 );`
	stype, _ := GetMySQLQueryType(sql)
	if stype != QUERY_MODIFY_SCHEMA {
		t.Error("stype:", stype)
	}
}

func TestAlterSchemaMySQLType(t *testing.T) {
	sql := `ALTER TABLE users ADD Email varchar(255);;`
	stype, _ := GetMySQLQueryType(sql)
	if stype != QUERY_MODIFY_SCHEMA {
		t.Error("stype:", stype)
	}
}
