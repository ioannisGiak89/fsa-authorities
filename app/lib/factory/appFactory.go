package factory

import (
	"net/http"
	"net/url"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/client"
	"github.com/ioannisGiak89/fsa-authorities/app/lib/service"
	"github.com/ioannisGiak89/fsa-authorities/app/lib/validator"
	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/ioannisGiak89/fsa-authorities/app/render"
	"github.com/jedib0t/go-pretty/v6/table"
)

// FSAFactory abstracts the creation of instances.
type FSAFactory interface {
	// BuildFSAService builds the FSA Service
	BuildFSAService(client client.FSAClient) service.HygieneRatingSystemService
	// BuildFSAClient builds the FSA Client
	BuildFSAClient(baseUrl *url.URL) client.FSAClient
	// BuildCompareTable builds a compare table
	BuildCompareTable([]model.FSASchemeRatingDistribution) *render.CompareTable
	// BuildTable builds a basic table for rendering
	BuildTable(rows []table.Row, header table.Row, subHeader table.Row) *render.Table
	// BuildValidator builds a validator
	BuildValidator() *validator.Validator
}

// AppFactory builds services
type AppFactory struct{}

// NewAppFactory creates a new FSAFactory instance.
func NewAppFactory() FSAFactory {
	return &AppFactory{}
}

func (f *AppFactory) BuildFSAService(client client.FSAClient) service.HygieneRatingSystemService {
	return service.NewFSAService(client)
}

func (f *AppFactory) BuildFSAClient(baseUrl *url.URL) client.FSAClient {
	return client.NewFSARestClient(baseUrl, &http.Client{})
}

func (f *AppFactory) BuildCompareTable(distributions []model.FSASchemeRatingDistribution) *render.CompareTable {
	return render.NewCompareTable(distributions)
}

func (f *AppFactory) BuildTable(rows []table.Row, header table.Row, subHeader table.Row) *render.Table {
	return render.NewTable(rows, header, subHeader)
}

func (f *AppFactory) BuildValidator() *validator.Validator {
	return validator.NewValidator()
}
