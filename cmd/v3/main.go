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
		userService         = services.NewUserService(repository)
		profileStatsService = services.NewProfileStatsService(repository)

		oauth2Controller = controllers.NewOAuth2Controller(userService, oauth2.Secret{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURI:  "",
			Scope:        "",
		})
		profileController      = controllers.NewProfileController(userService, profileStatsService)
		profileStatsController = controllers.NewProfileStatsController(userService, profileStatsService)
	)

	var r = gin.New()

	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/count.json", profileStatsController.GitHubDayWeekMonthTotalCount)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/day-week-month-total-count.svg", profileStatsController.GitHubDayWeekMonthTotalCountBadge)
	r.GET("/api/v1/github/profiles/:social_provider_user_id/views/total-count.svg", profileStatsController.TotalCountBadge)

	r.Static("/assets/files", "./public/assets/files")
	r.StaticFile("/favicon.ico", "./public/assets/files/favicon.ico")

	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})

	r.GET("/profile", func(ctx *gin.Context) {
		ctx.File("./public/profile.html")
	})

	r.GET("/stats", func(ctx *gin.Context) {
		ctx.File("./public/stats.html")
	})

	r.
		GET("/login/github", oauth2Controller.RedirectGitHubLogin).
		GET("/oauth/callback/github", oauth2Controller.CallbackGitHubLogin).
		GET("/logout", oauth2Controller.Logout)

	r.GET("/github/:username", profileController.GitHubProfile)

	server.Run(r.Handler())
}
