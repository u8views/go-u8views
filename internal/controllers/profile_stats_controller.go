package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/u8views/go-u8views/internal/badge"
	"github.com/u8views/go-u8views/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileStatsController struct {
	service *services.ProfileStatsService
}

func NewProfileStatsController(service *services.ProfileStatsService) *ProfileStatsController {
	return &ProfileStatsController{service: service}
}

type ProfileCountURI struct {
	UserID int64 `uri:"user_id" binding:"required"`
}

type ProfileCountResponse struct {
	DayCount   int64 `json:"day_count"`
	WeekCount  int64 `json:"week_count"`
	MonthCount int64 `json:"month_count"`
	TotalCount int64 `json:"total_count"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (c *ProfileStatsController) Count(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx)
	if done {
		return
	}

	ctx.JSON(http.StatusOK, &ProfileCountResponse{
		DayCount:   statsCount.DayCount,
		WeekCount:  statsCount.WeekCount,
		MonthCount: statsCount.MonthCount,
		TotalCount: statsCount.TotalCount,
	})
}

func (c *ProfileStatsController) CountBadge(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx)
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

	ctx.Data(http.StatusOK, "image/svg+xml", []byte(statsBadge))
}

func (c *ProfileStatsController) statsCount(ctx *gin.Context) (services.ProfileStatsCount, bool) {
	var uri ProfileCountURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("Cannot parse UserID from URI %s\n", err)

		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "Cannot parse UserID from URI",
		})

		return services.ProfileStatsCount{}, true
	}

	statsCount, err := c.service.StatsCount(ctx, uri.UserID, true)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "User not found",
		})

		return services.ProfileStatsCount{}, true
	}

	if err != nil {
		log.Printf("Database error %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return services.ProfileStatsCount{}, true
	}
	return statsCount, false
}

func statsBadge(day, week, month, total int64) (string, error) {
	dayCount, err := createBadge("Views per day", day)
	if err != nil {
		return "", err
	}
	weekCount, err := createBadge("Views per week", week)
	if err != nil {
		return "", err
	}
	monthCount, err := createBadge("Views per mount", month)
	if err != nil {
		return "", err
	}
	totalCount, err := createBadge("Total views", total)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<div style="display:flex; gap:10px"> %s %s %s %s</div>`, dayCount, weekCount, monthCount, totalCount), nil
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
