package controllers

import (
	"context"
	"net/http"
	"time"

	database "github.com/Mykal-Steele/movie-go/Server/MagicStreamMoviesServer/database"
	model "github.com/Mykal-Steele/movie-go/Server/MagicStreamMoviesServer/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var MovieCollection *mongo.Collection = database.OpenCollection("movies")
var GenreCollection *mongo.Collection = database.OpenCollection("genres")

func GetGenre() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var genres []model.Genre

		cursor, err := GenreCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch genres. "})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &genres); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies. "})
		}
		c.JSON(http.StatusOK, genres)
	}
}

// get movies from the database
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get ctx with 100 sec time out
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// schedules cancel as it can auto close
		defer cancel()

		//structure of the movies
		var movies []model.Movie

		// query the database, get the pointer to the data in the database as `cursor`
		cursor, err := MovieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies. "})
		}
		defer cursor.Close(ctx)

		// loop through entire result and put result into movies slice
		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies. "})
		}

		// return the movies
		c.JSON(http.StatusOK, movies)
	}
}

func GetHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "hello")
	}
}
