package pgxutils

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func PgSqlRowsToJson(rows pgx.Rows) ([]string, []map[string]interface{}) {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			itype := FieldType(fieldDescriptions[i])
			valuePtrs[i] = reflect.New(itype).Interface() // allocate pointer to type
		}
		rows.Scan(valuePtrs...)

		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface() // dereference pointer
			if str, ok := val.(sql.NullString); ok {
				if str.Valid {
					entry[col] = str.String
				} else {
					entry[col] = nil
				}
				continue
			}
			if bol, ok := val.(sql.NullBool); ok {
				if bol.Valid {
					entry[col] = bol.Bool
				} else {
					entry[col] = nil
				}
				continue
			}
			if float, ok := val.(sql.NullFloat64); ok {
				if float.Valid {
					entry[col] = float.Float64
				} else {
					entry[col] = nil
				}
				continue
			}
			if inte, ok := val.(sql.NullInt32); ok {
				if inte.Valid {
					entry[col] = inte.Int32
				} else {
					entry[col] = nil
				}
				continue
			}
			if inte, ok := val.(sql.NullInt64); ok {
				if inte.Valid {
					entry[col] = inte.Int64
				} else {
					entry[col] = nil
				}
				continue
			}
			if time, ok := val.(sql.NullTime); ok {
				if time.Valid {
					entry[col] = time.Time.String()
				} else {
					entry[col] = nil
				}
				continue
			}
			if tid, ok := val.(pgtype.TID); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = fmt.Sprintf("(%d,%d)", tid.BlockNumber, tid.OffsetNumber)
				}
				continue
			}
			if tid, ok := val.(pgtype.TextArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.VarcharArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.BoolArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.UUIDArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.DateArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int2Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int4Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int8Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Float4Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Float8Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[col] = nil
				} else {
					entry[col] = tid.Elements
				}
				continue
			}
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	// jsonData, _ := json.Marshal(tableData)
	return columns, tableData
}

func FieldType(fd pgproto3.FieldDescription) reflect.Type {
	switch fd.DataTypeOID {
	case pgtype.Float8OID:
		return reflect.TypeOf(sql.NullFloat64{})
	case pgtype.Float4OID:
		return reflect.TypeOf(sql.NullFloat64{})
	case pgtype.Int8OID:
		return reflect.TypeOf(sql.NullInt64{})
	case pgtype.Int4OID:
		return reflect.TypeOf(sql.NullInt32{})
	case pgtype.Int2OID:
		return reflect.TypeOf(sql.NullInt32{})
	case pgtype.VarcharOID, pgtype.BPCharArrayOID, pgtype.TextOID, pgtype.BPCharOID, pgtype.UUIDOID, pgtype.NameOID, LtreeOID:
		return reflect.TypeOf(sql.NullString{})
	case pgtype.BoolOID:
		return reflect.TypeOf(sql.NullBool{})
	case pgtype.NumericOID:
		return reflect.TypeOf(sql.NullFloat64{})
	case pgtype.DateOID, pgtype.TimestampOID, pgtype.TimestamptzOID:
		return reflect.TypeOf(sql.NullTime{})
	case pgtype.ByteaOID:
		return reflect.TypeOf([]byte(nil))
	case pgtype.TIDOID:
		return reflect.TypeOf(pgtype.TID{})
	default:
		return reflect.TypeOf(new(interface{})).Elem()
	}
}

const (
	ERRCODE_INVALID_PASSWORD                    = "28P01" // worng password
	ERRCODE_INVALID_AUTHORIZATION_SPECIFICATION = "28000" // db does not exist
)

const (
	LtreeOID = 16411
)

const (
	QUERY_READ   = iota
	QUERY_WRITE  = iota
	QUERY_ALTER  = iota
	QUERY_UNKOWN = -1
)

func GetPSQLQueryType(query string) int {
	// TODO: better query parsing method needed
	filteredQuery := strings.TrimSpace(strings.ToLower(query))
	if strings.Contains(filteredQuery, "returning") {
		return QUERY_READ
	}
	if strings.Contains(filteredQuery, "update") || strings.Contains(filteredQuery, "insert") || strings.Contains(filteredQuery, "truncate") {
		return QUERY_WRITE
	}
	if strings.Contains(filteredQuery, "alter") || strings.Contains(filteredQuery, "drop") {
		return QUERY_ALTER
	}
	if strings.HasPrefix(filteredQuery, "select") {
		return QUERY_READ
	}
	return QUERY_UNKOWN
}
