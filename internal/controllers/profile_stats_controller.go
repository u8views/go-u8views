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

func (c *ProfileStatsController) CountOnlyTotalBadge(ctx *gin.Context) {
	statsCount, done := c.statsCount(ctx)
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

func (c *ProfileStatsController) CountHitsSeeyoufarmStyleBadge(ctx *gin.Context) {
	// language=SVG
	const staticSVG = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="103" height="20">
 <linearGradient id="smooth" x2="0" y2="100%">
   <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
   <stop offset="1" stop-opacity=".1"/>
 </linearGradient>

 <mask id="round">
   <rect width="103" height="20" rx="3" ry="3" fill="#fff"/>
 </mask>

 <g mask="url(#round)">
   <rect width="30" height="20" fill="#555555"/>
   <rect x="30" width="73" height="20" fill="#79C83D"/>
   <rect width="103" height="20" fill="url(#smooth)"/>
 </g>

 <g fill="#fff" text-anchor="middle" font-family="Verdana,DejaVu Sans,Geneva,sans-serif" font-size="11"> 
   <text x="16" y="15" fill="#010101" fill-opacity=".3">hits</text>
   <text x="16" y="14" fill="#fff">hits</text>
   <text x="65.5" y="15" fill="#010101" fill-opacity=".3">123</text>
   <text x="65.5" y="14" fill="#fff">123</text>
 </g>
</svg>`

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", []byte(staticSVG))
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

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
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
