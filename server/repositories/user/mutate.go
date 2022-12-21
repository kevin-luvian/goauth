package user

import (
	"context"
	"errors"

	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
)

func (r *Repo) Create(ctx context.Context, usr user.User) (user.User, error) {
	defer prom.CollectRepoDuration("Create")()

	q := `INSERT INTO users (
			tag, name, email, hpass
		) VALUES (:values) RETURNING 
			id, tag, name, email, hpass, created_at, updated_at`

	row := r.db.QueryRowxContext(
		ctx,
		q,
		usr.Tag,
		usr.Name,
		usr.Email,
		usr.HPass,
	)

	returning := user.User{}

	err := row.StructScan(&returning)
	if err != nil {
		return user.User{}, err
	}

	return returning, nil
}

func (r *Repo) DeleteByID(ctx context.Context, id int) error {
	defer prom.CollectRepoDuration("DeleteByID")()

	q := `DELETE FROM users WHERE id = ?`

	result, err := r.db.ExecContect(ctx, q, id)
	if err != nil {
		return err
	}

	raff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if raff == 0 {
		return errors.New("id not found")
	}

	return nil
}
