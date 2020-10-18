package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"movie-backend/movie-app/config"
	"movie-backend/movie-app/movie/controller/rpc"
	"movie-backend/movie-app/movie/repository/http"
	"movie-backend/movie-app/movie/usecase"
	"net"
	"os"
)

var server *grpc.Server
var appConfig config.AppConfig
var env string

func init() {
	server = grpc.NewServer()

	// Initialize environment
	env = os.Getenv("STOCKBIT_ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigFile(fmt.Sprintf("movie-app/config/%s.json", env))
	_ = viper.ReadInConfig()
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	// Initialize movie service
	movieRepo := http.NewMovieRepository(appConfig.MovieDBConfig.BaseURL, appConfig.MovieDBConfig.APIKey)
	movieUC := usecase.NewMovieUseCase(movieRepo)
	rpc.NewGRPCServer(server, movieUC)
}

func main() {
	port := os.Getenv("PORT")
	var address string

	if port == "" {
		address = appConfig.Addr
	} else {
		address = fmt.Sprintf(":%v", port)
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
