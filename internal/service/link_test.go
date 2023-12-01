package service

import (
	"Ozon/domain"
	mock_service "Ozon/internal/service/mocks"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateShortLink(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRepository, fullLink, shortLink string)
	type args struct {
		FullLink  string
		ShortLink string
	}
	testTable := []struct {
		name string
		args
		mockBehavior
		want    string
		wantErr error
	}{
		{
			name: "ok",
			args: args{FullLink: "ozon.ru", ShortLink: "Lw1XBy9jH5"},
			mockBehavior: func(s *mock_service.MockRepository, fullLink, shortLink string) {
				s.EXPECT().CreateShortLink(context.Background(), fullLink, shortLink).Return(nil)
			},
			want:    "Lw1XBy9jH5",
			wantErr: nil,
		},
		{
			name:         "invalid format",
			args:         args{FullLink: "ozonru", ShortLink: ""},
			mockBehavior: func(r *mock_service.MockRepository, shortUrl, originalUrl string) {},
			want:         "",
			wantErr:      domain.ErrInvalidArgument,
		},
		{
			name:         "invalid format",
			args:         args{FullLink: "", ShortLink: ""},
			mockBehavior: func(r *mock_service.MockRepository, shortUrl, originalUrl string) {},
			want:         "",
			wantErr:      domain.ErrInvalidArgument,
		},
		{
			name:         "Bad request",
			args:         args{FullLink: "", ShortLink: "Lw1XBy9jH5"},
			mockBehavior: func(r *mock_service.MockRepository, shortUrl, originalUrl string) {},
			want:         "",
			wantErr:      domain.ErrInvalidArgument,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			shortL := mock_service.NewMockRepository(c)
			testCase.mockBehavior(shortL, testCase.FullLink, testCase.ShortLink)

			service := NewService(shortL)
			response, err := service.CreateShortLink(context.Background(), testCase.args.FullLink)
			assert.Equal(t, testCase.wantErr, err)
			assert.Equal(t, testCase.want, response)
		})
	}
}

func TestService_GetFullLink(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRepository, shortLink string)
	testTable := []struct {
		name      string
		ShortLink string
		mockBehavior
		want    string
		wantErr error
	}{
		{
			name:      "ok",
			ShortLink: "Lw1XBy9jH5",
			mockBehavior: func(s *mock_service.MockRepository, shortLink string) {
				s.EXPECT().GetFullLink(context.Background(), shortLink).Return("ozon.ru", nil)
			},
			want:    "ozon.ru",
			wantErr: nil,
		},
		{
			name:         "invalid argument",
			ShortLink:    "Lw1XBy9jH",
			mockBehavior: func(s *mock_service.MockRepository, shortLink string) {},
			want:         "",
			wantErr:      domain.ErrInvalidArgument,
		},
		{
			name:         "invalid argument",
			ShortLink:    "",
			mockBehavior: func(s *mock_service.MockRepository, shortLink string) {},
			want:         "",
			wantErr:      domain.ErrInvalidArgument,
		},
		{
			name:         "not found",
			ShortLink:    "Lw1XBy9jH2",
			mockBehavior: func(s *mock_service.MockRepository, shortLink string) {},
			want:         "",
			wantErr:      domain.ErrNoRecordFound,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			shortL := mock_service.NewMockRepository(c)
			testCase.mockBehavior(shortL, testCase.ShortLink)

			service := NewService(shortL)
			response, err := service.GetFullLink(context.Background(), testCase.ShortLink)
			assert.Equal(t, testCase.wantErr, err)
			assert.Equal(t, testCase.want, response)
		})
	}
}

func TestService_MakeShortLink(t *testing.T) {
	testTable := []struct {
		name      string
		shortLink string
		want      string
	}{
		{
			name:      "ok",
			shortLink: "ozon.ru",
			want:      "Lw1XBy9jH5",
		},
		{
			name:      "ok",
			shortLink: "ju.st",
			want:      "1tfEf1peQA",
		},
		{
			name:      "ok",
			shortLink: "do.it",
			want:      "A75VuzJHot",
		},
		{
			name:      "ok",
			shortLink: "want.sleep",
			want:      "FvAVQ7uIOA",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := mock_service.NewMockRepository(ctrl)
			service := NewService(s)
			response := service.MakeShortLink(testCase.shortLink)
			assert.Equal(t, testCase.want, response)
		})
	}
}
