package controllers

import (
	"log"
	"net/http"

	"github.com/u8views/go-u8views/internal/services"

	"github.com/gin-gonic/gin"
)

type UserStatsController struct {
	userService *services.UserService
}

func NewUserStatsController(userService *services.UserService) *UserStatsController {
	return &UserStatsController{userService: userService}
}

func (c *UserStatsController) UsersCreatedAtStatsByHour(ctx *gin.Context) {
	result, err := c.userService.UsersCreatedAtStatsByHour(ctx)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	ctx.JSON(http.StatusOK, result)
}
