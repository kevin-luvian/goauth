package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *Handler) HandlerGoogleLogin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/login/google", func(c *gin.Context) {
		state := util.RandString(32)

		redirect := h.authUC.GetGoogleLoginURL(c, state)

		gredis.Set(e.CACHE_GOOGLE_STATE+state, c.Request.Referer())

		logging.Infoln("c.Request.Referer()", c.Request.Referer())

		// session.Set("google-state", state)
		// session.Set("google-referer", c.Request.Referer())

		c.Header("Referer", c.Request.URL.String())
		c.Redirect(http.StatusTemporaryRedirect, redirect)
	})
}

func (h *Handler) HandlerAuthenticateGoogleRedirectOrigin(r gin.IRoutes) gin.IRoutes {
	return r.GET("/login/google/redirect", func(c *gin.Context) {
		var err error

		state := c.Query("state")
		code := c.Query("code")

		// check state
		referer, err := gredis.Get(e.CACHE_GOOGLE_STATE + state)
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

		// session.Set("google-referer", c.Request.Referer())

		// if !strings.Contains(usr.Email, "@tokopedia.com") {
		// 	errResponse(c, http.StatusUnauthorized, errors.New("unauthorized email address"))
		// 	return
		// }

		// eng, err := h.engineer.GetByEmail(c, usr.Email)
		// if err != nil {
		// 	log.Error(err)
		// }
		// usr.User.EngineerID = eng.ID

		// if eng.ID != 0 && usr.User.Access != user.FullAccess {
		// 	rbacGroups, _ := h.getUserGroupNames(usr.Email)
		// 	usr.Groups = rbacGroups
		// }

		// usr.RegisteredClaims = jwt.RegisteredClaims{
		// 	ExpiresAt: jwt.NewNumericDate(expirationTime),
		// }

		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, usr)
		// tokenString, err := token.SignedString([]byte(h.secret))
		// if err != nil {
		// 	errResponse(c, http.StatusInternalServerError, err)
		// 	return
		// }

		// referer, ok := store.Get("referer").(string)
		// if !ok {
		// 	errResponse(c, http.StatusInternalServerError, errors.New("invalid conversion"))
		// 	return
		// }

		// token := "tokenaskdnaksnd"

		// redirectURI := "h.redirectTo[urls.GetBaseURL(referer)]"
		// if redirectURI == "" {
		// 	app.Error(c, e.ERROR, errors.New("redirect uri not allowed"))
		// 	return
		// }

		token, err := h.authUC.SignJWT(c, usr)
		if err != nil {
			app.Error(c, e.ERROR, err)
			return
		}

		redirectLocation := fmt.Sprintf("%s?token=%s", referer, token)

		// c.Redirect(http.StatusTemporaryRedirect, redirectLocation)
		app.Success(c, map[string]interface{}{
			"loc":   redirectLocation,
			"usr":   usr,
			"token": token,
		})
	})
}
