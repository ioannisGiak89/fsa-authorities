package service_test

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/service"
	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/ioannisGiak89/fsa-authorities/testutils"
	"github.com/stretchr/testify/require"
)

// Implements the FSAClient interface. This struct is used to mock the FSARESTClient
type mockedHttpClient struct {
	baseUrl *url.URL
	MockGet func(path string) (*model.CustomResponse, error)
}

func (m mockedHttpClient) Get(path string) (*model.CustomResponse, error) {
	return m.MockGet(path)
}

func TestFSAService_GetAuthorities(t *testing.T) {
	t.Run("should return an AuthoritiesResponse", func(t *testing.T) {
		expectedResponse := testutils.GetFakeAuthoritiesResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		mockedResponse := &model.CustomResponse{
			StatusCode:   200,
			ResponseBody: jsonResponse,
		}

		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return mockedResponse, nil
			},
		})

		response, err := fsaService.GetAuthorities()
		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return nil, &model.HTTPError{
					StatusCode: 404,
					Message:    "Not found",
				}
			},
		})

		response, err := fsaService.GetAuthorities()
		assert.Nil(t, response)
		assert.Equal(t, err, &model.HTTPError{
			StatusCode: 404,
			Message:    "Not found",
		})
	})

	t.Run("should return an error if unmarshal fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return &model.CustomResponse{
					StatusCode:   200,
					ResponseBody: []byte{12, 12},
				}, nil
			},
		})

		response, err := fsaService.GetAuthorities()
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}

func TestFSAService_GetAuthorityByID(t *testing.T) {
	t.Run("should return an Authority", func(t *testing.T) {
		expectedResponse := testutils.GetFakeAuthority()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		mockedResponse := &model.CustomResponse{
			StatusCode:   200,
			ResponseBody: jsonResponse,
		}

		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return mockedResponse, nil
			},
		})

		response, err := fsaService.GetAuthorityByID("197")
		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return nil, &model.HTTPError{
					StatusCode: 404,
					Message:    "Not found",
				}
			},
		})

		response, err := fsaService.GetAuthorityByID("197")
		assert.Nil(t, response)
		assert.Equal(t, err, &model.HTTPError{
			StatusCode: 404,
			Message:    "Not found",
		})
	})

	t.Run("should return an error if unmarshal fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return &model.CustomResponse{
					StatusCode:   200,
					ResponseBody: []byte{12, 12},
				}, nil
			},
		})

		response, err := fsaService.GetAuthorityByID("197")
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}

func TestFSAService_GetEstablishments(t *testing.T) {
	t.Run("should return an EstablishmentsResponse", func(t *testing.T) {
		expectedResponse := testutils.GetFakeEstablishmentsResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		mockedResponse := &model.CustomResponse{
			StatusCode:   200,
			ResponseBody: jsonResponse,
		}

		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return mockedResponse, nil
			},
		})

		response, err := fsaService.GetEstablishments("197")
		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return nil, &model.HTTPError{
					StatusCode: 404,
					Message:    "Not found",
				}
			},
		})

		response, err := fsaService.GetEstablishments("197")
		assert.Nil(t, response)
		assert.Equal(t, err, &model.HTTPError{
			StatusCode: 404,
			Message:    "Not found",
		})
	})

	t.Run("should return an error if unmarshal fails", func(t *testing.T) {
		fsaService := service.NewFSAService(&mockedHttpClient{
			MockGet: func(path string) (*model.CustomResponse, error) {
				return &model.CustomResponse{
					StatusCode:   200,
					ResponseBody: []byte{12, 12},
				}, nil
			},
		})

		response, err := fsaService.GetEstablishments("197")
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}
