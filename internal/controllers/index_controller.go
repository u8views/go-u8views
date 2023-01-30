package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/templates"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	userService         *services.UserService
	profileStatsService *services.ProfileStatsService
}

func NewIndexController(userService *services.UserService, profileStatsService *services.ProfileStatsService) *IndexController {
	return &IndexController{userService: userService, profileStatsService: profileStatsService}
}

func (c *IndexController) Index(ctx *gin.Context) {
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
