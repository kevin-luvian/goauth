package v1

import (
	"testing"

	"github.com/adams-sarah/test2doc/test"
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

var router *gin.Engine
var server *test.Server

func AddRoutes() {
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/tags", GetTags)
	}
}

func TestRunner(t *testing.T) {
	var err error

	router = gin.New()
	AddRoutes()
	test.RegisterURLVarExtractor(util.MakeGINRouterExtractor(router))

	server, err = test.NewServer(router)
	if err != nil {
		panic(err.Error())
	}
	defer server.Finish()
}
