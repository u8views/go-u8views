package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/u8views/go-u8views/internal/badge"
	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"
	tmv2 "github.com/u8views/go-u8views/internal/templates/v2"

	"github.com/gin-gonic/gin"
)

const (
	// language=SVG
	pixel = `<svg xmlns="http://www.w3.org/2000/svg" width="1" height="1">
  <rect width="1" height="1" fill="white" />
</svg>`
)

type StatsController struct {
	userService  *services.UserService
	statsService *services.StatsService
}

func NewStatsController(userService *services.UserService, statsService *services.StatsService) *StatsController {
	return &StatsController{userService: userService, statsService: statsService}
}

type ProfileCountURI struct {
	SocialProviderUserID string `uri:"social_provider_user_id" binding:"required"`
}

type ProfileViewsStatsResponse struct {
	DayCount   int64 `json:"day_count"`
	WeekCount  int64 `json:"week_count"`
	MonthCount int64 `json:"month_count"`
	TotalCount int64 `json:"total_count"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (c *StatsController) GitHubDayWeekMonthTotalCount(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx, dbs.SocialProviderGithub)
	if done {
		return
	}

	ctx.JSON(http.StatusOK, &ProfileViewsStatsResponse{
		DayCount:   statsCount.DayCount,
		WeekCount:  statsCount.WeekCount,
		MonthCount: statsCount.MonthCount,
		TotalCount: statsCount.TotalCount,
	})
}

func (c *StatsController) DayCountBadge(ctx *gin.Context) {
	period := time.Now().UTC().Truncate(time.Hour).AddDate(0, 0, -1)

	statsCount, done := c.timePeriodStatsCount(ctx, dbs.SocialProviderGithub, period)
	if done {
		return
	}

	c.renderImage(ctx, []byte(tmv2.DayBadge(statsCount)))
}

func (c *StatsController) WeekCountBadge(ctx *gin.Context) {
	period := time.Now().UTC().Truncate(time.Hour).AddDate(0, 0, -7)

	statsCount, done := c.timePeriodStatsCount(ctx, dbs.SocialProviderGithub, period)
	if done {
		return
	}

	c.renderImage(ctx, []byte(tmv2.WeekBadge(statsCount)))
}

func (c *StatsController) MonthCountBadge(ctx *gin.Context) {
	period := time.Now().UTC().Truncate(time.Hour).AddDate(0, -1, 0)

	statsCount, done := c.timePeriodStatsCount(ctx, dbs.SocialProviderGithub, period)
	if done {
		return
	}

	c.renderImage(ctx, []byte(tmv2.MonthBadge(statsCount)))
}

func (c *StatsController) TotalCountBadge(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx, dbs.SocialProviderGithub)
	if done {
		return
	}

	totalCountBadge, err := createBadge("Total views", statsCount.TotalCount)
	if err != nil {
		log.Printf("Cannot generate badge %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Cannot generate badge",
		})

		return
	}

	c.renderImage(ctx, totalCountBadge)
}

func (c *StatsController) GitHubDayWeekMonthTotalCountBadge(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx, dbs.SocialProviderGithub)
	if done {
		return
	}

	// Temporary solution to reduce traffic volume
	if statsCount.HourCount > 1024 {
		c.renderImage(ctx, []byte(pixel))

		return
	}

	c.renderImage(ctx, []byte(tmv2.Badge(statsCount)))
}

func (c *StatsController) Pixel(ctx *gin.Context) {
	done := c.increment(ctx, dbs.SocialProviderGithub)
	if done {
		return
	}

	c.renderImage(ctx, []byte(pixel))
}

func (c *StatsController) GitHubStats(ctx *gin.Context) {
	socialProviderUserID, done := c.parseSocialProviderUserID(ctx)
	if done {
		return
	}

	userID, done := c.toUserID(ctx, dbs.SocialProviderGithub, socialProviderUserID)
	if done {
		return
	}

	result, err := c.statsService.Stats(ctx, userID)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *StatsController) UsersCreatedAtStatsByDay(ctx *gin.Context) {
	result, err := c.userService.UsersCreatedAtStatsByDay(ctx)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *StatsController) ReferralsStats(ctx *gin.Context) {
	socialProviderUserID, done := c.parseSocialProviderUserID(ctx)
	if done {
		return
	}

	userID, done := c.toUserID(ctx, dbs.SocialProviderGithub, socialProviderUserID)
	if done {
		return
	}

	result, err := c.statsService.ReferralsStats(ctx, userID)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *StatsController) timePeriodStatsCount(ctx *gin.Context, provider dbs.SocialProvider, period time.Time) (statsCount int64, done bool) {
	socialProviderUserID, done := c.parseSocialProviderUserID(ctx)
	if done {
		return
	}

	userID, done := c.toUserID(ctx, provider, socialProviderUserID)
	if done {
		return
	}

	increment := strings.HasPrefix(ctx.GetHeader("User-Agent"), "github-camo")

	statsCount, err := c.statsService.TimePeriodStatsCount(ctx, userID, increment, period)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return statsCount, true
	}

	return statsCount, false
}

func (c *StatsController) statsCount(ctx *gin.Context, provider dbs.SocialProvider) (statsCount services.ProfileViewsStats, done bool) {
	socialProviderUserID, done := c.parseSocialProviderUserID(ctx)
	if done {
		return
	}

	userID, done := c.toUserID(ctx, provider, socialProviderUserID)
	if done {
		return
	}

	increment := strings.HasPrefix(ctx.GetHeader("User-Agent"), "github-camo")

	statsCount, err := c.statsService.StatsCount(ctx, userID, increment)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "User not found (stats)",
		})

		return statsCount, true
	}

	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return statsCount, true
	}

	return statsCount, false
}

func (c *StatsController) increment(ctx *gin.Context, provider dbs.SocialProvider) (done bool) {
	socialProviderUserID, done := c.parseSocialProviderUserID(ctx)
	if done {
		return
	}

	userID, done := c.toUserID(ctx, provider, socialProviderUserID)
	if done {
		return
	}

	increment := strings.HasPrefix(ctx.GetHeader("User-Agent"), "github-camo")
	if !increment {
		return
	}

	err := c.statsService.Increment(ctx, userID)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return true
	}

	return false
}

func (c *StatsController) parseSocialProviderUserID(ctx *gin.Context) (string, bool) {
	var uri ProfileCountURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("Cannot parse SocialProviderUserID from URI %s\n", err)

		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "Cannot parse SocialProviderUserID from URI",
		})

		return "", true
	}

	return uri.SocialProviderUserID, false
}

func (c *StatsController) toUserID(ctx *gin.Context, provider dbs.SocialProvider, socialProviderUserID string) (int64, bool) {
	userID, err := c.userService.GetBySocialProvider(ctx, provider, socialProviderUserID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "User not found (social)",
		})

		return 0, true
	}
	if err != nil {
		log.Printf("Database error (social) %s\n", err)
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return 0, true
	}

	return userID, false
}

func (c *StatsController) renderImage(ctx *gin.Context, data []byte) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", data)
}

func createBadge(subject string, count int64) ([]byte, error) {
	svg, err := badge.RenderBytes(
		subject,
		strconv.FormatInt(count, 10),
	)
	if err != nil {
		return nil, err
	}

	return svg, nil
}
