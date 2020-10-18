package movie

import "context"

type Movie struct {
	Title     string `json:"title"`
	Year      string `json:"year"`
	ImdbID    string `json:"imdbID"`
	Type      string `json:"type"`
	PosterURL string `json:"posterURL"`
}

type MovieListResponse struct {
	Movies    []Movie `json:"movies"`
	TotalData int64   `json:"totalData"`
}

type MovieUseCase interface {
	Search(ctx context.Context, keyword string, page int64) (res MovieListResponse, err error)
}
