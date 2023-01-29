package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/u8views/go-u8views/internal/badge"
	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage/dbs"

	"github.com/gin-gonic/gin"
)

type ProfileStatsController struct {
	userService         *services.UserService
	profileStatsService *services.ProfileStatsService
}

func NewProfileStatsController(userService *services.UserService, profileStatsService *services.ProfileStatsService) *ProfileStatsController {
	return &ProfileStatsController{userService: userService, profileStatsService: profileStatsService}
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

func (c *ProfileStatsController) GitHubDayWeekMonthTotalCount(ctx *gin.Context) {
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

func (c *ProfileStatsController) TotalCountBadge(ctx *gin.Context) {
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

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", []byte(totalCountBadge))
}

func (c *ProfileStatsController) GitHubDayWeekMonthTotalCountBadge(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx, dbs.SocialProviderGithub)
	if done {
		return
	}

	statsBadge, err := statsBadge(
		statsCount.DayCount,
		statsCount.WeekCount,
		statsCount.MonthCount,
		statsCount.TotalCount,
	)
	if err != nil {
		log.Printf("Cannot generate badge %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Cannot generate badge",
		})

		return
	}

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", []byte(statsBadge))
}

func (c *ProfileStatsController) statsCount(ctx *gin.Context, provider dbs.SocialProvider) (statsCount services.ProfileViewsStats, done bool) {
	var uri ProfileCountURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("Cannot parse SocialProviderUserID from URI %s\n", err)

		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "Cannot parse SocialProviderUserID from URI",
		})

		return statsCount, true
	}

	userID, err := c.userService.GetBySocialProvider(ctx, provider, uri.SocialProviderUserID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "User not found (social)",
		})

		return statsCount, true
	}
	if err != nil {
		log.Printf("Database error (social) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return statsCount, true
	}

	statsCount, err = c.profileStatsService.StatsCount(ctx, userID, true)
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

func statsBadge(day, week, month, total int64) (string, error) {
	dayBadge, err := createBadge("Views per day", day)
	if err != nil {
		return "", err
	}
	weekBadge, err := createBadge("Views per week", week)
	if err != nil {
		return "", err
	}
	monthBadge, err := createBadge("Views per month", month)
	if err != nil {
		return "", err
	}
	totalBadge, err := createBadge("Total views", total)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<div style="display:flex; gap:10px"> %s %s %s %s</div>`, dayBadge, weekBadge, monthBadge, totalBadge), nil
}

func createBadge(subject string, count int64) (string, error) {
	svg, err := badge.RenderBytes(
		subject,
		strconv.FormatInt(count, 10),
	)
	if err != nil {
		return "", err
	}

	return string(svg), nil
}
