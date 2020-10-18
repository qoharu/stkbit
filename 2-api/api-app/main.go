package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"movie-backend/api-app/config"
	"movie-backend/api-app/movie/controller/http"
	"movie-backend/api-app/movie/repository/rpc"
	"movie-backend/api-app/movie/usecase"
	"movie-backend/movie-app/movie/controller/rpc/pb"
	"os"
)

var r *gin.Engine
var appConfig config.AppConfig
var env string

func init() {
	r = gin.Default()

	// Initialize environment
	env = os.Getenv("STOCKBIT_ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigFile(fmt.Sprintf("api-app/config/%s.json", env))
	_ = viper.ReadInConfig()
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	// Initialize Movie Service
	movieServiceConn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", appConfig.MovieServiceClient.Host, appConfig.MovieServiceClient.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("can't connect: %s", err)
	}
	movieServiceClient := pb.NewMovieServiceClient(movieServiceConn)
	movieRepo := rpc.NewMovieRPCRepository(movieServiceClient)
	movieUC := usecase.NewMovieUseCase(movieRepo)
	http.CreateMovieController(r.Group("/v1"), movieUC)
}

func main() {
	port := os.Getenv("PORT")
	var address string

	if port == "" {
		address = appConfig.Addr
	} else {
		address = fmt.Sprintf(":%v", port)
	}
	r.Run(address)
}
