package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"movie-backend/movie-app/movie"
	"movie-backend/movie-app/movie/repository"
	"movie-backend/movie-app/movie/repository/mock_repository"
	"reflect"
	"testing"
)

func Test_movieUseCase_SearchMovie(t *testing.T) {
	// Prepare mock data
	var movies []movie.Movie
	movies = append(movies, movie.Movie{
		Title:  "Midnight Dinner",
		Year:   "2020",
		ImdbID: "ti19120",
		Type:   "series",
		Poster: "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg",
	})
	expectedResult := movie.MovieSearchResponse{
		Search:       movies,
		TotalResults: "1",
		Response:     "True",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)
	mockRepo.EXPECT().Search(movie.MovieSearchSpec{
		Keyword: "Dinner",
		Page:    1,
	}).Return(expectedResult, nil)

	mockRepo.EXPECT().Search(movie.MovieSearchSpec{
		Keyword: "ToManyWordsThatIsNotFound",
		Page:    1,
	}).Return(movie.MovieSearchResponse{}, errors.New("connection error"))

	type fields struct {
		movieRepo repository.MovieRepository
	}
	type args struct {
		spec movie.MovieSearchSpec
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes movie.MovieSearchResponse
		wantErr bool
	}{
		{
			name:   "search success",
			fields: fields{movieRepo: mockRepo},
			args: args{spec: movie.MovieSearchSpec{
				Keyword: "Dinner",
				Page:    1,
			}},
			wantRes: expectedResult,
			wantErr: false,
		},
		{
			name:   "search error",
			fields: fields{movieRepo: mockRepo},
			args: args{spec: movie.MovieSearchSpec{
				Keyword: "ToManyWordsThatIsNotFound",
				Page:    1,
			}},
			wantRes: movie.MovieSearchResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewMovieUseCase(tt.fields.movieRepo)
			gotRes, err := useCase.SearchMovie(tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SearchMovie() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
