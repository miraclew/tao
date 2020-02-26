package sqlschema

import "github.com/miraclew/tao/tools/tao/parser/proto3"

func MapMessage2CreateTable(message *proto3.Message) (*CreateTable, error) {
	createTable := &CreateTable{
		TableName: message.Name,
		Pk:        "Id",
		Fields:    []*Field{},
	}

	for _, entry := range message.Entries {
		f, err := MapField(entry.Field)
		if err != nil {
			return nil, err
		}
		createTable.Fields = append(createTable.Fields, f)
	}

	return createTable, nil
}

func MapField(field *proto3.Field) (*Field, error) {
	fType := MapType(field.Type)
	var fDefault string
	fDefault = defaults[fType]

	f := &Field{
		Name:    field.Name,
		Type:    fType,
		Null:    false,
		Default: fDefault,
	}

	return f, nil
}

var defaults = map[string]string{
	"datetime": "CURRENT_TIMESTAMP",
}
