package util

import (
	"testing"

	"github.com/kevin-luvian/goauth/server/pkg/assert"
)

func TestUtil_EncodeGOB(t *testing.T) {
	testCases := []struct {
		name string
		arg  interface{}
		want []uint8
	}{{
		name: "success",
		arg:  "abcdefghijklmnopqrstuvwxyz",
		want: []uint8([]byte{0x1d, 0xc, 0x0, 0x1a, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a}),
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeGob(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUtil_EncodeMD5(t *testing.T) {
	testCases := []struct {
		name string
		arg  interface{}
		want string
	}{{
		name: "success",
		arg:  "abcdefghijklmnopqrstuvwxyz",
		want: "8d1d0ba2bc9670b2f4b7ddf60d034425",
	}, {
		name: "success_int",
		arg:  120,
		want: "4b4ac996ccda482227dad331279aae74",
	}, {
		name: "success_struct",
		arg: struct {
			name string
		}{"name"},
		want: "b5d104f542a4c833f27e61f4cd9c68b9",
	}, {
		name: "success_array",
		arg:  []string{"a", "b", "c", "d", "e"},
		want: "783861278069d483430a55b3fa86e1ee",
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeMD5(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
