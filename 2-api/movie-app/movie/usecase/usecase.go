package usecase

import (
	"movie-backend/movie-app/movie"
	"movie-backend/movie-app/movie/repository"
)

type movieUseCase struct {
	movieRepo repository.MovieRepository
}

func (useCase *movieUseCase) SearchMovie(spec movie.MovieSearchSpec) (res movie.MovieSearchResponse, err error) {
	return useCase.movieRepo.Search(spec)
}

func NewMovieUseCase(movieRepo repository.MovieRepository) movie.UseCase {
	return &movieUseCase{movieRepo: movieRepo}
}
