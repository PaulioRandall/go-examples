package server

import (
	"context"
	"net/http"
	"time"

	"github.com/PaulioRandall/ranch-go/routing/moo"

	"github.com/PaulioRandall/go-examples/server/shared/logfmt"
)

func putInCtx(r *http.Request, k, v interface{}) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, k, v)
	return r.WithContext(ctx)
}

func logRequest(w http.ResponseWriter, r *http.Request) *http.Request {
	logfmt.LogInfo(r.URL.String())
	return nil
}

func timeRequest(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now().UnixNano()
		h.ServeHTTP(w, r)
		stop := time.Now().UnixNano()

		t := time.Duration(stop - start)
		logfmt.LogInfo("Finished in %v microseconds", t.Microseconds())
	}

	return http.HandlerFunc(f)
}

func replacePathWith(newPath string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = newPath
		f(w, r)
	}
}

func replacePathsWith(newPath string) moo.PreWrapper {
	return func(w http.ResponseWriter, r *http.Request) *http.Request {
		r.URL.Path = newPath
		return r
	}
}
