package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"movie-backend/api-app/movie"
	"movie-backend/api-app/movie/repository"
	"movie-backend/api-app/movie/repository/mock_repository"
	"movie-backend/movie-app/movie/controller/rpc/pb"
	"reflect"
	"testing"
)

func Test_movieUseCase_Search(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepoResponse := pb.MovieSearchResponse{
		Search:       nil,
		TotalResults: 0,
		Response:     true,
	}

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)
	mockRepo.EXPECT().Search(nil, "Dinner", int64(100)).Return(mockRepoResponse, nil)

	expectedResult := movie.MovieListResponse{
		Movies:    nil,
		TotalData: 0,
	}

	type fields struct {
		movieRepo repository.MovieRepository
	}
	type args struct {
		ctx     context.Context
		keyword string
		page    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes movie.MovieListResponse
		wantErr bool
	}{
		{
			name:   "success search with empty result",
			fields: fields{movieRepo: mockRepo},
			args: args{
				ctx:     nil,
				keyword: "Dinner",
				page:    100,
			},
			wantRes: expectedResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			muc := NewMovieUseCase(tt.fields.movieRepo)
			gotRes, err := muc.Search(tt.args.ctx, tt.args.keyword, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Search() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
