package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"movie-backend/api-app/movie/controller/http"
	"movie-backend/api-app/movie/repository/rpc"
	"movie-backend/api-app/movie/usecase"
	"movie-backend/movie-app/movie/controller/rpc/pb"
	"os"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	// Initialize Movie Service
	movieServiceConn, err := grpc.Dial(":9080", grpc.WithInsecure())
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
		address = ":55501"
	} else {
		address = fmt.Sprintf(":%v", port)
	}
	r.Run(address)
}
