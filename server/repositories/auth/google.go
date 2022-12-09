package auth

import (
	"context"
	"encoding/json"
	"io"
)

// GoogleLoginURL get login url
func (r *Repo) GoogleLoginURL(ctx context.Context, state string) string {
	return r.oauth.AuthCodeURL(state)
}

// GetGoogleUserInfo get user info after authenticate
func (r *Repo) GetGoogleUserInfo(ctx context.Context, code string) error {
	var (
		gUserInfo struct {
			Name          string `json:"name" db:"name"`
			Email         string `json:"email" db:"email"`
			VerifiedEmail bool   `json:"verified_email"`
			Picture       string `json:"picture"`
			Hd            string `json:"hd"`
		}
	)

	tok, err := r.oauth.Exchange(ctx, code)
	if err != nil {
		return err
	}

	client := r.oauth.Client(ctx, tok)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return err
	}
	defer userInfo.Body.Close()

	data, _ := io.ReadAll(userInfo.Body)
	err = json.Unmarshal(data, &gUserInfo)
	if err != nil {
		return err
	}

	return nil
}
