package main

import (
	"context"
	"fmt"
	"net/http"
)

const requestIdKey = "rid"

func main() {

	http.Handle("/", middleWare(http.HandlerFunc(handler)))

	http.ListenAndServe(":9800", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	reqId := requestIdFromContext(r.Context())
	fmt.Fprintln(w, "Request Id: ", reqId)
	return
}

func requestIdFromContext(ctx context.Context) string {
	return ctx.Value(requestIdKey).(string)
}
func newContextWithRequestId(ctx context.Context, r *http.Request) context.Context {
	reqId := r.Header.Get("X-Request-Id")
	if reqId == "" {
		reqId = "0"
	}
	return context.WithValue(ctx, requestIdKey, reqId)
}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := newContextWithRequestId(r.Context(), r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
