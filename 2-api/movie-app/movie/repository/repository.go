package repository

import "movie-backend/movie-app/movie"

// MovieRepository ...
type MovieRepository interface {
	Search(spec movie.MovieSearchSpec) (movie.MovieSearchResponse, error)
}
