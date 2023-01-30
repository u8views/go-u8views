package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	"github.com/u8views/go-u8views/internal/templates"

	"github.com/gin-gonic/gin"
)

type ProfileURI struct {
	Username string `uri:"username" binding:"required"`
}

type ProfileController struct {
	userService         *services.UserService
	profileStatsService *services.ProfileStatsService
}

func NewProfileController(userService *services.UserService, profileStatsService *services.ProfileStatsService) *ProfileController {
	return &ProfileController{userService: userService, profileStatsService: profileStatsService}
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

	stats, err := c.profileStatsService.StatsCount(ctx, user.ID, false)
	if err != nil {
		log.Printf("Cannot fetch views stats by id = %d err: %v", user.ID, err)
	}

	var currentUserView templates.ProfileView

	currentUserID := parseCookieUserID(ctx)
	if currentUserID > 0 {
		user, err := c.userService.GetByID(ctx, currentUserID)
		if err != nil {
			log.Printf("Cannot fetch user by id = %d err: %v", currentUserID, err)

			// NOP
		} else {
			currentUserView = templates.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			}
		}
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.Profile(
		templates.ProfileView{
			ID:                   user.ID,
			SocialProviderUserID: user.SocialProviderUserID,
			Username:             user.Username,
			Name:                 user.Name,
		},
		currentUserView,
		stats,
	)))
}
