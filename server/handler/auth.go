package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/app"
	"github.com/kevin-luvian/goauth/server/pkg/e"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
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
		parsed, err := url.Parse(c.Request.Referer())
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		referer := parsed.Scheme + "://" + parsed.Host + "/redirect"

		data := google.GoogleOAuthState{
			Referer: referer,
			Type:    google.GoogleSignup,
		}

		state := util.RandString(30)
		err = gredis.SetStruct(e.CACHE_GOOGLE_STATE+state, data, 12*time.Hour)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		redirect := h.authUC.GoogleRedirectURL(c, state)

		c.Redirect(http.StatusTemporaryRedirect, redirect)
	})
}

func (h *Handler) HandlerGoogleLogin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/login/google", func(c *gin.Context) {
		parsed, err := url.Parse(c.Request.Referer())
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		referer := parsed.Scheme + "://" + parsed.Host + "/redirect"
		state := util.RandString(30)
		redirect := h.authUC.GoogleRedirectURL(c, state)
		data := google.GoogleOAuthState{
			Referer: referer,
			Type:    google.GoogleLogin,
		}

		gredis.SetStruct(e.CACHE_GOOGLE_STATE+state, data, 12*time.Hour)

		c.Redirect(http.StatusTemporaryRedirect, redirect)
	})
}

func (h *Handler) HandlerAuthenticateGoogleRedirectOrigin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/redirect/google", func(c *gin.Context) {
		var (
			referer string
			err     error
		)

		// redirect return
		redirectError := func(err error) {
			redirectLocation := fmt.Sprintf("%s?err=%s", referer, err.Error())
			c.Redirect(http.StatusTemporaryRedirect, redirectLocation)
		}

		state := c.Query("state")
		code := c.Query("code")
		cacheData := google.GoogleOAuthState{}

		// check state
		err = gredis.GetStruct(e.CACHE_GOOGLE_STATE+state, &cacheData)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		referer = cacheData.Referer

		gredis.Delete(e.CACHE_GOOGLE_STATE + state)

		gInfo, err := h.authUC.GetGoogleProfileInfo(c, code)
		if err != nil {
			redirectError(err)
			return
		}

		var usr user.User = user.User{}
		if cacheData.Type == google.GoogleSignup {
			usr, err = h.authUC.Signup(c, "", gInfo.Name, gInfo.Email)
			if err != nil {
				redirectError(err)
				return
			}
		} else if cacheData.Type == google.GoogleLogin {
			usr, err = h.authUC.GetByEmail(c, gInfo.Email)
			if err != nil {
				redirectError(err)
				return
			}
		}

		accessToken, refreshToken, err := h.authUC.SignJWT(c, usr)
		if err != nil {
			redirectError(err)
			return
		}

		redirectLocation := fmt.Sprintf("%s?token=%s", referer, accessToken)
		setRefreshTokenCookie(c, refreshToken)
		c.Redirect(http.StatusTemporaryRedirect, redirectLocation)
	})
}

func (h *Handler) HandlerRefreshToken(r gin.IRoutes) gin.IRoutes {
	return r.GET("/refresh-token", func(c *gin.Context) {
		refreshToken := getRefreshTokenCookie(c)
		usr, err := h.authUC.ParseJWTRefreshToken(c, refreshToken)
		if err != nil {
			app.Error(c, e.FORBIDDEN, err)
			return
		}

		accessToken, refreshToken, err := h.authUC.SignJWT(c, usr)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		setRefreshTokenCookie(c, refreshToken)
		app.Success(c, gin.H{"token": accessToken})
	})
}
