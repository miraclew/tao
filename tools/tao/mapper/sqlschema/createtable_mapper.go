package sqlschema

import (
	"database/sql"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"fmt"
	"strings"
)

func MapCreateTables(messages []*proto3.Message) (*CreateTables, error) {
	res := &CreateTables{
		Items: []*CreateTable{},
	}
	for _, msg := range messages {
		ct, err := MapCreateTable(msg)
		if err != nil {
			return nil, err
		}
		res.Items = append(res.Items, ct)
	}
	return res, nil
}

func MapCreateTable(message *proto3.Message) (*CreateTable, error) {
	createTable := &CreateTable{
		TableName: message.Name,
		Pk:        "Id",
		Columns:   []*Column{},
	}

	for _, entry := range message.Entries {
		if entry.Field == nil {
			continue
		}
		if entry.Field.Type.Reference == "Key" {
			k, err := MapKey(entry.Field)
			if err != nil {
				return nil, err
			}
			createTable.Keys = append(createTable.Keys, k)
		} else {
			f, err := MapColumn(entry.Field)
			if err != nil {
				return nil, err
			}
			createTable.Columns = append(createTable.Columns, f)
		}
	}

	return createTable, nil
}

// Key PK_GroupId_UserId = 100;
// Key K_GroupId = 101;
// Key K_UserId = 102;
// Key UK_OwnerId = 103;

func MapKey(field *proto3.Field) (*Key, error) {
	parts := strings.Split(field.Name, "_")
	if len(parts) < 2 {
		return nil, fmt.Errorf("bad db key: %s", field.Name)
	}
	k := &Key{
		Type:    parts[0],
		Columns: parts[1:],
	}
	return k, nil
}

func MapColumn(field *proto3.Field) (*Column, error) {
	fType := GetDbType(field)
	if fType == "" {
		fType = MapType(field.Type)
	}

	var fDefault sql.NullString
	fDefault.Valid = true

	if field.Name == "Id" {
		fDefault.Valid = false
	}

	if fType == "datetime" {
		fDefault.String = "CURRENT_TIMESTAMP"
	} else if strings.HasPrefix(fType, "int") || strings.HasPrefix(fType, "bigint") || strings.HasPrefix(fType, "tinyint") {
		fDefault.String = "0"
	} else if strings.HasPrefix(fType, "char") || strings.HasPrefix(fType, "varchar") {
		fDefault.String = "''"
	} else {
		fDefault.Valid = false
	}

	f := &Column{
		Name:    field.Name,
		Type:    fType,
		Null:    false,
		Default: fDefault,
	}

	return f, nil
}

func GetDbType(field *proto3.Field) string {
	for _, option := range field.Options {
		if option.Name == "db_type" && option.Value.String != nil {
			return *option.Value.String
		}
	}
	return ""
}
