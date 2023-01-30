package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/u8views/go-u8views/internal/models/oauth2"
	"github.com/u8views/go-u8views/internal/oauth2/github"
	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	"github.com/u8views/go-u8views/internal/templates"

	"github.com/gin-gonic/gin"
)

const (
	userCookieKey = "u8views_user_id"
)

type SocialProviderQuery struct {
	Code string `form:"code" binding:"required"`
}

type OAuth2Controller struct {
	userService *services.UserService
	github      oauth2.Secret
}

func NewOAuth2Controller(userService *services.UserService, github oauth2.Secret) *OAuth2Controller {
	return &OAuth2Controller{userService: userService, github: github}
}

func (s *OAuth2Controller) RedirectGitHubLogin(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, github.Redirect(s.github))
}

func (s *OAuth2Controller) CallbackGitHubLogin(ctx *gin.Context) {
	s.callbackLogin(ctx, s.github, github.User)
}

func (s *OAuth2Controller) Logout(ctx *gin.Context) {
	delCookieUserID(ctx)

	// Sending browser cookies during a 302 redirect
	// https://stackoverflow.com/questions/4694089/sending-browser-cookies-during-a-302-redirect
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.RawRedirect("/")))
}

func (s *OAuth2Controller) callbackLogin(ctx *gin.Context, secret oauth2.Secret, userGetter func(secret oauth2.Secret, code string) (oauth2.SocialProviderUser, error)) {
	var query SocialProviderQuery

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(fmt.Sprintf("HTTP GET query %q required", "code")))

		return
	}

	socialProviderUser, err := userGetter(secret, query.Code)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))

		return
	}

	userID, err := s.userService.Upsert(
		ctx,
		dbs.SocialProviderGithub,
		socialProviderUser.ID,
		socialProviderUser.Username,
		socialProviderUser.Name,
	)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))

		return
	}

	setCookieUserID(ctx, userID)

	// Sending browser cookies during a 302 redirect
	// https://stackoverflow.com/questions/4694089/sending-browser-cookies-during-a-302-redirect
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.RawRedirect("/github/"+socialProviderUser.Username)))
}

func setCookieUserID(ctx *gin.Context, userID int64) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		userCookieKey,
		strconv.FormatInt(userID, 10),
		14*86400,
		"/",
		ctx.Request.Host,
		true,
		true,
	)
}

func delCookieUserID(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		userCookieKey,
		"",
		0,
		"/",
		ctx.Request.Host,
		true,
		true,
	)
}

func parseCookieUserID(r *gin.Context) int64 {
	cookie, err := r.Cookie(userCookieKey)
	if err == http.ErrNoCookie {
		return 0
	}

	if err != nil {
		// NOP

		return 0
	}

	if cookie == "" {
		return 0
	}

	userID, err := strconv.ParseInt(cookie, 10, 64)
	if err != nil {
		// NOP

		return 0
	}

	return userID
}
