package auth

import (
	"context"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/pkg/assert"
	"golang.org/x/oauth2"
	"gopkg.in/h2non/gock.v1"
)

func TestRepo_GetGoogleUserInfo(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockIOAuth)
		want    google.GoogleUserInfo
		wantErr bool
	}{{
		name: "success",
		args: args{
			code: "404",
		},
		mock: func(a args, mock *MockIOAuth) {
			rtok := &oauth2.Token{}
			mock.EXPECT().Exchange(gomock.Any(), a.code).Return(rtok, nil)
			mock.EXPECT().Client(gomock.Any(), rtok).Return(&http.Client{})
			gock.New("https://www.googleapis.com/oauth2/v2/userinfo").
				Reply(200).BodyString(`{"name":"roger"}`)
		},
		want: google.GoogleUserInfo{
			Name: "roger",
		},
		wantErr: false,
	}, {
		name: "error_exchange",
		args: args{
			code: "404",
		},
		mock: func(a args, mock *MockIOAuth) {
			mock.EXPECT().Exchange(gomock.Any(), a.code).Return(&oauth2.Token{}, assert.ErrMock)
		},
		want:    google.GoogleUserInfo{},
		wantErr: true,
	}, {
		name: "error_getting_userinfo",
		args: args{
			code: "404",
		},
		mock: func(a args, mock *MockIOAuth) {
			rtok := &oauth2.Token{}
			mock.EXPECT().Exchange(gomock.Any(), a.code).Return(rtok, nil)
			mock.EXPECT().Client(gomock.Any(), rtok).Return(&http.Client{})
			gock.New("https://www.googleapis.com/oauth2/v2/userinfo").ReplyError(assert.ErrMock)
		},
		want:    google.GoogleUserInfo{},
		wantErr: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOAuth := NewMockIOAuth(ctrl)

			r := New(mockOAuth)

			tt.mock(tt.args, mockOAuth)

			got, err := r.GetGoogleUserInfo(ctx, tt.args.code)
			assert.Equal(t, tt.want, got)
			assert.WantError(t, tt.wantErr, err)
		})
	}
}

func TestRepo_GoogleRedirectURL(t *testing.T) {
	type args struct {
		ctx   context.Context
		state string
	}
	tests := []struct {
		name string
		args args
		mock func(a args, mock *MockIOAuth)
		want string
	}{
		{
			name: "Test Success",
			args: args{
				ctx:   context.Background(),
				state: "state",
			},
			mock: func(a args, mock *MockIOAuth) {
				mock.EXPECT().AuthCodeURL(a.state).Return("http://authurl")
			},
			want: "http://authurl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOAuth := NewMockIOAuth(ctrl)

			r := New(&oauth2.Config{})

			r.oauth = mockOAuth

			tt.mock(tt.args, mockOAuth)

			if got := r.GoogleRedirectURL(tt.args.ctx, tt.args.state); got != tt.want {
				t.Errorf("GoogleRedirectURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
