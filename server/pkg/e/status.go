package e

import "net/http"

var statusFlags = map[int]int{
	SUCCESS:        http.StatusOK,
	ERROR:          http.StatusInternalServerError,
	INVALID_PARAMS: http.StatusBadRequest,
	FORBIDDEN:      http.StatusForbidden,
}

// GetStatus get status code information based on Code
func GetStatus(code int) int {
	msg, ok := statusFlags[code]
	if ok {
		return msg
	}

	return statusFlags[ERROR]
}
