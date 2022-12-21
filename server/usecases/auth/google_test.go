package auth

import (
	"context"
	"math/rand"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/assert"
)

func TestUseCase_GetByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		arg       args
		mockRepo  func(a args, mock MockIUserRepo)
		want      user.User
		wantError bool
	}{{
		name: "success",
		arg: args{
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{{
					ID:    10,
					Name:  "bob",
					Email: "bob@gmail.co",
				}}, 0, nil)
		},
		want: user.User{
			ID:    10,
			Name:  "bob",
			Email: "bob@gmail.co",
		},
		wantError: false,
	}, {
		name: "error_empty_users",
		arg: args{
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
		},
		want:      user.User{},
		wantError: true,
	}, {
		name: "error_get_repo",
		arg: args{
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, assert.ErrMock)
		},
		want:      user.User{},
		wantError: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := NewMockIUserRepo(ctrl)

			u := New(Dependencies{
				UserRepo: mockUserRepo,
			})

			tt.mockRepo(tt.arg, *mockUserRepo)

			got, err := u.GetByEmail(ctx, tt.arg.email)
			assert.Equal(t, tt.want, got)
			assert.WantError(t, tt.wantError, err)
		})
	}
}

func TestUseCase_Signup(t *testing.T) {
	mockTime := time.Now()

	type args struct {
		tag      string
		name     string
		email    string
		randSeed int64
	}
	tests := []struct {
		name      string
		arg       args
		mockRepo  func(a args, mock MockIUserRepo)
		want      user.User
		wantError bool
	}{{
		name: "success",
		arg: args{
			tag:   "user-123",
			name:  "bob",
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)

			mock.EXPECT().Create(gomock.Any(), user.User{
				Tag:   a.tag,
				Name:  a.name,
				Email: a.email,
				HPass: "",
			}).Return(user.User{
				ID:    1,
				Tag:   a.tag,
				Name:  a.name,
				Email: a.email,
			}, nil)
		},
		want: user.User{
			ID:        1,
			Tag:       "user-123",
			Name:      "bob",
			Email:     "bob@gmail.co",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
		},
		wantError: false,
	}, {
		name: "success_generate_tag",
		arg: args{
			tag:      "",
			name:     "bob",
			email:    "bob@gmail.co",
			randSeed: 1234567,
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil).Times(2)

			mock.EXPECT().Create(gomock.Any(), user.User{
				Tag:   "kRFlQxwyCmRfWjMyF",
				Name:  a.name,
				Email: a.email,
				HPass: "",
			}).Return(user.User{
				ID:    1,
				Tag:   "kRFlQxwyCmRfWjMyF",
				Name:  a.name,
				Email: a.email,
			}, nil)
		},
		want: user.User{
			ID:        1,
			Tag:       "kRFlQxwyCmRfWjMyF",
			Name:      "bob",
			Email:     "bob@gmail.co",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
		},
		wantError: false,
	}, {
		name: "error_user_already_exist",
		arg: args{
			tag:   "tag",
			name:  "bob",
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 1, nil)
		},
		want: user.User{
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
		},
		wantError: true,
	}, {
		name: "error_repo_create",
		arg: args{
			tag:   "tag",
			name:  "bob",
			email: "bob@gmail.co",
		},
		mockRepo: func(a args, mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
			mock.EXPECT().Create(gomock.Any(), gomock.Any()).
				Return(user.User{}, assert.ErrMock)
		},
		want: user.User{
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
		},
		wantError: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rand.Seed(tt.arg.randSeed)

			mockUserRepo := NewMockIUserRepo(ctrl)

			u := New(Dependencies{
				UserRepo: mockUserRepo,
			})

			tt.mockRepo(tt.arg, *mockUserRepo)

			got, err := u.Signup(ctx, tt.arg.tag, tt.arg.name, tt.arg.email)
			got.CreatedAt = mockTime
			got.UpdatedAt = mockTime

			assert.Equal(t, tt.want, got)
			assert.WantError(t, tt.wantError, err)
		})
	}
}

func TestUseCase_GenerateTag(t *testing.T) {
	tests := []struct {
		name     string
		randSeed int64
		want     string
		mockRepo func(mock MockIUserRepo)
	}{{
		name:     "success",
		randSeed: 12345,
		want:     "sALeNEWaRfnjAArNq",
		mockRepo: func(mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
		},
	}, {
		name:     "success_retries_3",
		randSeed: 123,
		want:     "bTfQtvYZwsXmCScap",
		mockRepo: func(mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 1, nil).Times(2)
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, assert.ErrMock)
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 0, nil)
		},
	}, {
		name:     "success_retries_exhausted",
		randSeed: 1234567,
		want:     "QNnWqatvKqiEvOrZdiCE",
		mockRepo: func(mock MockIUserRepo) {
			mock.EXPECT().Get(gomock.Any(), gomock.Any()).
				Return([]user.User{}, 1, nil).AnyTimes()
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rand.Seed(tt.randSeed)

			mockUserRepo := NewMockIUserRepo(ctrl)

			u := New(Dependencies{
				UserRepo: mockUserRepo,
			})

			tt.mockRepo(*mockUserRepo)

			got := u.GenerateTag(ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}
