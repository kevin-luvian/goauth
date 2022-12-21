package util

import (
	"math/rand"
	"reflect"
	"time"
)

func Setup() {
	rand.Seed(time.Now().UnixNano())
}

type CompareFunc func(i int, first interface{}, second interface{}) bool

func RandString(size int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}

func Contains(in interface{}, value interface{}, comparators ...CompareFunc) bool {
	if len(comparators) == 0 {
		comparators = []CompareFunc{func(_ int, first, second interface{}) bool {
			return first == second
		}}
	}

	comparator := comparators[0]

	if inV := reflect.ValueOf(in); inV.IsValid() {
		for i := 0; i < inV.Len(); i++ {
			if current := inV.Index(i); current.IsValid() {
				if comparator(i, value, current.Interface()) {
					return true
				}
			}
		}
	}

	return false
}
