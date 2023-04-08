package sqliteutils

import (
	"database/sql"
	"reflect"
	"strconv"

	"github.com/xwb1989/sqlparser"
)

func SqliteRowsToJson(rows *sql.Rows) ([]string, []map[string]interface{}) {

	fieldDescriptions, _ := rows.ColumnTypes()
	columns, _ := rows.Columns()
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			itype := fieldDescriptions[i].ScanType()
			valuePtrs[i] = reflect.New(itype).Interface() // allocate pointer to type
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i := range columns {
			iStr := strconv.Itoa(i)
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface() // dereference pointer
			if str, ok := val.(sql.NullString); ok {
				if str.Valid {
					entry[iStr] = str.String
				} else {
					entry[iStr] = nil
				}
				continue
			}
			if bol, ok := val.(sql.NullBool); ok {
				if bol.Valid {
					entry[iStr] = bol.Bool
				} else {
					entry[iStr] = nil
				}
				continue
			}
			if float, ok := val.(sql.NullFloat64); ok {
				if float.Valid {
					entry[iStr] = float.Float64
				} else {
					entry[iStr] = nil
				}
				continue
			}
			if inte, ok := val.(sql.NullInt32); ok {
				if inte.Valid {
					entry[iStr] = inte.Int32
				} else {
					entry[iStr] = nil
				}
				continue
			}
			if inte, ok := val.(sql.NullInt64); ok {
				if inte.Valid {
					entry[iStr] = inte.Int64
				} else {
					entry[iStr] = nil
				}
				continue
			}
			if time, ok := val.(sql.NullTime); ok {
				if time.Valid {
					entry[iStr] = time.Time.String()
				} else {
					entry[iStr] = nil
				}
				continue
			}
			b, ok := val.(sql.RawBytes)
			if ok {
				if b == nil {
					v = nil
				} else {
					v = string(b)
				}
			} else {
				v = val
			}
			entry[iStr] = v
		}
		tableData = append(tableData, entry)
	}

	return columns, tableData
}

const (
	QUERY_READ          = iota
	QUERY_WRITE         = iota
	QUERY_MODIFY_SCHEMA = iota
	QUERY_UNKOWN        = -1
)

func GetSQLiteQueryType(query string) (queryType int, isReturningRows bool) {
	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return QUERY_UNKOWN, false
	}
	switch stmt.(type) {
	case *sqlparser.Select:
		return QUERY_READ, true
	case *sqlparser.Insert:
	case *sqlparser.Update:
		return QUERY_WRITE, false
	case *sqlparser.DDL:
		return QUERY_MODIFY_SCHEMA, false
	}
	return QUERY_UNKOWN, false
}
