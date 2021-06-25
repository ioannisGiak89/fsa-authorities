package client_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/client"
	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockedHttpClient is used to mock any functions from http.Client
type mockedHttpClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (cl *mockedHttpClient) Do(req *http.Request) (*http.Response, error) {
	return cl.MockDo(req)
}

func TestFsaRestClient_Get(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	require.NoError(t, err)

	t.Run("should return an error if the request fails", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network request failed")
				},
			},
		)

		responseBody, err := fsaClient.Get("path/to/endpoint")

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("network request failed"), err)
	})

	t.Run("should return an error if status code is 404", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("not found"))),
						StatusCode: http.StatusNotFound,
					}, nil
				},
			},
		)

		responseBody, err := fsaClient.Get("path/to/form3/resource/endpoint")

		assert.Equal(t, &model.HttpError{
			StatusCode: http.StatusNotFound,
			Message:    "not found",
		}, err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return an error if status code is 403", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("Forbidden"))),
						StatusCode: http.StatusForbidden,
					}, nil
				},
			},
		)

		responseBody, err := fsaClient.Get("path/to/form3/resource/endpoint")

		assert.Equal(t, &model.HttpError{
			StatusCode: http.StatusForbidden,
			Message:    "Forbidden",
		}, err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return an error if status code is 400", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("Bad request"))),
						StatusCode: http.StatusBadRequest,
					}, nil
				},
			},
		)

		responseBody, err := fsaClient.Get("path/to/form3/resource/endpoint")

		assert.Equal(t, &model.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
		}, err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return an error if status code is 500", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("Interval Server"))),
						StatusCode: http.StatusInternalServerError,
					}, nil
				},
			},
		)

		responseBody, err := fsaClient.Get("path/to/form3/resource/endpoint")

		assert.Equal(t, &model.HttpError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Interval Server",
		}, err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return a custom response is status code is 200", func(t *testing.T) {
		fsaClient := client.NewFsaRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("A valid response"))),
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		)
		mockedBodyToBytes, err := ioutil.ReadAll(ioutil.NopCloser(bytes.NewReader([]byte("A valid response"))))
		require.NoError(t, err)

		cr, err := fsaClient.Get("path/to/endpoint")

		assert.Equal(t, &model.CustomResponse{
			StatusCode:   http.StatusOK,
			ResponseBody: mockedBodyToBytes,
		}, cr)
		assert.Nil(t, err)
	})
}
