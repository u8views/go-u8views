package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

const sitemapPageSize = 1024

func (c *SitemapController) SitemapGithubProfiles(ctx *gin.Context) {
	totalCount, err := c.userService.GetUsernamesCount(ctx)
	if err != nil {
		log.Printf("Database error (sitemap count) %s\n", err)
		c.renderError(ctx, "Database error")
		return
	}

	if totalCount <= sitemapPageSize {
		usernames, err := c.userService.GetAllUsernames(ctx)
		if err != nil {
			log.Printf("Database error (sitemap) %s\n", err)
			c.renderError(ctx, "Database error")
			return
		}
		ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfiles(usernames)))
		return
	}

	// Return sitemap index
	ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfilesIndex(totalCount, sitemapPageSize)))
}

func (c *SitemapController) SitemapGithubProfilesPaginated(ctx *gin.Context) {
	filename := ctx.Param("filename")

	if !strings.HasSuffix(filename, ".xml") {
		c.renderError(ctx, "Invalid file format")
		return
	}

	rangeStr := strings.TrimSuffix(filename, ".xml")

	parts := strings.Split(rangeStr, "-")
	log.Printf("Split parts: %v (length: %d)", parts, len(parts))

	if len(parts) != 2 {
		c.renderError(ctx, fmt.Sprintf("Invalid range format. Got %d parts: %v", len(parts), parts))
		return
	}

	begin, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		c.renderError(ctx, "Invalid begin parameter")
		return
	}

	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		c.renderError(ctx, "Invalid end parameter")
		return
	}

	if end-begin > sitemapPageSize {
		c.renderError(ctx, "Page size too large")
		return
	}

	usernames, err := c.userService.GetUsernamesPaginated(ctx, begin, end-begin)
	if err != nil {
		log.Printf("Database error (sitemap paginated) %s\n", err)
		c.renderError(ctx, "Database error")
		return
	}

	ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfiles(usernames)))
}

func (c *SitemapController) renderError(ctx *gin.Context, message string) {
	// language=XML
	errorMessage := `<?xml version="1.0" encoding="UTF-8"?>
<error>
    <message>` + message + `</message>
</error>`
	ctx.Data(http.StatusInternalServerError, "text/xml", []byte(errorMessage))
}
