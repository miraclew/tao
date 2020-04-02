package sqlschema

import (
	"database/sql"
	"fmt"
	"strings"
)

type CreateTables struct {
	Items []*CreateTable
}

type CreateTable struct {
	TableName string
	Pk        string
	Columns   []*Column
	Keys      []*Key
}

type Column struct {
	Name    string
	Type    string
	Null    bool
	Default sql.NullString
}

type Key struct {
	Type    string // 1: PRIMARY KEY, 2: UNIQUE KEY, 3 KEY
	Name    string
	Columns []string
}

var mysqlKey = map[string]string{"PK": "PRIMARY KEY", "K": "KEY", "UK": "UNIQUE KEY"}

// PRIMARY KEY (`id`),
// KEY `group_id` (`group_id`,`user_id`)
func (k *Key) String() string {
	wrapColumns := make([]string, len(k.Columns))
	for i, column := range k.Columns {
		wrapColumns[i] = fmt.Sprintf("`%s`", column)
	}
	return fmt.Sprintf("%s (%s)", mysqlKey[k.Type], strings.Join(wrapColumns, ","))

}
