package auth

import (
	"context"
	"encoding/json"
	"io"

	"github.com/kevin-luvian/goauth/server/entity/google"
)

// GoogleLoginURL get login url
func (r *Repo) GoogleLoginURL(ctx context.Context, state string) string {
	return r.oauth.AuthCodeURL(state)
}

// GetGoogleUserInfo get user info after authenticate
func (r *Repo) GetGoogleUserInfo(ctx context.Context, code string) (google.UserInfo, error) {
	var gUserInfo google.UserInfo

	tok, err := r.oauth.Exchange(ctx, code)
	if err != nil {
		return gUserInfo, err
	}

	client := r.oauth.Client(ctx, tok)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return gUserInfo, err
	}
	defer userInfo.Body.Close()

	data, _ := io.ReadAll(userInfo.Body)
	err = json.Unmarshal(data, &gUserInfo)
	if err != nil {
		return gUserInfo, err
	}

	return gUserInfo, nil
}
