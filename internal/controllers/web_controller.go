package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	tv1 "github.com/u8views/go-u8views/internal/templates/v1"
	tv2 "github.com/u8views/go-u8views/internal/templates/v2"

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

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tv2.Index(sessionProfile, c.exampleProfile(), charity, nil)))

		return
	}

	if len(users) == 0 {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tv2.Index(sessionProfile, c.exampleProfile(), charity, nil)))

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

	profiles := make([]tv2.FullProfileView, len(users))
	for i, user := range users {
		stats := userViewsStatsMap[user.ID]

		profiles[i] = tv2.FullProfileView{
			ProfileView: tv2.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			},
			ProfileViewsStats: tv1.ProfileViewsStats{
				DayCount:   stats.DayCount,
				WeekCount:  stats.WeekCount,
				MonthCount: stats.MonthCount,
				TotalCount: user.TotalCount,
			},
			CreatedAt: user.CreatedAt,
		}
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tv2.Index(sessionProfile, c.exampleProfile(), charity, profiles)))
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

	stats, err := c.statsService.StatsCount(ctx, user.ID, false)
	if err != nil {
		log.Printf("Cannot fetch views stats by id = %d err: %v", user.ID, err)
	}

	var currentUserView tv1.ProfileView

	currentUserID := parseCookieUserID(ctx, userCookieKey)
	if currentUserID > 0 {
		user, err := c.userService.GetByID(ctx, currentUserID)
		if err != nil {
			log.Printf("Cannot fetch user by id = %d err: %v", currentUserID, err)

			// NOP
		} else {
			currentUserView = tv1.ProfileView{
				ID:                   user.ID,
				SocialProviderUserID: user.SocialProviderUserID,
				Username:             user.Username,
				Name:                 user.Name,
			}
		}
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tv1.Profile(
		tv1.ProfileView{
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
	sessionProfile, totalCount := c.sessionProfile(ctx)
	charity := totalCount > 0

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tv2.Stats(sessionProfile, charity)))
}

func (c *WebController) sessionProfile(ctx *gin.Context) (tv2.ProfileView, int64) {
	var (
		sessionProfile tv2.ProfileView
		totalCount     int64
	)

	sessionUserID := parseCookieUserID(ctx, userCookieKey)
	if sessionUserID > 0 {
		user, err := c.userService.GetByID(ctx, sessionUserID)
		if err != nil {
			log.Printf("Cannot fetch user by id = %d err: %v", sessionUserID, err)

			// NOP
		} else {
			sessionProfile = tv2.ProfileView{
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

func (c *WebController) exampleProfile() tv2.ProfileView {
	return tv2.ProfileView{
		ID:                   1,
		SocialProviderUserID: "63663261",
		Username:             "YaroslavPodorvanov",
		Name:                 "Yaroslav Podorvanov",
	}
}