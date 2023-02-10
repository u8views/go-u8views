package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	templates "github.com/u8views/go-u8views/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type ProfileURI struct {
	Username string `uri:"username" binding:"required"`
}

type WebController struct {
	userService         *services.UserService
	profileStatsService *services.ProfileStatsService
}

func NewWebController(userService *services.UserService, profileStatsService *services.ProfileStatsService) *WebController {
	return &WebController{userService: userService, profileStatsService: profileStatsService}
}

func (c *WebController) Index(ctx *gin.Context) {
	const limit = 32

	users, err := c.userService.Users(ctx, limit)
	if err != nil {
		log.Printf("Cannot fetch users %s\n", err)

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.Index(nil)))

		return
	}

	if len(users) == 0 {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.Index(nil)))

		return
	}

	userIDs := make([]int64, len(users))
	for i, user := range users {
		userIDs[i] = user.ID
	}

	now := time.Now().UTC().Truncate(time.Hour)
	userViewsStatsMap, err := c.profileStatsService.UserDayWeekMonthViewsStatsMap(ctx, userIDs, now)
	if err != nil {
		log.Printf("Cannot fetch stats %s\n", err)

		// NOP
	}

	profiles := make([]templates.FullProfileView, len(users))
	for i, user := range users {
		stats := userViewsStatsMap[user.ID]

		profiles[i] = templates.FullProfileView{
			ProfileView: templates.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			},
			ProfileViewsStats: templates.ProfileViewsStats{
				DayCount:   stats.DayCount,
				WeekCount:  stats.WeekCount,
				MonthCount: stats.MonthCount,
				TotalCount: user.TotalCount,
			},
			CreatedAt: user.CreatedAt,
		}
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(templates.Index(profiles)))
}

func (c *WebController) GitHubProfile(ctx *gin.Context) {
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

func (c *WebController) Stats(ctx *gin.Context) {

}
