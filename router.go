package main

import (
	"github.com/gorilla/mux"
	"github.com/docker/docker/client"
	"net/http"
	"compress/gzip"
	"strings"
	"io"
	"log"
	"fmt"
)

func NewRouter(cli *client.Client) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//router.Use(GzipMiddleware)
	router.Use(LogMiddleware)
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

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &responseObserver{ResponseWriter: w}
		handler.ServeHTTP(o, r)

		addr := r.RemoteAddr
		if i := strings.LastIndex(addr, ":"); i != -1 {
			addr = addr[:i]
		}
		log.Printf("DEBUG: %s |> %q %d %d %q %q",
			addr,
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			o.status,
			o.written,
			r.Referer(),
			r.UserAgent())
	})
}
