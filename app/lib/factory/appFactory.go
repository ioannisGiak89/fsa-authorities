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

// Factory abstracts the creation of instances.
type FsaFactory interface {
	// BuildFsaService builds Fsa Service
	BuildFsaService(client client.FsaClient) service.HygieneRatingSystemService
	// BuildFsaClient builds the Fsa Client
	BuildFsaClient(baseUrl *url.URL) client.FsaClient
	// BuildCompareTable builds a compare table
	BuildCompareTable([]model.FsaSchemeRatingDistribution) *render.CompareTable
	// BuildTable builds a basic table for rendering
	BuildTable(rows []table.Row, header table.Row, subHeader table.Row) *render.Table
	// BuildValidator builds a validator
	BuildValidator() *validator.Validator
}

// AppFactory builds services
type AppFactory struct{}

// FsaFactory creates a Calculator
func NewAppFactory() FsaFactory {
	return &AppFactory{}
}

func (f *AppFactory) BuildFsaService(client client.FsaClient) service.HygieneRatingSystemService {
	return service.NewFsaService(client)
}

func (f *AppFactory) BuildFsaClient(baseUrl *url.URL) client.FsaClient {
	return client.NewFsaRestClient(baseUrl, &http.Client{})
}

func (f *AppFactory) BuildCompareTable(distributions []model.FsaSchemeRatingDistribution) *render.CompareTable {
	return render.NewCompareTable(distributions)
}

func (f *AppFactory) BuildTable(rows []table.Row, header table.Row, subHeader table.Row) *render.Table {
	return render.NewTable(rows, header, subHeader)
}

func (f *AppFactory) BuildValidator() *validator.Validator {
	return validator.NewValidator()
}
