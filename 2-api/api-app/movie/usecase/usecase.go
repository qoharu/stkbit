package usecase

import (
	"context"
	"movie-backend/api-app/movie"
	"movie-backend/api-app/movie/repository"
)

type movieUseCase struct {
	movieRepo repository.MovieRepository
}

func (muc movieUseCase) Search(ctx context.Context, keyword string, page int64) (res movie.MovieListResponse, err error) {
	result, err := muc.movieRepo.Search(ctx, keyword, page)

	var movieResults []movie.Movie
	for _, movieRes := range result.Search {
		movieResults = append(movieResults, movie.Movie{
			Title:     movieRes.Title,
			Year:      movieRes.Year,
			ImdbID:    movieRes.ImdbID,
			Type:      movieRes.Type,
			PosterURL: movieRes.Poster,
		})
	}

	return movie.MovieListResponse{
		Movies:    movieResults,
		TotalData: result.TotalResults,
	}, nil
}

func NewMovieUseCase(movieRepo repository.MovieRepository) movie.MovieUseCase {
	return movieUseCase{movieRepo: movieRepo}
}
