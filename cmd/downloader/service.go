package downloader

import (
	"errors"
	"net/http"
)

type DLService interface {
	Fetch(targetURL string) (<-chan *http.Response, <-chan error)
}

type fetchDLService struct{}

func NewDLService() DLService {
	return &fetchDLService{}
}

func (*fetchDLService) Fetch(targetURL string) (<-chan *http.Response, <-chan error) {
	out := make(chan *http.Response, 1)
	errs := make(chan error, 1)

	if targetURL == "" {
		errs <- errors.New("domain cannot be empty string")
		return out, errs
	}
	go func() {
		client := http.Client{}
		resp, _ := client.Get(targetURL)
		out <- resp
		close(out)
		close(errs)
	}()
	return out, errs
}
