package controllers

import (
	"log"
	"net/http"

	"github.com/u8views/go-u8views/internal/services"
	tmv2 "github.com/u8views/go-u8views/internal/templates/v2"

	"github.com/gin-gonic/gin"
)

type SitemapController struct {
	userService *services.UserService
}

func NewSitemapController(userService *services.UserService) *SitemapController {
	return &SitemapController{userService: userService}
}

func (c *SitemapController) SitemapGithubProfiles(ctx *gin.Context) {
	users, err := c.userService.GetAllUsernames(ctx)
	if err != nil {
		log.Printf("Database error (sitemap) %s\n", err)

		// language=XML
		const errorMessage = `<?xml version="1.0" encoding="UTF-8"?>
<error>
	<message>Database error</message>
</error>`
		ctx.Data(http.StatusInternalServerError, "text/xml", []byte(errorMessage))

		return
	}

	usernames := make([]string, len(users))
	for i, user := range users {
		usernames[i] = user.Username
	}
	ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfiles(usernames)))
}
