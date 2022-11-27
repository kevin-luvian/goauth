package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/app"
	"github.com/kevin-luvian/goauth/server/pkg/e"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	fmt.Println("OK")
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"ok": true,
	})
}
