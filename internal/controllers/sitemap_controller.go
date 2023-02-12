package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/u8views/go-u8views/internal/services"
	templates "github.com/u8views/go-u8views/internal/templates/v1"
	"log"
	"net/http"
)

type SiteMapController struct {
	userService *services.UserService
}

func NewSiteMapController(userService *services.UserService) *SiteMapController {
	return &SiteMapController{userService: userService}
}

func (c *SiteMapController) GetSiteMap(ctx *gin.Context) {
	users, err := c.userService.GetAllUsernames(ctx)
	if err != nil {
		log.Printf("Database error (stats) %s\n", err)

		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			ErrorMessage: "Database error",
		})

		return
	}

	profiles := make([]templates.ProfileView, len(users))
	for i, user := range users {
		profiles[i] = templates.ProfileView{
			Username: user.Username,
		}
	}
	ctx.Data(http.StatusOK, "text/xml; charset=utf-8", []byte(templates.SitemapGithubProfile(profiles)))
}
