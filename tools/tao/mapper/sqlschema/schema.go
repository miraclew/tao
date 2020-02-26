package sqlschema

type CreateTable struct {
	TableName string
	Pk        string
	Fields    []*Field
}

type Field struct {
	Name    string
	Type    string
	Null    bool
	Default string
}
