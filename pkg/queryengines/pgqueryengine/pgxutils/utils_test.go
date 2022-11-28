package pgxutils

import (
	"testing"
)

func TestReadPGSQLType(t *testing.T) {
	sql := `SELECT * FROM users WHERE bio = 'UPDATE'`
	stype, _ := GetPSQLQueryType(sql)
	if stype != QUERY_READ {
		t.Error("stype:", stype)
	}
}

func TestWritePGSQLType(t *testing.T) {
	sql := `UPDATE users SET bio = 'SELECT' WHERE name = 'ALTER'`
	stype, _ := GetPSQLQueryType(sql)
	if stype != QUERY_WRITE {
		t.Error("stype:", stype)
	}
}

func TestModifySchemaPGSQLType(t *testing.T) {
	sql := `CREATE TABLE users(
		id serial PRIMARY KEY,
		name VARCHAR (255) UNIQUE NOT NULL
	 );`
	stype, _ := GetPSQLQueryType(sql)
	if stype != QUERY_MODIFY_SCHEMA {
		t.Error("stype:", stype)
	}
}

func TestWriteReturningRowsPGSQLType(t *testing.T) {
	sql := `UPDATE users SET bio = 'SELECT' WHERE name = 'ALTER' RETURNING ctid;`
	stype, isReturningRows := GetPSQLQueryType(sql)
	if stype != QUERY_WRITE || !isReturningRows {
		t.Error("stype: ", stype)
		t.Error("isReturningRows: ", isReturningRows)
	}
}
