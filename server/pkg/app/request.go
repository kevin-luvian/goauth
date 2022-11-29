package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Infoln(err.Key, err.Message)
	}
}
