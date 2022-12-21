package user

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/db"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
)

func (r *Repo) Get(ctx context.Context, param db.GetDBParam) ([]user.User, int, error) {
	defer prom.CollectRepoDuration("Get")()

	var (
		users = []user.User{}
		total = 0
	)

	q := `SELECT 
			id,
			tag,
			name,
			email,
			hpass,
			created_at,
			updated_at
		  FROM users %s`
	qCount := `SELECT COUNT(1) FROM users %s`

	err := r.db.GetWithCount(ctx, q, qCount, param, &users, &total)

	return users, total, r.getError(err)
}
