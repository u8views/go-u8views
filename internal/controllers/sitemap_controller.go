package controllers

import (
	"fmt"
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

const (
	sitemapProfilesLimit    = 1024
	sitemapProfilesMaxLimit = 8096
)

func (c *SitemapController) SitemapGithubProfilesIndex(ctx *gin.Context) {
	totalCount, err := c.userService.UsersCount(ctx)
	if err != nil {
		log.Printf("Database error (sitemap count) %s\n", err)
		c.renderError(ctx, http.StatusInternalServerError, "Database error")
		return
	}

	// Return sitemap index
	ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfilesIndex(totalCount, sitemapProfilesLimit)))
}

func (c *SitemapController) SitemapGithubProfiles(ctx *gin.Context) {
	type sitemapParams struct {
		Offset int32 `uri:"offset"`
		Limit  int32 `uri:"limit"`
	}

	var params sitemapParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		c.renderError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if params.Offset <= 0 {
		c.renderError(ctx, http.StatusBadRequest, "Offset required")
		return
	}

	if params.Limit <= 0 {
		c.renderError(ctx, http.StatusBadRequest, "Limit required")
		return
	}

	if params.Limit > sitemapProfilesMaxLimit {
		c.renderError(ctx, http.StatusBadRequest, fmt.Sprintf("Limit max %d", sitemapProfilesMaxLimit))
		return
	}

	usernames, err := c.userService.Usernames(ctx, params.Offset, params.Limit)
	if err != nil {
		log.Printf("Database error (sitemap paginated) %s\n", err)
		c.renderError(ctx, http.StatusInternalServerError, "Database error")
		return
	}

	ctx.Data(http.StatusOK, "application/xml", []byte(tmv2.SitemapGithubProfiles(usernames)))
}

func (c *SitemapController) renderError(ctx *gin.Context, code int, message string) {
	// language=XML
	errorMessage := `<?xml version="1.0" encoding="UTF-8"?>
<error>
    <message>` + message + `</message>
</error>`
	ctx.Data(code, "text/xml", []byte(errorMessage))
}
