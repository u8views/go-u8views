package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"

	"github.com/gin-gonic/gin"
)

type ProfileURI struct {
	Username string `uri:"username" binding:"required"`
}

type ProfileController struct {
	userService *services.UserService
}

func NewProfileController(userService *services.UserService) *ProfileController {
	return &ProfileController{userService: userService}
}

func (c *ProfileController) GitHubProfile(ctx *gin.Context) {
	var uri ProfileURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("Cannot parse Username from URI %s\n", err)

		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Cannot parse Username from URI"))

		return
	}

	user, err := c.userService.GetBySocialProviderUsername(ctx, dbs.SocialProviderGithub, uri.Username)
	if err == sql.ErrNoRows {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(fmt.Sprintf("User not found")))

		return
	}

	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))

		return
	}

	if parseCookieUserID(ctx) == user.ID {
		// current user so show "badge copy"
	} else if user.ID == 0 {
		// show GitHub signup proposal
	}

	ctx.File("./public/profile.html")
}
