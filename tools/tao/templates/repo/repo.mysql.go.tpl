{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"database/sql"
	"{{.Module}}/{{.Pkg}}"
	"github.com/miraclew/tao/pkg/slice"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type MysqlRepo struct {
	DB *sqlx.DB
}

func (r *MysqlRepo) Query(ctx context.Context, q *QueryParams) ([]*{{.Pkg}}.{{.Name|title}}, error) {
	vs := []*{{.Pkg}}.{{.Name|title}}{}
	var err error
	var values []interface{}
	var selectStatement = "SELECT * FROM `{{.Name}}` "

	if len(q.Ids) > 0 {
		selectStatement += fmt.Sprintf(" WHERE id IN (?%s) ", strings.Repeat(", ?", len(q.Ids)-1))
		values = append(values, slice.Int64Slice(q.Ids).InterfaceSlice()...)
		q.Limit = 100
		q.Offset = 0
	} else if q.Filter != nil {
		cond, vs := q.Filter.Conditions()
		values = append(values, vs...)
		selectStatement += fmt.Sprintf("WHERE %s ", cond)
	}

	if q.Sort == "" {
		q.Sort = "Id desc"
	}
	selectStatement += " ORDER BY " + q.Sort

	values = append(values, q.Offset, q.Limit)
	err = r.DB.SelectContext(ctx, &vs, selectStatement+" LIMIT ?,?", values...)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (r *MysqlRepo) Get(ctx context.Context, q *GetParams) (*{{.Pkg}}.{{.Name|title}}, error) {
	var v {{.Pkg}}.{{.Name|title}}
	var err error
	if q.Id != 0 {
		err = r.DB.GetContext(ctx, &v, "SELECT * FROM `{{.Name}}` WHERE Id=?", q.Id)
	} else {
		cond, vs := q.Filter.Conditions()
		query := fmt.Sprintf("SELECT * FROM `{{.Name}}` WHERE %s", cond)
		err = r.DB.GetContext(ctx, &v, query, vs...)
	}
	if err != nil && err == sql.ErrNoRows {
		return nil, ErrRecordNotFound
	}
	return &v, err
}

func (r *MysqlRepo) Update(ctx context.Context, vs Values, id int64) error {
	if len(vs) == 0 {
		return nil
	}

	ss, values := vs.UpdateSet()
	values = append(values, id)

	query := fmt.Sprintf("UPDATE `{{.Name}}` SET %s WHERE Id=?", ss)
	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *MysqlRepo) Create(ctx context.Context, v *{{.Pkg}}.{{.Name|title}}) (int64, error) {
	{{- $fields := .PrimaryModel.InsertFields }}
	res, err := r.DB.ExecContext(ctx, `INSERT INTO ` + "`{{.Name}}`" + ` ({{$fields | join ", "}}) VALUES (?{{", ?" | repeat (int (sub (len $fields) 1))}})`,
		{{range $fields}}v.{{.}},{{end}})
	if err != nil {
		return 0, err
	}
	v.Id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return v.Id, nil
}

func (r *MysqlRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM `{{.Name}}` WHERE Id=?", id)
	return err
}
