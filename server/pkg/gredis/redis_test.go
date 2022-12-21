package gredis

import (
	"testing"
	"time"

	"github.com/kevin-luvian/goauth/server/pkg/assert"
	"github.com/rafaeljusto/redigomock"
)

func TestRedis_Setup(t *testing.T) {
	// invalid settings panic
	assert.Panic(t, func() {
		Setup()
	})
}

func TestRedis_Ping(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(conn *redigomock.Conn)
		wantErr bool
	}{{
		name: "test_success",
		mock: func(conn *redigomock.Conn) {
			MockExpectPing(conn, nil)
		},
		wantErr: false,
	}, {
		name: "test_error",
		mock: func(conn *redigomock.Conn) {
			MockExpectPing(conn, assert.ErrMock)
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MockRedis()
			tt.mock(conn)

			err := Ping()
			assert.WantError(t, tt.wantErr, err)
		})
	}
}

func TestRedis_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		arg     args
		mock    func(conn *redigomock.Conn, a args)
		want    string
		wantErr bool
	}{{
		name: "test_success",
		arg: args{
			key: "key1",
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectGet(conn, a.key, "val1", nil)
		},
		want:    "val1",
		wantErr: false,
	}, {
		name: "test_error",
		arg: args{
			key: "key1",
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectGet(conn, a.key, "", assert.ErrMock)
		},
		want:    "",
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MockRedis()
			tt.mock(conn, tt.arg)

			got, err := Get(tt.arg.key)
			assert.WantError(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedis_GetStruct(t *testing.T) {
	type mockType struct {
		Name string
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		arg     args
		mock    func(conn *redigomock.Conn, a args)
		want    mockType
		wantErr bool
	}{{
		name: "test_success",
		arg: args{
			key: "key1",
		},
		mock: func(conn *redigomock.Conn, a args) {
			target := mockType{
				Name: "test123",
			}
			MockExpectGetStruct(conn, a.key, target, nil)
		},
		want: mockType{
			Name: "test123",
		},
		wantErr: false,
	}, {
		name: "test_error",
		arg: args{
			key: "key1",
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectGetStruct(conn, a.key, nil, assert.ErrMock)
		},
		want:    mockType{},
		wantErr: true,
	}, {
		name: "test_error_type",
		arg: args{
			key: "key1",
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectGetStruct(conn, a.key, "ntype", nil)
		},
		want:    mockType{},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MockRedis()
			tt.mock(conn, tt.arg)

			target := mockType{}
			err := GetStruct(tt.arg.key, &target)

			assert.WantError(t, tt.wantErr, err)
			assert.Equal(t, tt.want, target)
		})
	}
}

func TestRedis_Set(t *testing.T) {
	mockTTL := 15 * time.Minute

	type args struct {
		key   string
		value string
		ttl   *time.Duration
	}

	tests := []struct {
		name    string
		arg     args
		mock    func(conn *redigomock.Conn, a args)
		wantErr bool
	}{{
		name: "test_success",
		arg: args{
			key:   "key1",
			value: "val1",
			ttl:   &mockTTL,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSet(conn, a.key, a.value, mockTTL, nil)
		},
		wantErr: false,
	}, {
		name: "test_success_default",
		arg: args{
			key:   "key1",
			value: "val1",
			ttl:   nil,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSet(conn, a.key, a.value, DefaultTTL, nil)
		},
		wantErr: false,
	}, {
		name: "test_error",
		arg: args{
			key:   "key1",
			value: "val1",
			ttl:   nil,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSet(conn, a.key, a.value, DefaultTTL, assert.ErrMock)
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MockRedis()
			tt.mock(conn, tt.arg)

			var err error
			if tt.arg.ttl != nil {
				err = Set(tt.arg.key, tt.arg.value, *tt.arg.ttl)
			} else {
				err = Set(tt.arg.key, tt.arg.value)
			}

			assert.WantError(t, tt.wantErr, err)
		})
	}
}

func TestRedis_SetStruct(t *testing.T) {
	mockTTL := 15 * time.Minute

	type args struct {
		key   string
		value interface{}
		ttl   *time.Duration
	}

	tests := []struct {
		name    string
		arg     args
		mock    func(conn *redigomock.Conn, a args)
		wantErr bool
	}{{
		name: "test_success",
		arg: args{
			key: "key1",
			value: struct {
				Name string
			}{
				Name: "abc",
			},
			ttl: &mockTTL,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSetStruct(conn, a.key, a.value, mockTTL, nil)
		},
		wantErr: false,
	}, {
		name: "test_success_default",
		arg: args{
			key: "key1",
			value: struct {
				Name string
			}{
				Name: "abc",
			},
			ttl: nil,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSetStruct(conn, a.key, a.value, DefaultTTL, nil)
		},
		wantErr: false,
	}, {
		name: "test_error",
		arg: args{
			key: "key1",
			value: struct {
				Name string
			}{
				Name: "abc",
			},
			ttl: nil,
		},
		mock: func(conn *redigomock.Conn, a args) {
			MockExpectSetStruct(conn, a.key, a.value, DefaultTTL, assert.ErrMock)
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MockRedis()
			tt.mock(conn, tt.arg)

			var err error
			if tt.arg.ttl != nil {
				err = SetStruct(tt.arg.key, tt.arg.value, *tt.arg.ttl)
			} else {
				err = SetStruct(tt.arg.key, tt.arg.value)
			}

			assert.WantError(t, tt.wantErr, err)
		})
	}
}
