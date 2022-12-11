package util

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/kevin-luvian/goauth/server/pkg/logging"
)

func ComputeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s += len(r.URL.String())
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}

func GetClientIPAddr(req *http.Request) string {
	ipSlice := []string{
		req.Header.Get("X-FORWARDED-FOR"),
		req.Header.Get("X-Forwarded-For"),
		req.Header.Get("x-forwarded-for"),
	}

	for _, ip := range ipSlice {
		if ip != "" {
			return ip
		}
	}

	return strings.Split(req.RemoteAddr, ":")[0]
}

func CreateCORSRule(urls []string) []*regexp.Regexp {
	matcher := []*regexp.Regexp{}
	for _, c := range urls {
		if c == "" {
			continue
		}

		reg, err := regexp.Compile(c)
		if err != nil {
			logging.Errorf("Error compiling cors url=%s, err=%v", c, err)
			continue
		}
		matcher = append(matcher, reg)
	}
	return matcher
}

func CheckOrigin(match []*regexp.Regexp) func(origin string) bool {
	return func(origin string) bool {
		for _, m := range match {
			ok := m.MatchString(origin)
			if ok || origin == "*" {
				return ok
			}
		}
		return false
	}
}
