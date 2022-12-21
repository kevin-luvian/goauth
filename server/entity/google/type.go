package google

import "github.com/kevin-luvian/goauth/server/pkg/util"

type GoogleAuthType int8

const (
	GoogleLogin GoogleAuthType = iota
	GoogleSignup
)

type GoogleUserInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Hd            string `json:"hd"`
}

type GoogleOAuthState struct {
	Referer string
	Type    GoogleAuthType
}

func (o *GoogleOAuthState) Valid() bool {
	return util.Contains([]GoogleAuthType{GoogleLogin, GoogleSignup}, o.Type)
}
