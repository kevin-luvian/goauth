package auth

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/user"
)

func (u *UseCase) GetGoogleLoginURL(c context.Context, state string) string {
	return u.authRepo.GoogleLoginURL(c, state)
}

func (u *UseCase) GetGoogleProfileInfo(c context.Context, state string, code string) (user.User, error) {
	var usr user.User

	gInfo, err := u.authRepo.GetGoogleUserInfo(c, code)
	if err != nil {
		return usr, err
	}

	usr.Name = gInfo.Name
	usr.Email = gInfo.Email

	return usr, nil
}
