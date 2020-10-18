package rpc

import (
	"context"
	"log"
	"movie-backend/api-app/movie/repository"
	"movie-backend/movie-app/movie/controller/rpc/pb"
)

type movieRPCRepository struct {
	serviceClient pb.MovieServiceClient
}

func (m movieRPCRepository) Search(ctx context.Context, keyword string, page int64) (res pb.MovieSearchResponse, err error) {
	result, err := m.serviceClient.SearchMovie(ctx, &pb.MovieSearchSpec{
		Keyword: keyword,
		Page:    page,
	})

	if err != nil {
		log.Println(err.Error())
		return pb.MovieSearchResponse{}, err
	}

	return *result, nil
}

func NewMovieRPCRepository(serviceClient pb.MovieServiceClient) repository.MovieRepository {
	return &movieRPCRepository{serviceClient: serviceClient}
}
