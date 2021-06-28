package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
)

// FSAClient defines the FSAClient client
type FSAClient interface {
	// Get does a get request to an endpoint
	Get(path string) (*model.CustomResponse, error)
}

// HTTPClient interface. This interface is implemented by http.Client and is used for mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// FSARESTClient implements the FSAClient interface
type FSARESTClient struct {
	baseUrl *url.URL
	client  HTTPClient
}

// NewFSARestClient returns a new FSARESTClient instance.
func NewFSARestClient(baseUrl *url.URL, httpClient HTTPClient) *FSARESTClient {
	return &FSARESTClient{
		baseUrl: baseUrl,
		client:  httpClient,
	}
}

func (cl *FSARESTClient) Get(path string) (*model.CustomResponse, error) {
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

	return nil, &model.HTTPError{
		StatusCode: res.StatusCode,
		Message:    string(resBody),
	}
}
