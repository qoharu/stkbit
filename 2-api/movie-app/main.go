package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"movie-backend/movie-app/movie/controller/rpc"
	"movie-backend/movie-app/movie/repository/http"
	"movie-backend/movie-app/movie/usecase"
	"net"
)

var server *grpc.Server

func init() {
	server = grpc.NewServer()

	// Initialize movie service
	movieRepo := http.NewMovieRepository("https://www.omdbapi.com/", "faf7e5bb&s")
	movieUC := usecase.NewMovieUseCase(movieRepo)
	rpc.NewGRPCServer(server, movieUC)
}

func main() {
	lis, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}