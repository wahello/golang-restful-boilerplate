package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router = gin.Default()

// Router returns the api router
func Router() http.Handler {
	router.Use(func(ctx *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				resp := Response{}
				resp.Title = "Something went wrong"
				resp.Errors = rvr.(error)
				resp.Status = http.StatusInternalServerError
				resp.ServerJSON(ctx.Writer)
				return
			}
		}()
	})

	router.GET("/", func(ctx *gin.Context) {
		resp := Response{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"name": "movie-pie",
			},
		}
		resp.ServerJSON(ctx.Writer)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	v1 := router.Group("/v1")
	v1.POST("/login", login)
	v1.POST("/register", register)
	v1.GET("/profile", profile)

	movies := v1.Group("movies")
	movies.GET("/search", searchMovie)
	movies.GET("/favourite", favouriteMovie)
}
