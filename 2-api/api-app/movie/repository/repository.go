package repository

import (
	"context"
	"movie-backend/movie-app/movie/controller/rpc/pb"
)

type MovieRepository interface {
	Search(ctx context.Context, keyword string, page int64) (res pb.MovieSearchResponse, err error)
}
