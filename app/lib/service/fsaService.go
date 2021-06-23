package service

import (
	"encoding/json"

	"github.com/ioannisGiak89/compare-fsa-ratings/app/lib/client"
	"github.com/ioannisGiak89/compare-fsa-ratings/app/model"
)

// HygieneRatingSystemService defines the Hygiene Rating System Service
type HygieneRatingSystemService interface {
	// GetAuthorities gets a list with all the authorities and some basic information
	GetAuthorities() (*model.AuthoritiesResponse, error)

	// GetAuthorities gets a list with all the authorities and some basic information
	GetAuthorityByID(authorityId string) (*model.Authority, error)

	// GetEstablishments gets all the Establishments for a given authority id
	GetEstablishments(authorityId string) (*model.EstablishmentsResponse, error)
}

// FsaService implements the HygieneRatingSystemService interface
type FsaService struct {
	fsaClient client.FsaClient
}

// NewFsaService creates a new FsaService
func NewFsaService(fsaClient *client.FsaClient) HygieneRatingSystemService {
	return &FsaService{
		fsaClient: *fsaClient,
	}
}

func (fs *FsaService) GetAuthorityByID(authorityId string) (*model.Authority, error) {

	responseBody, err := fs.fsaClient.Get("Authorities/" + authorityId)

	if err != nil {
		return nil, err
	}

	var authority model.Authority
	err = json.Unmarshal(responseBody, &authority)

	if err != nil {
		return nil, err
	}

	return &authority, nil
}

func (fs *FsaService) GetAuthorities() (*model.AuthoritiesResponse, error) {
	responseBody, err := fs.fsaClient.Get("Authorities/basic")

	if err != nil {
		return nil, err
	}

	var authorities model.AuthoritiesResponse
	err = json.Unmarshal(responseBody, &authorities)

	if err != nil {
		return nil, err
	}

	return &authorities, nil
}

func (fs *FsaService) GetEstablishments(authorityId string) (*model.EstablishmentsResponse, error) {
	responseBody, err := fs.fsaClient.Get("Establishments?localAuthorityId=" + authorityId)

	if err != nil {
		return nil, err
	}

	var establishments model.EstablishmentsResponse
	err = json.Unmarshal(responseBody, &establishments)

	if err != nil {
		return nil, err
	}

	return &establishments, nil
}
