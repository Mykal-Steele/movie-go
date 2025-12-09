package main

import (
	"fmt"

	controller "github.com/Mykal-Steele/movie-go/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", controller.GetHello())

	router.GET("/movies", controller.GetMovies())
	router.GET("/genres", controller.GetGenre())
	if err := router.Run(":8080"); err != nil {
		fmt.Println("There is an error: ", err)
	}
}
