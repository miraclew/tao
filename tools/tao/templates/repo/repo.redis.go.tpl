{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"github.com/redis-go/redis"
	"{{.Module}}/{{.Pkg}}"
)

type RedisRepo struct {
	client *redis.Client
}

func (r *RedisRepo) Query(ctx context.Context, q *QueryParams) ([]{{.Pkg}}.{{.Name|title}}, error) {
	return nil, nil
}

func (r *RedisRepo) Get(ctx context.Context, id string) (*{{.Pkg}}.{{.Name|title}}, error) {
	return nil, nil
}

func (r *RedisRepo) Update(ctx context.Context, attrs map[string]interface{}, id string) error {
	return nil
}

func (r *RedisRepo) Create(ctx context.Context, req *{{.Pkg}}.{{.Name|title}}) (string, error) {
	return "", nil
}

func (r *RedisRepo) Delete(ctx context.Context, id string) error {
	return nil
}
