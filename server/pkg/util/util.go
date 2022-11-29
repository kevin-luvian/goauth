package util

import (
	"net/http"
	"strings"

	"github.com/adams-sarah/test2doc/doc/parse"
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func MakeGINRouterExtractor(engine *gin.Engine) parse.URLVarExtractor {
	// routes := engine.Routes()
	// routes[0].Method
	// routes[0].
	return func(req *http.Request) map[string]string {
		// httprouter Lookup func needs a trailing slash on path
		// path := req.URL.Path
		// if !strings.HasSuffix(path, "/") {
		// 	path += "/"
		// }

		params := req.URL.Query()

		// _, params, ok := router.
		// if !ok {
		// 	return nil
		// }

		paramsMap := make(map[string]string, len(params))
		for k, p := range params {
			paramsMap[k] = strings.Join(p, ", ")
		}

		return paramsMap
	}
}
