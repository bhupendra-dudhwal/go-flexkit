package middleware

import "github.com/google/uuid"

func getRequestID() string {
	return uuid.NewString()
}
