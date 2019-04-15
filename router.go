package main

import (
	"github.com/gorilla/mux"
	"github.com/docker/docker/client"
	"net/http"
	"compress/gzip"
	"strings"
	"io"
)

func NewRouter(cli *client.Client) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//router.Use(GzipMiddleware)
	router.HandleFunc("/", Info).Methods("GET")
	router.HandleFunc("/logs/{container}", Logs(cli)).Methods("GET")
	return router
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}
