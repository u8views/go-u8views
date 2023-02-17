package main

import (
	"os"

	"github.com/u8views/go-u8views/internal/controllers"
	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/env"
	"github.com/u8views/go-u8views/internal/models/oauth2"
	"github.com/u8views/go-u8views/internal/server"
	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	var (
		dsn = env.Must("POSTGRES_DSN")
	)

	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	storage.MustMigrateUp(pgConnection)

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	var (
		userService  = services.NewUserService(repository)
		statsService = services.NewStatsService(repository)

		oauth2Controller = controllers.NewOAuth2Controller(userService, oauth2.Secret{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURI:  "",
			Scope:        "",
		})
		webController     = controllers.NewWebController(userService, statsService)
		statsController   = controllers.NewStatsController(userService, statsService)
		sitemapController = controllers.NewSitemapController(userService)
	)

	var r = gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/count.json", statsController.GitHubDayWeekMonthTotalCount)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/stats.json", statsController.GitHubStats)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/day-week-month-total-count.svg", statsController.GitHubDayWeekMonthTotalCountBadge)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/total-count.svg", statsController.TotalCountBadge)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/referrals/stats.json", statsController.ReferralsStats)
	r.GET("/api/v1/users/stats.json", statsController.UsersCreatedAtStatsByDay)

	r.
		GET("/login/github", oauth2Controller.RedirectGitHubLogin).
		GET("/oauth/callback/github", oauth2Controller.CallbackGitHubLogin).
		GET("/logout", oauth2Controller.Logout)

	r.GET("/", webController.Index)
	r.GET("/design", func(ctx *gin.Context) {
		ctx.File("./public/design/v2/index-auth.html")
	})

	r.GET("/github/:username", webController.GitHubProfile)
	r.GET("/design/github/:username", func(ctx *gin.Context) {
		ctx.File("./public/design/v2/profile-auth.html")
	})

	r.GET("/sitemap-github-profiles.xml", sitemapController.SitemapGithubProfiles)

	r.GET("/stats", webController.Stats)
	r.GET("/design/stats", func(ctx *gin.Context) {
		ctx.File("./public/design/v2/stats-auth.html")
	})

	r.Static("/assets/images", "./public/assets/images")
	r.Static("/assets/js", "./public/assets/js")

	r.StaticFile("/android-icon-144x144.png", "./public/android-icon-144x144.png")
	r.StaticFile("/android-icon-192x192.png", "./public/android-icon-192x192.png")
	r.StaticFile("/android-icon-36x36.png", "./public/android-icon-36x36.png")
	r.StaticFile("/android-icon-48x48.png", "./public/android-icon-48x48.png")
	r.StaticFile("/android-icon-72x72.png", "./public/android-icon-72x72.png")
	r.StaticFile("/android-icon-96x96.png", "./public/android-icon-96x96.png")
	r.StaticFile("/apple-icon-114x114.png", "./public/apple-icon-114x114.png")
	r.StaticFile("/apple-icon-120x120.png", "./public/apple-icon-120x120.png")
	r.StaticFile("/apple-icon-144x144.png", "./public/apple-icon-144x144.png")
	r.StaticFile("/apple-icon-152x152.png", "./public/apple-icon-152x152.png")
	r.StaticFile("/apple-icon-180x180.png", "./public/apple-icon-180x180.png")
	r.StaticFile("/apple-icon-57x57.png", "./public/apple-icon-57x57.png")
	r.StaticFile("/apple-icon-60x60.png", "./public/apple-icon-60x60.png")
	r.StaticFile("/apple-icon-72x72.png", "./public/apple-icon-72x72.png")
	r.StaticFile("/apple-icon-76x76.png", "./public/apple-icon-76x76.png")
	r.StaticFile("/apple-icon.png", "./public/apple-icon.png")
	r.StaticFile("/apple-icon-precomposed.png", "./public/apple-icon-precomposed.png")
	r.StaticFile("/favicon-16x16.png", "./public/favicon-16x16.png")
	r.StaticFile("/favicon-32x32.png", "./public/favicon-32x32.png")
	r.StaticFile("/favicon-96x96.png", "./public/favicon-96x96.png")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.StaticFile("/ms-icon-144x144.png", "./public/ms-icon-144x144.png")
	r.StaticFile("/ms-icon-150x150.png", "./public/ms-icon-150x150.png")
	r.StaticFile("/ms-icon-310x310.png", "./public/ms-icon-310x310.png")
	r.StaticFile("/ms-icon-70x70.png", "./public/ms-icon-70x70.png")

	r.StaticFile("/manifest.json", "./public/manifest.json")
	r.StaticFile("/browserconfig.xml", "./public/browserconfig.xml")

	r.StaticFile("/humans.txt", "./public/humans.txt")
	r.StaticFile("/robots.txt", "./public/robots.txt")

	r.StaticFile("/sitemap-main.xml", "./public/sitemap-main.xml")
	r.StaticFile("/sitemap.xml", "./public/sitemap.xml")

	server.Run(r.Handler())
}
