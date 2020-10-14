package downloader

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDLService struct {
	mock.Mock
}

type MockDownloader struct {
	mock.Mock
}

func (mock *MockDLService) Fetch(domain string) (<-chan *http.Response, <-chan error) {
	out := make(chan *http.Response, 1)
	errs := make(chan error, 1)
	args := mock.Called()

	out <- args.Get(0).(*http.Response)
	errs <- args.Error(1)
	return out, errs
}

func (mock *MockDownloader) Get(domain string) (Posts, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(Posts), args.Error(1)
}

func TestGetDownloadNoURL(t *testing.T) {
	domain := ""

	testDLService := NewDLService()
	testDownloader := NewDownloader(testDLService, true, false, false)
	results, err := testDownloader.Get(domain)

	assert.Nil(t, results)
	assert.NotNil(t, err)

}

func TestNonExistentDomain(t *testing.T) {
	domain := "http://example.test"

	testDLService := NewDLService()
	testDownloader := NewDownloader(testDLService, true, false, false)
	results, err := testDownloader.Get(domain)

	assert.Nil(t, results)
	assert.NotNil(t, err)

}

func TestJekyll(t *testing.T) {
	domain := "http://example.test"
	response := http.Response{
		Status: "200",
		Body:   ioutil.NopCloser(strings.NewReader("test body")),
	}

	mockDLService := new(MockDLService)
	mockDLService.On("Fetch").Return(&response, nil)
	testDownloader := NewDownloader(mockDLService, false, true, false)
	results, err := testDownloader.Get(domain)
	assert.Nil(t, results)
	assert.Equal(t, "Jekyll marshalling method not currently supported", err.Error())
}

func TestHugo(t *testing.T) {
	domain := "http://example.test"
	response := http.Response{
		Status: "200",
		Body:   ioutil.NopCloser(strings.NewReader("test body")),
	}

	mockDLService := new(MockDLService)
	mockDLService.On("Fetch").Return(&response, nil)
	testDownloader := NewDownloader(mockDLService, false, false, true)
	results, err := testDownloader.Get(domain)
	assert.Nil(t, results)
	assert.Equal(t, "Hugo marshalling method not currently supported", err.Error())
}

func TestGetMockDownload(t *testing.T) {
	domain := "http://example.test"

	mockDownloader := new(MockDownloader)
	post := Post{Title: "test", Slug: "test-test", Content: "test content"}
	var posts Posts = []Post{post}

	mockDownloader.On("Get").Return(posts, nil)

	results, _ := mockDownloader.Get(domain)

	assert.NotNil(t, results)
	assert.Equal(t, "test", results[0].Title)
	assert.Equal(t, "test-test", results[0].Slug)
}

func TestUnmarshalWP(t *testing.T) {
	jsonfile, err := os.Open("../fixtures/test-wp-output.json")
	assert.Nil(t, err)
	defer jsonfile.Close()

	buf := new(bytes.Buffer)
	_, err1 := buf.ReadFrom(jsonfile)
	assert.Nil(t, err1)

	var p Posts
	err2 := p.unmarshalWP(buf.Bytes())
	assert.Nil(t, err2)
	assert.Equal(t, "Test1 Title", p[0].Title)
}
