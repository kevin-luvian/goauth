package util

import (
	"fmt"

	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	fmt.Println("abcd setup")
}
