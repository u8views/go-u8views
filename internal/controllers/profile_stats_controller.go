package controllers

import (
	"database/sql"
	"log"
	"net/http"

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
	var uri ProfileCountURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		log.Printf("cannot parse uri %s\n", err)

		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			ErrorMessage: "Cannot parse UserID from URI",
		})

		return
	}

	statsCount, err := c.service.StatsCount(ctx, uri.UserID)
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

	ctx.JSON(http.StatusOK, &ProfileCountResponse{
		DayCount:   statsCount.DayCount,
		WeekCount:  statsCount.WeekCount,
		MonthCount: statsCount.MonthCount,
		TotalCount: statsCount.TotalCount,
	})
}
