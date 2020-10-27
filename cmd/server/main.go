package server

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	// Dir sets the target directory for the assets (exported for tests)
	Dir string

	// Logs Turn output logs on/off (exported for tests)
	Logs bool
)

// GZipServer is the main application struct
// - httpServer is the main server instance
// - gzon flags whether gzip compression is on/off
type GZipServer struct {
	httpServer http.Server
	gzon       bool
	// Running    chan struct{}
}

type gZipWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

// Run will fire up the server to run the demo
// - port defines the server's port
// - gzip enables/disables gzip compression
// - assetDir what files to server
// - logs enables/disables output logs to the screen
func Run(port string, gzip bool, assetDir string, logs bool) error {
	Dir = assetDir
	Logs = logs

	gz := &GZipServer{
		httpServer: http.Server{
			Addr:    port,
			Handler: http.HandlerFunc(fileHandler),
		},
		gzon: gzip,
		// Running: make(chan struct{}),
	}

	err := gz.Start()
	if err != nil {
		return err
	}
	return nil
}

// Start will fire up the server using ListenAndServe
func (gs *GZipServer) Start() error {
	if Logs {
		log.Printf("starting server on port: %v", gs.httpServer.Addr)
		log.Printf("gzip compression is: %v", boolHelper(gs.gzon))
	}
	err := gs.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// Shutdown will stop the server by calling Close on the server
func (gs *GZipServer) Shutdown() error {
	if Logs {
		log.Printf("shutting down server on port: %v", gs.httpServer.Addr)
	}
	gs.httpServer.Close()
	// close(gs.Running)
	if err := gs.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (gs *GZipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var writer http.ResponseWriter = w
	if gs.gzon && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gzw := gZipWriter{ResponseWriter: w, WriteCloser: gzip.NewWriter(w)}
		defer gzw.Close()
		writer = gzw
		w.Header().Add("Content-Encoding", "gzip")
	}
	gs.httpServer.Handler.ServeHTTP(writer, r)
}

func (gzw gZipWriter) Write(b []byte) (int, error) {
	return gzw.WriteCloser.Write(b)
}

func fileHandler(writer http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(Dir)).ServeHTTP(writer, r)
}

func boolHelper(state bool) string {
	if state {
		return "enabled"
	}
	return "disabled"
}
