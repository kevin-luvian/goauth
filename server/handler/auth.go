package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/entity"
	"github.com/kevin-luvian/goauth/server/pkg/app"
	"github.com/kevin-luvian/goauth/server/pkg/e"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

func (h *Handler) HandlerAuthPing(r gin.IRoutes) gin.IRoutes {
	return r.GET("/ping", func(c *gin.Context) {
		app.Success(c, map[string]interface{}{
			"ok": true,
		})
	})
}

func (h *Handler) HandlerGoogleSignup(r gin.IRoutes) gin.IRoutes {
	return r.GET("/signup/google", func(c *gin.Context) {
		state := util.RandString(32)

		redirect := h.authUC.GetGoogleLoginURL(c, state)

		data := entity.GoogleOAuthState{
			Referer: c.Request.Referer(),
			Type:    entity.GoogleSignup,
		}

		gredis.SetStruct(e.CACHE_GOOGLE_STATE+state, data, 5*time.Minute)

		c.Header("Referer", c.Request.URL.String())
		c.Redirect(http.StatusTemporaryRedirect, redirect)
	})
}

func (h *Handler) HandlerGoogleLogin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/login/google", func(c *gin.Context) {
		state := util.RandString(32)

		redirect := h.authUC.GetGoogleLoginURL(c, state)

		data := entity.GoogleOAuthState{
			Referer: c.Request.Referer(),
			Type:    entity.GoogleLogin,
		}

		gredis.SetStruct(e.CACHE_GOOGLE_STATE+state, data)

		logging.Infoln("c.Request.Referer()", c.Request.Referer())

		c.Header("Referer", c.Request.URL.String())
		c.Redirect(http.StatusTemporaryRedirect, redirect)
	})
}

func (h *Handler) HandlerAuthenticateGoogleRedirectOrigin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/login/google/redirect", func(c *gin.Context) {
		var err error

		state := c.Query("state")
		code := c.Query("code")
		cacheData := entity.GoogleOAuthState{}

		// check state
		err = gredis.GetStruct(e.CACHE_GOOGLE_STATE+state, &cacheData)
		if err != nil {
			app.Error(c, e.FORBIDDEN, err)
			return
		}

		gredis.Delete(e.CACHE_GOOGLE_STATE + state)

		usr, err := h.authUC.GetGoogleProfileInfo(c, state, code)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		if cacheData.Type == entity.GoogleSignup {
			// check if user not exists

		} else if cacheData.Type == entity.GoogleLogin {
			// check if user exists
			app.Error(c, e.FORBIDDEN, errors.New("google login is not yet supported"))
			return
		}

		token, err := h.authUC.SignJWT(c, usr)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		redirectLocation := fmt.Sprintf("%s?token=%s", cacheData.Referer, token)

		c.Redirect(http.StatusTemporaryRedirect, redirectLocation)
		// app.Success(c, map[string]interface{}{
		// 	"loc":   redirectLocation,
		// 	"usr":   usr,
		// 	"token": token,
		// })
	})
}
