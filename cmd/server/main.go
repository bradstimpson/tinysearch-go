package server

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

type GZipServer struct {
	AssetDir string
}

type gZipWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

func Run(port string, gzip bool, assetDir string) {
	var err error

	if gzip {
		http.Handle("/", http.FileServer(http.Dir(assetDir)))
		err = http.ListenAndServe(port, new(GZipServer))
	} else {
		err = http.ListenAndServe(port, http.FileServer(http.Dir(assetDir)))
	}
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}

func (gs *GZipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var writer http.ResponseWriter
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gzw := gZipWriter{ResponseWriter: w, WriteCloser: gzip.NewWriter(w)}
		defer gzw.Close()
		writer = gzw
		w.Header().Add("Content-Encoding", "gzip")
	} else {
		writer = w
	}
	http.DefaultServeMux.ServeHTTP(writer, r)
}

func (gzw gZipWriter) Write(b []byte) (int, error) {
	return gzw.WriteCloser.Write(b)
}
