package auth

import (
	"context"
	"errors"

	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/db"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

func (u *UseCase) GoogleRedirectURL(ctx context.Context, state string) string {
	return u.authRepo.GoogleRedirectURL(ctx, state)
}

func (u *UseCase) GetGoogleProfileInfo(ctx context.Context, code string) (google.GoogleUserInfo, error) {
	return u.authRepo.GetGoogleUserInfo(ctx, code)
}

func (u *UseCase) GetByEmail(ctx context.Context, email string) (user.User, error) {
	// get user information
	users, _, err := u.userRepo.Get(ctx, db.GetDBParam{
		Filters:      []db.Filter{{Field: "email", Value: email}},
		DisableCount: true,
	})
	if err != nil || len(users) <= 0 {
		return user.User{}, errors.New("user doesn't exists")
	}

	return users[0], nil
}

func (u *UseCase) Signup(ctx context.Context, tag, name, email string) (user.User, error) {
	// check if users already exists
	_, total, err := u.userRepo.Get(ctx, db.GetDBParam{
		Filters:    []db.Filter{{Field: "email", Value: email}},
		DisableGet: true,
	})
	if err != nil || total > 0 {
		return user.User{}, errors.New("user already exists")
	}

	if tag == "" {
		tag = u.GenerateTag(ctx)
	}

	usr, err := u.userRepo.Create(ctx, user.User{
		Tag:   tag,
		Name:  name,
		Email: email,
		HPass: "",
	})
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (u *UseCase) GenerateTag(ctx context.Context) string {
	tag := ""
	for i := 1; i <= 5; i++ {
		tag = util.RandString(17)

		_, total, err := u.userRepo.Get(ctx, db.GetDBParam{
			Filters:    []db.Filter{{Field: "tag", Value: tag}},
			DisableGet: true,
		})
		if err == nil && total == 0 {
			return tag
		} else if err != nil {
			logging.Errorf("failed finding user tag %s, err: %v", tag, err)
		}
	}

	return util.RandString(20)
}
