package render

import (
	"fmt"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/jedib0t/go-pretty/v6/table"
)

// CompareTable is used to render tables for comparing distributions of food standards results.
type CompareTable struct {
	ratingDistributions []model.FSASchemeRatingDistribution
	Table
}

// NewCompareTable creates a new CompareTable instance.
func NewCompareTable(ratingDistributions []model.FSASchemeRatingDistribution) *CompareTable {
	return &CompareTable{ratingDistributions: ratingDistributions}
}

// CreateTableAndRender checks the scheme type of establishments, and renders an appropriate results
// table for that scheme type.
func (t *CompareTable) CreateTableAndRender() {
	switch t.ratingDistributions[0].(type) {
	case *model.FHRSSchemeRatingDistribution:
		t.createAndRenderFHRSTable()
	case *model.FHISSchemeRatingDistribution:
		t.createAndRenderFHISTable()
	default:
		fmt.Println("scheme type not found")
	}
}

// createAndRenderFHRSTable creates and renders a FHRS results table.
func (t *CompareTable) createAndRenderFHRSTable() {
	t.header = table.Row{"Rating"}
	t.subHeader = table.Row{""}
	fiveStarRow := table.Row{"Five Stars"}
	fourStarRow := table.Row{"Four Stars"}
	threeStarRow := table.Row{"Three Stars"}
	twoStarRow := table.Row{"Two Stars"}
	oneStarRow := table.Row{"One Stars"}
	zeroStarRow := table.Row{"Zero Stars"}
	exemptRow := table.Row{"Exempt"}
	passRow := table.Row{"Pass"}
	awaitingInspectionRow := table.Row{"Awaiting Inspection"}

	for _, fsd := range t.ratingDistributions {
		if fhrsDistribution, ok := fsd.(*model.FHRSSchemeRatingDistribution); ok {
			fhrsDistribution.CalculatePercentages()
			t.header = append(t.header, fhrsDistribution.Authority.Name, fhrsDistribution.Authority.Name)
			t.subHeader = append(t.subHeader, "Percentage", "Total")
			fiveStarRow = append(fiveStarRow, fhrsDistribution.FiveStar.Percentage, fhrsDistribution.FiveStar.Total)
			fourStarRow = append(fourStarRow, fhrsDistribution.FourStar.Percentage, fhrsDistribution.FourStar.Total)
			threeStarRow = append(threeStarRow, fhrsDistribution.ThreeStar.Percentage, fhrsDistribution.ThreeStar.Total)
			twoStarRow = append(twoStarRow, fhrsDistribution.TwoStar.Percentage, fhrsDistribution.TwoStar.Total)
			oneStarRow = append(oneStarRow, fhrsDistribution.OneStar.Percentage, fhrsDistribution.OneStar.Total)
			zeroStarRow = append(zeroStarRow, fhrsDistribution.ZeroStar.Percentage, fhrsDistribution.ZeroStar.Total)
			exemptRow = append(exemptRow, fhrsDistribution.Exempt.Percentage, fhrsDistribution.Exempt.Total)
			passRow = append(passRow, fhrsDistribution.Pass.Percentage, fhrsDistribution.Pass.Total)
			awaitingInspectionRow = append(awaitingInspectionRow, fhrsDistribution.AwaitingInspection.Percentage, fhrsDistribution.AwaitingInspection.Total)
		}
	}

	t.rows = []table.Row{
		fiveStarRow,
		fourStarRow,
		threeStarRow,
		twoStarRow,
		oneStarRow,
		zeroStarRow,
		exemptRow,
		passRow,
		awaitingInspectionRow,
	}

	t.Render()
}

// createAndRenderFHISTable creates and renders a FHIS results table.
func (t *CompareTable) createAndRenderFHISTable() {
	t.header = table.Row{"Rating"}
	t.subHeader = table.Row{""}
	passRow := table.Row{"Pass"}
	exemptRow := table.Row{"Exempt"}
	improvementRequiredRow := table.Row{"Improvement Required"}
	awaitingPublicationRow := table.Row{"Awaiting Publication"}
	awaitingInspectionRow := table.Row{"Awaiting Inspection"}
	passAndEatSafeRow := table.Row{"Pass And Eat Safe"}

	for _, fsd := range t.ratingDistributions {
		if f, ok := fsd.(*model.FHISSchemeRatingDistribution); ok {
			f.CalculatePercentages()
			t.header = append(t.header, f.Authority.Name, f.Authority.Name)
			t.subHeader = append(t.subHeader, "Percentage", "Total")
			passRow = append(passRow, f.Pass.Percentage, f.Pass.Total)
			exemptRow = append(exemptRow, f.Exempt.Percentage, f.Exempt.Total)
			passAndEatSafeRow = append(passAndEatSafeRow, f.PassAndEatSafe.Percentage, f.PassAndEatSafe.Total)
			awaitingInspectionRow = append(awaitingInspectionRow, f.AwaitingInspection.Percentage, f.AwaitingInspection.Total)
			awaitingPublicationRow = append(awaitingPublicationRow, f.AwaitingPublication.Percentage, f.AwaitingPublication.Total)
			improvementRequiredRow = append(improvementRequiredRow, f.ImprovementRequired.Percentage, f.ImprovementRequired.Total)
		}
	}

	t.rows = []table.Row{
		passRow,
		improvementRequiredRow,
		awaitingPublicationRow,
		exemptRow,
		awaitingInspectionRow,
		passAndEatSafeRow,
	}

	t.Render()
}
