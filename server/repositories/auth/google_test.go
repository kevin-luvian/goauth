package auth

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"golang.org/x/oauth2"
)

// func TestRepo_GetGoogleUserInfo(t *testing.T) {
// 	type fields struct {
// 		oauth *oauth2.Config
// 	}
// 	type args struct {
// 		ctx  context.Context
// 		code string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		mock    func(a args, mock *MockoauthInterface)
// 		want    user.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "Test Success",
// 			fields: fields{
// 				oauth: &oauth2.Config{
// 					ClientID:     "client",
// 					ClientSecret: "secret",
// 					RedirectURL:  "http://localhost:8181/test",
// 					Endpoint:     google.Endpoint,
// 				},
// 				db: nil,
// 			},
// 			args: args{},
// 			mock: func(a args, mock *MockoauthInterface) {
// 				mock.EXPECT().Exchange(a.ctx, a.code).Return(&oauth2.Token{}, nil)
// 				mock.EXPECT().Client(a.ctx, gomock.Any()).Return(&http.Client{})
// 				gock.New("https://www.googleapis.com/oauth2/v2/userinfo").
// 					Reply(200).BodyString("{}")
// 			},
// 			want:    user.User{},
// 			wantErr: false,
// 		}, {
// 			name: "Test Error - Exchange Token",
// 			fields: fields{
// 				oauth: &oauth2.Config{
// 					ClientID:     "client",
// 					ClientSecret: "secret",
// 					RedirectURL:  "http://localhost:8181/test",
// 					Endpoint:     google.Endpoint,
// 				},
// 				db: nil,
// 			},
// 			args: args{},
// 			mock: func(a args, mock *MockoauthInterface) {
// 				mock.EXPECT().Exchange(a.ctx, a.code).
// 					Return(&oauth2.Token{}, errors.New("failed exchange token"))
// 			},
// 			want:    user.User{},
// 			wantErr: true,
// 		}, {
// 			name: "Test Error - Parsing Response",
// 			fields: fields{
// 				oauth: &oauth2.Config{
// 					ClientID:     "client",
// 					ClientSecret: "secret",
// 					RedirectURL:  "http://localhost:8181/test",
// 					Endpoint:     google.Endpoint,
// 				},
// 				db: nil,
// 			},
// 			args: args{},
// 			mock: func(a args, mock *MockoauthInterface) {
// 				mock.EXPECT().Exchange(a.ctx, a.code).Return(&oauth2.Token{}, nil)
// 				mock.EXPECT().Client(a.ctx, gomock.Any()).Return(&http.Client{})
// 				gock.New("https://www.googleapis.com/oauth2/v2/userinfo").
// 					Reply(200).BodyString("{")
// 			},
// 			want:    user.User{},
// 			wantErr: true,
// 		}, {
// 			name: "Test Error - Request Failed",
// 			fields: fields{
// 				oauth: &oauth2.Config{
// 					ClientID:     "client",
// 					ClientSecret: "secret",
// 					RedirectURL:  "http://localhost:8181/test",
// 					Endpoint:     google.Endpoint,
// 				},
// 				db: nil,
// 			},
// 			args: args{},
// 			mock: func(a args, mock *MockoauthInterface) {
// 				mock.EXPECT().Exchange(a.ctx, a.code).Return(&oauth2.Token{}, nil)
// 				mock.EXPECT().Client(a.ctx, gomock.Any()).Return(&http.Client{})
// 				gock.New("https://www.googleapis.com/oauth2/v2/userinfo").
// 					Reply(200).SetError(errors.New("some error"))
// 			},
// 			want:    user.User{},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mockOAuth := NewMockoauthInterface(ctrl)

// 			r := New(tt.fields.db, tt.fields.oauth)
// 			r.oauth = mockOAuth
// 			tt.mock(tt.args, mockOAuth)
// 			got, err := r.GetGoogleUserInfo(tt.args.ctx, tt.args.code)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetGoogleUserInfo() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetGoogleUserInfo() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestRepo_GoogleLoginURL(t *testing.T) {
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

			if got := r.GoogleLoginURL(tt.args.ctx, tt.args.state); got != tt.want {
				t.Errorf("GoogleLoginURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
