package middleware

import (
	"gateway-go/auth"
	"net/http"
	"strings"
)

func InterceptJWT(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		token, err := retrieveTokenFromRequest(request)

		if err != nil {
			writeErrorToResponse(writer, err)
		}

		valid, err := auth.VerifyWithSecret(token)

		if err != nil || !valid {
			writeErrorToResponse(writer, err)
			return
		}

		h.ServeHTTP(writer, request)
	})
}

func writeErrorToResponse(writer http.ResponseWriter, err error)  {

	writer.WriteHeader(http.StatusForbidden)

	writer.Header().Add("Content-Type", "application/json")

	writer.Write([]byte(err.Error()))
}

type WrongTokenError struct {}

func (e *WrongTokenError) Error() string{
	return "Empty token"
}

func newWrongTokenError() *WrongTokenError {
	return &WrongTokenError{}
}

func retrieveTokenFromRequest(request *http.Request) (string, error) {

	authorization := request.Header.Get("Authorization")

	if authorization == "" {
		return authorization, newWrongTokenError()
	}

	splitted := strings.Split(authorization, " ")

	if len(splitted) != 2 {
		return authorization, newWrongTokenError()
	}

	return splitted[1], nil
}
