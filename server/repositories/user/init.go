package user

//go:generate mockgen -source=./init.go -destination=mock_user_repo.go -package=user

import (
	"errors"
	"strings"

	"github.com/kevin-luvian/goauth/server/pkg/db"
)

type Repo struct {
	db *db.DB
}

func New(db *db.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) getError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case strings.Contains(err.Error(), `pq: insert or update on table "eis_mst_service" violates foreign key constraint "service_repository_id_fkey"`):
		return errors.New("err user not found")
	default:
		return err
	}
}
