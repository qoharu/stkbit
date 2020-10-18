package http

import (
	"encoding/json"
	"fmt"
	"log"
	"movie-backend/movie-app/movie"
	"movie-backend/movie-app/movie/repository"
	"net/http"
	"strconv"
)

type movieRepository struct {
	movieDBBaseURL string
	movieDBAPIKey  string
}

func (mr *movieRepository) Search(spec movie.MovieSearchSpec) (response movie.MovieSearchResponse, err error) {
	searchURL := fmt.Sprintf(
		"%s?apikey=%s&s=%s&page=%s",
		mr.movieDBBaseURL,
		mr.movieDBAPIKey,
		spec.Keyword,
		strconv.Itoa(int(spec.Page)),
	)

	resp, err := http.Get(searchURL)

	if err != nil {
		log.Println(err)
		return movie.MovieSearchResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&response)
	log.Println(response)
	return response, nil
}

func NewMovieRepository(movieDBBaseURL string, movieDBAPIKey string) repository.MovieRepository {
	return &movieRepository{movieDBBaseURL: movieDBBaseURL, movieDBAPIKey: movieDBAPIKey}
}
