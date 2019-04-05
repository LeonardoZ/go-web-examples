package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/foo", Chain(foo, Method("GET"), Logging()))
	http.HandleFunc("/bar", Chain(bar, Method("GET"), Logging()))

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func createNewMiddleware() Middleware {
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			// do something here
			next(w, r)
		}
		return handler
	}
	return middleware
}

func Logging() Middleware {
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			next(w, r)
		}
		return handler
	}
	return middleware
}

func Method(m string) Middleware {
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			next(w, r)
		}
		return handler
	}
	return middleware
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
