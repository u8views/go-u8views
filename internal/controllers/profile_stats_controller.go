package controllers

import (
	"database/sql"
	"fmt"
	"github.com/u8views/go-u8views/internal/badge"
	"github.com/u8views/go-u8views/internal/services"
	"log"
	"net/http"
	"strconv"

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
	var uri ProfileCountURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("cannot parse uri %s\n", err)

		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "Cannot parse UserID from URI",
		})

		return
	}

	statsCount, err := c.service.StatsCount(ctx, uri.UserID, true)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusOK, &ProfileCountResponse{
			DayCount:   0,
			WeekCount:  0,
			MonthCount: 0,
			TotalCount: 0,
		})

		return
	}

	if err != nil {
		log.Printf("cannot parse uri %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	//ctx.JSON(http.StatusOK, &ProfileCountResponse{
	//	DayCount:   statsCount.DayCount,
	//	WeekCount:  statsCount.WeekCount,
	//	MonthCount: statsCount.MonthCount,
	//	TotalCount: statsCount.TotalCount,
	//})

	counters := ProfileCountResponse{
		DayCount:   statsCount.DayCount,
		WeekCount:  statsCount.WeekCount,
		MonthCount: statsCount.MonthCount,
		TotalCount: statsCount.TotalCount,
	}

	ctx.Writer.WriteHeader(http.StatusOK)

	ctx.Writer.Write([]byte(StatTemplate(
		strconv.FormatInt(counters.DayCount, 10),
		strconv.FormatInt(statsCount.WeekCount, 10),
		strconv.FormatInt(statsCount.MonthCount, 10),
		strconv.FormatInt(statsCount.TotalCount, 10),
	)))

}

func CreateBadge(subject string, status string, colorSVG badge.Color) string {
	svg, err := badge.RenderBytes(
		subject,
		status,
		colorSVG,
	)
	if err != nil {
		panic(err)
	}
	return string(svg)
}

func StatTemplate(day, week, month, total string) string {
	dayCount := CreateBadge("Views per day", day, "blue")
	weekCount := CreateBadge("Views per week", week, "green")
	monthCount := CreateBadge("Views per mount", month, "#ff0000")
	totalCount := CreateBadge("Total views", total, "orange")
	return fmt.Sprintf(`<div style="display:flex; gap:10px"> %s %s %s %s</div>`, dayCount, weekCount, monthCount, totalCount)

}
