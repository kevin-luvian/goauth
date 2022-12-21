package auth

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	gomock "github.com/golang/mock/gomock"
	"github.com/kevin-luvian/goauth/server/entity/token"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/assert"
)

func TestUseCase_SignJWT(t *testing.T) {
	accessExpiry := 10 * time.Minute
	refreshExpiry := 10 * time.Hour

	type wants struct {
		accessToken  token.AccessToken
		refreshToken token.RefreshToken
	}
	type args struct {
		usr user.User
	}

	tests := []struct {
		name    string
		args    args
		want    wants
		wantErr bool
	}{{
		name: "success",
		args: args{
			usr: user.User{
				ID:   90,
				Name: "bob",
			},
		},
		want: wants{
			accessToken: token.AccessToken{
				UserID:   90,
				UserName: "bob",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpiry)),
				},
			},
			refreshToken: token.RefreshToken{
				UserID: 90,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpiry)),
				},
			},
		},
		wantErr: false,
	}, {
		name: "empty_user",
		args: args{
			usr: user.User{},
		},
		want: wants{
			accessToken: token.AccessToken{
				UserID:   0,
				UserName: "",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpiry)),
				},
			},
			refreshToken: token.RefreshToken{
				UserID: 0,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpiry)),
				},
			},
		},
		wantErr: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := New(Dependencies{})

			gotAccessToken, gotRefreshToken, err := r.SignJWT(ctx, tt.args.usr)

			// if want error, stop comparing others
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			aTokenClaims := token.AccessToken{}
			tkn, _ := jwt.ParseWithClaims(gotAccessToken, &aTokenClaims, func(token *jwt.Token) (interface{}, error) {
				return []byte(""), nil
			})
			assert.True(t, tkn.Valid)
			assert.Equal(t, tt.want.accessToken, aTokenClaims)

			rTokenClaims := token.RefreshToken{}
			tkn, _ = jwt.ParseWithClaims(gotRefreshToken, &rTokenClaims, func(token *jwt.Token) (interface{}, error) {
				return []byte(""), nil
			})
			assert.True(t, tkn.Valid)
			assert.Equal(t, tt.want.refreshToken, rTokenClaims)
		})
	}
}

func TestUseCase_ParseJWTAccessToken(t *testing.T) {
	type args struct {
		accessToken token.AccessToken
		secret      string
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(a args, mock *MockIUserRepo)
		want     user.User
		wantErr  bool
	}{{
		name: "success",
		args: args{
			accessToken: token.AccessToken{
				UserID:   10,
				UserName: "bob",
			},
			secret: "secretz",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{{
					ID:   a.accessToken.UserID,
					Name: "bobby",
				}}, 0, nil)
		},
		want: user.User{
			ID:   10,
			Name: "bobby",
		},
		wantErr: false,
	}, {
		name: "error_no_user_found",
		args: args{
			accessToken: token.AccessToken{
				UserID:   10,
				UserName: "bob",
			},
			secret: "secretz",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
		},
		want:    user.User{},
		wantErr: true,
	}, {
		name: "error_invalid_secret",
		args: args{
			accessToken: token.AccessToken{
				UserID:   10,
				UserName: "bob",
			},
			secret: "invalid",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {},
		want:     user.User{},
		wantErr:  true,
	}, {
		name: "error_invalid_secret_empty",
		args: args{
			accessToken: token.AccessToken{
				UserID:   10,
				UserName: "bob",
			},
			secret: "",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {},
		want:     user.User{},
		wantErr:  true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockIUserRepo(ctrl)

			r := New(Dependencies{
				UserRepo: mockRepo,
			})
			r.secrets.access = "secretz"

			tt.mockRepo(tt.args, mockRepo)

			jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tt.args.accessToken)
			accessToken, _ := jwtToken.SignedString([]byte(tt.args.secret))

			got, err := r.ParseJWTAccessToken(ctx, accessToken)
			assert.Equal(t, tt.want, got)
			assert.WantError(t, tt.wantErr, err)
		})
	}
}

func TestUseCase_ParseJWTRefreshToken(t *testing.T) {
	type args struct {
		refreshToken token.RefreshToken
		secret       string
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(a args, mock *MockIUserRepo)
		want     user.User
		wantErr  bool
	}{{
		name: "success",
		args: args{
			refreshToken: token.RefreshToken{
				UserID: 10,
			},
			secret: "secretz",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{{
					ID: 10,
				}}, 0, nil)
		},
		want: user.User{
			ID: 10,
		},
		wantErr: false,
	}, {
		name: "error_no_user_found",
		args: args{
			refreshToken: token.RefreshToken{
				UserID: 10,
			},
			secret: "secretz",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
		},
		wantErr: true,
	}, {
		name: "error_invalid_secret",
		args: args{
			refreshToken: token.RefreshToken{
				UserID: 10,
			},
			secret: "invalid",
		},
		mockRepo: func(a args, mock *MockIUserRepo) {},
		wantErr:  true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockIUserRepo(ctrl)

			r := New(Dependencies{
				UserRepo: mockRepo,
			})
			r.secrets.refresh = "secretz"

			tt.mockRepo(tt.args, mockRepo)

			jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tt.args.refreshToken)
			accessToken, _ := jwtToken.SignedString([]byte(tt.args.secret))

			got, err := r.ParseJWTRefreshToken(ctx, accessToken)
			assert.Equal(t, tt.want, got)
			assert.WantError(t, tt.wantErr, err)
		})
	}
}
