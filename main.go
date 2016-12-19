package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func main() {

	r := gin.Default()

	r.GET("/movie", GetMovies)
	r.GET("/movie/:id", GetMovie)
	r.PUT("/movie", PutMovie)
	r.StaticFile("/swagger.yaml", "./swagger.yaml")

	r.Run(":8080")

}

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	Rate     int    `json:"rate" binding:"required"`
	Producer string `json:"producer" binding:"required"`
}

var movies []Movie

func init() {
	movies = make([]Movie, 0)

	data, err := ioutil.ReadFile("movies.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &movies); err != nil {
		panic(err)
	}

	fmt.Println(movies)
}

func GetMovies(c *gin.Context) {
	c.JSON(200, movies)
}
func GetMovie(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		Abort(c, 400, errors.New("ID should be int"))
		return
	}

	for _, movie := range movies {
		if movie.Id == id {
			c.JSON(200, movie)
			return
		}
	}

	Abort(c, 400, errors.New("Movie not found"))

}
func PutMovie(c *gin.Context) {
	var movie Movie

	if err := c.BindJSON(&movie); err != nil {
		Abort(c, 400, err)
		return
	}
	movie.Id = len(movies)
	movies = append(movies, movie)

	c.JSON(200, map[string]int{"id": movie.Id})
}

func Abort(c *gin.Context, code int, err error) {
	c.JSON(code, map[string]string{"error": err.Error()})
}
