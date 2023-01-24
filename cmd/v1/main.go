package main

import (
	"github.com/gin-gonic/gin"
	"github.com/u8views/go-u8views/internal/controllers"
	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/env"
	"github.com/u8views/go-u8views/internal/services"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	var (
		dsn  = env.Must("POSTGRES_DSN")
		port = env.Must("PORT")
	)

	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	var (
		profileStatsController = controllers.NewProfileStatsController(services.NewProfileStatsService(repository))
	)

	var r = gin.New()

	r.GET("/:user_id/count", profileStatsController.Count)
	r.GET("/:user_id/count.svg", profileStatsController.CountBadge)

	r.LoadHTMLGlob("cmd/v1/public/*.html")

	r.Static("/assets/files", "./cmd/v1/public/assets/files")

	//r.StaticFS("/", http.Dir("./cmd/v1/public/"))

	r.GET("/u8views-main.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "u8views-main.html", gin.H{})
	})

	r.GET("/u8views-profile.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "u8views-profile.html", gin.H{})
	})

	r.GET("/u8views-stat.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "u8views-stat.html", gin.H{})
	})

	var serverErr = r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if serverErr != nil {
		log.Fatalln(serverErr)

		return
	}
}
