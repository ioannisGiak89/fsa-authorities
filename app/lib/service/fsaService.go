package service

import (
	"encoding/json"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/client"
	"github.com/ioannisGiak89/fsa-authorities/app/model"
)

// HygieneRatingSystemService defines the Hygiene Rating System Service
type HygieneRatingSystemService interface {
	// GetAuthorities gets a list with all the authorities and some basic information
	GetAuthorities() (*model.AuthoritiesResponse, error)

	// GetAuthorityByID gets a specific authority by it's ID.
	GetAuthorityByID(authorityId string) (*model.Authority, error)

	// GetEstablishments gets all the Establishments for a given authority id
	GetEstablishments(authorityId string) (*model.EstablishmentsResponse, error)
}

// FSAService implements the HygieneRatingSystemService interface
type FSAService struct {
	fsaClient client.FSAClient
}

// NewFSAService creates a new FSAService
func NewFSAService(fsaClient client.FSAClient) HygieneRatingSystemService {
	return &FSAService{
		fsaClient: fsaClient,
	}
}

func (fs *FSAService) GetAuthorityByID(authorityId string) (*model.Authority, error) {
	cr, err := fs.fsaClient.Get("Authorities/" + authorityId)

	if err != nil {
		return nil, err
	}

	var authority model.Authority
	err = json.Unmarshal(cr.ResponseBody, &authority)

	if err != nil {
		return nil, err
	}

	return &authority, nil
}

func (fs *FSAService) GetAuthorities() (*model.AuthoritiesResponse, error) {
	cr, err := fs.fsaClient.Get("Authorities/basic")

	if err != nil {
		return nil, err
	}

	var authorities model.AuthoritiesResponse
	err = json.Unmarshal(cr.ResponseBody, &authorities)

	if err != nil {
		return nil, err
	}

	return &authorities, nil
}

func (fs *FSAService) GetEstablishments(authorityId string) (*model.EstablishmentsResponse, error) {
	cr, err := fs.fsaClient.Get("Establishments?localAuthorityId=" + authorityId)

	if err != nil {
		return nil, err
	}

	var establishments model.EstablishmentsResponse
	err = json.Unmarshal(cr.ResponseBody, &establishments)

	if err != nil {
		return nil, err
	}

	return &establishments, nil
}
