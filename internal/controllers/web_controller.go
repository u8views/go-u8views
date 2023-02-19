package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	tmv1 "github.com/u8views/go-u8views/internal/templates/v1"
	tmv2 "github.com/u8views/go-u8views/internal/templates/v2"

	"github.com/gin-gonic/gin"
)

type ProfileURI struct {
	Username string `uri:"username" binding:"required"`
}

type WebController struct {
	userService  *services.UserService
	statsService *services.StatsService
}

func NewWebController(userService *services.UserService, statsService *services.StatsService) *WebController {
	return &WebController{userService: userService, statsService: statsService}
}

func (c *WebController) Index(ctx *gin.Context) {
	const limit = 32

	sessionProfile, totalCount := c.sessionProfile(ctx)
	charity := totalCount > 0

	users, err := c.userService.Users(ctx, limit)
	if err != nil {
		log.Printf("Cannot fetch users %s\n", err)

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmv2.Index(sessionProfile, c.exampleProfile(), charity, nil)))

		return
	}

	if len(users) == 0 {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmv2.Index(sessionProfile, c.exampleProfile(), charity, nil)))

		return
	}

	userIDs := make([]int64, len(users))
	for i, user := range users {
		userIDs[i] = user.ID
	}

	now := time.Now().UTC().Truncate(time.Hour)
	userViewsStatsMap, err := c.statsService.UserDayWeekMonthViewsStatsMap(ctx, userIDs, now)
	if err != nil {
		log.Printf("Cannot fetch stats %s\n", err)

		// NOP
	}

	profiles := make([]tmv2.FullProfileView, len(users))
	for i, user := range users {
		stats := userViewsStatsMap[user.ID]

		profiles[i] = tmv2.FullProfileView{
			ProfileView: tmv2.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			},
			ProfileViewsStats: tmv1.ProfileViewsStats{
				DayCount:   stats.DayCount,
				WeekCount:  stats.WeekCount,
				MonthCount: stats.MonthCount,
				TotalCount: user.TotalCount,
			},
			CreatedAt: user.CreatedAt,
		}
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmv2.Index(sessionProfile, c.exampleProfile(), charity, profiles)))
}

func (c *WebController) GitHubProfile(ctx *gin.Context) {
	var uri ProfileURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("Cannot parse Username from URI %s\n", err)

		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Cannot parse Username from URI"))

		return
	}

	sessionProfile, totalCount := c.sessionProfile(ctx)
	charity := totalCount > 0

	currentPageProfile, err := c.getProfileByUsername(ctx, sessionProfile, uri.Username)
	if err == sql.ErrNoRows {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(fmt.Sprintf("User not found")))

		return
	}
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))

		return
	}

	stats, err := c.statsService.StatsCount(ctx, currentPageProfile.ID, false)
	if err != nil {
		log.Printf("Cannot fetch views stats by id = %d err: %v", currentPageProfile.ID, err)
	}

	exampleProfile := sessionProfile
	if exampleProfile.ID == 0 {
		exampleProfile = c.exampleProfile()
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmv2.Profile(
		sessionProfile,
		currentPageProfile,
		exampleProfile,
		charity,
		stats,
	)))
}

func (c *WebController) Stats(ctx *gin.Context) {
	sessionProfile, totalCount := c.sessionProfile(ctx)
	charity := totalCount > 0

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmv2.Stats(sessionProfile, charity)))
}

func (c *WebController) sessionProfile(ctx *gin.Context) (tmv2.ProfileView, int64) {
	var (
		sessionProfile tmv2.ProfileView
		totalCount     int64
	)

	sessionUserID := parseCookieUserID(ctx, userCookieKey)
	if sessionUserID > 0 {
		user, err := c.userService.GetByID(ctx, sessionUserID)
		if err != nil {
			log.Printf("Cannot fetch user by id = %d err: %v", sessionUserID, err)

			// NOP
		} else {
			sessionProfile = tmv2.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			}
		}

		totalCount, err = c.statsService.TotalCount(ctx, sessionUserID)
		if err != nil {
			log.Printf("Cannot fetch total count by user_id = %d err: %v", sessionUserID, err)

			// NOP
		}
	}

	return sessionProfile, totalCount
}

func (c *WebController) getProfileByUsername(ctx *gin.Context, sessionProfile tmv2.ProfileView, username string) (tmv2.ProfileView, error) {
	if strings.EqualFold(username, sessionProfile.Username) {
		return sessionProfile, nil
	}

	user, err := c.userService.GetBySocialProviderUsername(ctx, dbs.SocialProviderGithub, username)
	if err != nil {
		return tmv2.ProfileView{}, err
	}

	return tmv2.ProfileView{
		ID:                   user.ID,
		SocialProviderUserID: user.SocialProviderUserID,
		Username:             user.Username,
		Name:                 user.Name,
	}, nil
}

func (c *WebController) exampleProfile() tmv2.ProfileView {
	return tmv2.ProfileView{
		ID:                   1,
		SocialProviderUserID: "63663261",
		Username:             "YaroslavPodorvanov",
		Name:                 "Yaroslav Podorvanov",
	}
}
