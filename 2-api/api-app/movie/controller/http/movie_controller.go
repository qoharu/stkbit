package http

import (
	"github.com/gin-gonic/gin"
	"movie-backend/api-app/movie"
	"strconv"
)

type movieController struct {
	movieUseCase movie.MovieUseCase
}

func CreateMovieController(r *gin.RouterGroup, movieUseCase movie.MovieUseCase) {
	controller := &movieController{movieUseCase: movieUseCase}
	r.GET("/search", controller.Search)
}

func (mc *movieController) Search(c *gin.Context) {
	searchWord := c.Query("searchword")
	pagination, _ := strconv.Atoi(c.Query("pagination"))

	searchRes, err := mc.movieUseCase.Search(c, searchWord, int64(pagination))

	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    searchRes,
		})
	}

}
