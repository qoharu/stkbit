package rpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"movie-backend/movie-app/movie"
	"movie-backend/movie-app/movie/controller/rpc/pb"
	"movie-backend/movie-app/movie/repository/mock_repository"
	"movie-backend/movie-app/movie/usecase"
	"reflect"
	"testing"
)

func Test_gRPCServer_SearchMovie(t *testing.T) {
	// Prepare mock data
	var repoMovies []movie.Movie
	repoMovies = append(repoMovies, movie.Movie{
		Title:  "Midnight Dinner",
		Year:   "2020",
		ImdbID: "ti19120",
		Type:   "series",
		Poster: "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg",
	})
	repoResult := movie.MovieSearchResponse{
		Search:       repoMovies,
		TotalResults: "1",
		Response:     "True",
	}

	var movies []*pb.MovieDisplay
	movies = append(movies, &pb.MovieDisplay{
		Title:  repoMovies[0].Title,
		Year:   repoMovies[0].Year,
		ImdbID: repoMovies[0].ImdbID,
		Type:   repoMovies[0].Type,
		Poster: repoMovies[0].Poster,
	})

	expectedResult := pb.MovieSearchResponse{
		Search:       movies,
		TotalResults: 1,
		Response:     true,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)
	mockRepo.EXPECT().Search(movie.MovieSearchSpec{
		Keyword: "Dinner",
		Page:    1,
	}).Return(repoResult, nil)

	type fields struct {
		useCase movie.UseCase
	}
	type args struct {
		ctx  context.Context
		spec *pb.MovieSearchSpec
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.MovieSearchResponse
		wantErr bool
	}{
		{
			name:   "search success",
			fields: fields{useCase: usecase.NewMovieUseCase(mockRepo)},
			args: args{
				ctx: nil,
				spec: &pb.MovieSearchSpec{
					Keyword: "Dinner",
					Page:    1,
				},
			},
			want:    &expectedResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gRPCServer{
				useCase: tt.fields.useCase,
			}
			got, err := g.SearchMovie(tt.args.ctx, tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}
