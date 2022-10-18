package auth

import (
	"context"
	"fmt"
	"net/http"
)

type CtxKey string

const CtxReqIdKey CtxKey = "x-api-key"

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		CtxKey := req.Header.Get("x-api-key")
		if CtxKey == "" {
			CtxKey = "NONE"
			fmt.Fprintf(res, "Header not found\n")
		}

		ctx1 := context.WithValue(req.Context(), CtxReqIdKey, CtxKey)
		ctx2 := req.WithContext(ctx1)

		handler.ServeHTTP(res, ctx2)
	})
}

func Yo() string {
	return "Yo!"
}
