package client

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

// FsaClient defines the Form3 Resources client
type FsaClient interface {
    Get(path string) ([]byte, error)
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

// Creates a new Form3 rest client
func NewFsa3RestClient(baseUrl *url.URL, httpClient HTTPClient) FsaClient {
    return &FsaRestClient{
        baseUrl: baseUrl,
        client:  httpClient,
    }
}

// Get does a get request to an endpoint
func (cl *FsaRestClient) Get(path string) ([]byte, error) {
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

    if res.StatusCode != http.StatusOK {
        return nil, errors.New(string(resBody))
    }

    return resBody, nil
}
