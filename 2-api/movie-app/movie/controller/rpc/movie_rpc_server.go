package rpc

import (
	"context"
	"google.golang.org/grpc"
	"movie-backend/movie-app/movie"
	"movie-backend/movie-app/movie/controller/rpc/pb"
	"strconv"
)

type gRPCServer struct {
	useCase movie.UseCase
}

func NewGRPCServer(server *grpc.Server, useCase movie.UseCase) {
	pb.RegisterMovieServiceServer(server, &gRPCServer{useCase: useCase})
}

func (g *gRPCServer) SearchMovie(ctx context.Context, spec *pb.MovieSearchSpec) (*pb.MovieSearchResponse, error) {

	result, err := g.useCase.SearchMovie(movie.MovieSearchSpec{
		Keyword: spec.Keyword,
		Page:    spec.Page,
	})

	if err != nil {
		return nil, err
	}

	var movieDisplays []*pb.MovieDisplay
	for _, movieRes := range result.Search {
		movieDisplays = append(movieDisplays, &pb.MovieDisplay{
			Title:  movieRes.Title,
			Year:   movieRes.Year,
			ImdbID: movieRes.ImdbID,
			Type:   movieRes.Type,
			Poster: movieRes.Poster,
		})
	}

	totalResult, _ := strconv.Atoi(result.TotalResults)
	response := pb.MovieSearchResponse{
		Search:       movieDisplays,
		TotalResults: int64(totalResult),
		Response:     result.Response == "True",
	}

	return &response, nil
}
