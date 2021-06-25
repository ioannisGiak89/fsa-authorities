package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
)

// FsaClient defines the FsaClient client
type FsaClient interface {
	// Get does a get request to an endpoint
	Get(path string) (*model.CustomResponse, error)
}

// HTTPClient interface. This interface is implemented by http.Client and is used for mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// FsaRestClient implements the FsaClient interface
type FsaRestClient struct {
	baseUrl *url.URL
	client  HTTPClient
}

// Creates a new FsaRestClient
func NewFsaRestClient(baseUrl *url.URL, httpClient HTTPClient) FsaClient {
	return &FsaRestClient{
		baseUrl: baseUrl,
		client:  httpClient,
	}
}

func (cl *FsaRestClient) Get(path string) (*model.CustomResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(
		"%s%s",
		cl.baseUrl.String(),
		path,
	), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-version", "2")

	res, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		return &model.CustomResponse{
			StatusCode:   res.StatusCode,
			ResponseBody: resBody,
		}, nil
	}

	return nil, &model.HttpError{
		StatusCode: res.StatusCode,
		Message:    string(resBody),
	}
}
