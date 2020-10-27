package server

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	startup()
	code := m.Run()
	// shutdown(m)
	os.Exit(code)
}

func startup() {
	wd, _ := os.Getwd()
	if strings.HasSuffix(wd, "server") {
		wd = strings.TrimSuffix(wd, "/cmd/server")
	}
	Dir = wd + "/assets"
	Logs = false
}

func TestFileHandler(t *testing.T) {

	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:9090/",
		nil,
	)
	assert.Nil(t, err)

	rec := httptest.NewRecorder()
	fileHandler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "CONTACT FORM")
}

func TestGzipServer(t *testing.T) {
	cases := []struct {
		gzip bool
		resp string
		logs bool
	}{
		{true, "CONTACT FORM", true},
		{false, "CONTACT FORM", false},
	}
	// byte buffer for capturing log output
	var buf bytes.Buffer
	for _, c := range cases {
		if c.logs {
			Logs = c.logs
			log.SetOutput(&buf)
		}
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:9091/",
			nil,
		)
		assert.Nil(t, err)

		gz := &GZipServer{
			httpServer: http.Server{
				Addr:    ":9091",
				Handler: http.HandlerFunc(fileHandler),
			},
			gzon: c.gzip,
		}
		rec := httptest.NewRecorder()

		if c.gzip {
			req.Header.Set("Accept-Encoding", "gzip")
		}
		gz.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, decompress(rec.Body.String()), c.resp)
	}
}

func TestStart(t *testing.T) {
	cases := []struct {
		gzip bool
		resp string
	}{
		{true, "CONTACT FORM"},
		{false, "CONTACT FORM"},
	}
	for _, c := range cases {
		gz := &GZipServer{
			httpServer: http.Server{
				Addr:    ":9091",
				Handler: http.HandlerFunc(fileHandler),
			},
			gzon: c.gzip,
			// Running: make(chan struct{}),
		}
		serviceRunning := make(chan struct{})
		serviceDone := make(chan struct{})
		go func() {
			close(serviceRunning)
			_ = gz.Start()
			defer close(serviceDone)
		}()

		// wait until the goroutine started to run (1)
		<-serviceRunning

		// interact with your service to test whatever you want
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:9091/",
			nil,
		)
		assert.Nil(t, err)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()
		rb, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Contains(t, decompress(string(rb)), c.resp)

		// stop the service (2)
		// time.AfterFunc(3*time.Second, func() {
		err = gz.Shutdown()
		if err != nil {
			return
		}

		// wait until the service is shutdown (3)
		<-serviceDone
	}
}

func decompress(in string) string {
	if http.DetectContentType([]byte(in)) != "application/x-gzip" {
		return in
	}
	r := bytes.NewReader([]byte(in))
	gr, _ := gzip.NewReader(r)
	result, _ := ioutil.ReadAll(gr)
	return string(result)
}
