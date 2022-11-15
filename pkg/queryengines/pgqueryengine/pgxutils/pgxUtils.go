package pgxutils

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"slashbase.com/backend/internal/utils"
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
			if tid, ok := val.(pgtype.TID); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = fmt.Sprintf("(%d,%d)", tid.BlockNumber, tid.OffsetNumber)
				}
				continue
			}
			if tid, ok := val.(pgtype.TextArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.VarcharArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.BoolArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.UUIDArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.DateArray); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int2Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int4Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Int8Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Float4Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			if tid, ok := val.(pgtype.Float8Array); ok {
				if tid.Status == pgtype.Null || tid.Status == pgtype.Undefined {
					entry[iStr] = nil
				} else {
					entry[iStr] = tid.Elements
				}
				continue
			}
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[iStr] = v
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
	filteredQuery := strings.TrimSpace(strings.ToLower(query))
	filteredQuery = strings.ReplaceAll(filteredQuery, "\n", " ")
	regexr, _ := regexp.Compile(`'.+'|".+"`)
	regexr.ReplaceAllString(filteredQuery, "")
	filteredQuery = strings.ReplaceAll(filteredQuery, "  ", " ")
	tokens := strings.Split(filteredQuery, " ")
	if utils.ContainsString(tokens, "returning") || utils.ContainsString(tokens, "with") {
		return QUERY_READ
	}
	if utils.ContainsString(tokens, "update") || utils.ContainsString(tokens, "insert") || utils.ContainsString(tokens, "truncate") {
		return QUERY_WRITE
	}
	if utils.ContainsString(tokens, "alter") || utils.ContainsString(tokens, "drop") || utils.ContainsString(tokens, "create") {
		return QUERY_ALTER
	}
	if utils.ContainsString(tokens, "select") {
		return QUERY_READ
	}
	return QUERY_UNKOWN
}

func QueryToDataModel(fieldQueryData []map[string]interface{}, constraintsQueryData []map[string]interface{}) []map[string]interface{} {
	fields := []map[string]interface{}{}

	constraintMap := map[int32]map[string]interface{}{}
	for _, constraint := range constraintsQueryData {
		conkey := constraint["0"].([]pgtype.Int2)
		for _, colKey := range conkey {
			constraintMap[int32(colKey.Int)] = constraint
		}
	}

	for _, fieldData := range fieldQueryData {
		conkey := fieldData["0"].(int32)
		constraint := constraintMap[conkey]
		field := map[string]interface{}{
			"name":       fieldData["1"].(string),
			"type":       fieldData["2"].(string),
			"isNullable": fieldData["3"].(string) == "YES",
			"isPrimary":  false,
		}
		tags := []string{}
		if constraint["2"] != nil {
			field["isPrimary"] = rune(constraint["2"].(int8)) == 'p'
			if rune(constraint["2"].(int8)) == 'u' {
				tags = append(tags, "Unique")
			}
			if rune(constraint["2"].(int8)) == 'c' {
				tags = append(tags, "Check: "+constraint["1"].(string))
			}
			if rune(constraint["2"].(int8)) == 'f' {
				tags = append(tags, "Foreign Key: "+constraint["1"].(string))
			}
			if rune(constraint["2"].(int8)) == 't' {
				tags = append(tags, "Trigger: "+constraint["1"].(string))
			}
			if rune(constraint["2"].(int8)) == 'x' {
				tags = append(tags, "Exclusion: "+constraint["1"].(string))
			}
		}
		if fieldData["4"] != nil {
			coldef := fieldData["4"].(string)
			tags = append(tags, "Default: "+coldef)
		}
		if fieldData["5"] != nil {
			maxLen := fieldData["5"].(int32)
			tags = append(tags, "Max Length: "+strconv.Itoa(int(maxLen)))
		}
		field["tags"] = tags
		fields = append(fields, field)
	}

	return fields
}
