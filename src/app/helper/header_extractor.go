package helper

import (
	"errors"
	"net/http"
	"strings"
)

func HeaderExtractor(key string, r *http.Request) (string, error) {
	header := r.Header.Get(key)
	if header == "" {
		return "", errors.New("authorization token header not found")
	}

	switch key {
	case "Authorization":
		authHeaderParts := strings.Fields(header)
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return "", errors.New("your bearer token not found")
		}
		return authHeaderParts[1], nil
	default:
		return header, nil
	}
}
