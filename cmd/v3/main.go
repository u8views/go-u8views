package main

import (
	"log"

	"github.com/u8views/go-u8views/internal/controllers"
	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/env"
	"github.com/u8views/go-u8views/internal/server"
	"github.com/u8views/go-u8views/internal/services"
	"github.com/u8views/go-u8views/internal/storage"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	var (
		dsn  = env.Must("POSTGRES_DSN")
		port = env.Must("PORT")
	)

	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	storage.MustMigrateUp(pgConnection)

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	var (
		profileStatsController = controllers.NewProfileStatsController(services.NewProfileStatsService(repository))
	)

	var r = gin.New()

	r.GET("/:user_id/count", profileStatsController.Count)
	r.GET("/:user_id/count.svg", profileStatsController.CountBadge)

	r.Static("/assets/files", "./public/assets/files")

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/profile", func(c *gin.Context) {
		c.File("./public/profile.html")
	})

	r.GET("/stats", func(c *gin.Context) {
		c.File("./public/stats.html")
	})

	var serverErr = r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if serverErr != nil {
		log.Fatalln(serverErr)

		return
	}

	server.Run(r.Handler())
}
