package render

import (
	"fmt"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/jedib0t/go-pretty/v6/table"
)

// CompareTable table used for comparing distributions
type CompareTable struct {
	ratingDistributions []model.FsaSchemeRatingDistribution
	Table
}

// NewCompareTable creates a new CompareTable
func NewCompareTable(ratingDistributions []model.FsaSchemeRatingDistribution) *CompareTable {
	return &CompareTable{ratingDistributions: ratingDistributions}
}

// CreateTableAndRender checks the scheme type of establishments
func (t *CompareTable) CreateTableAndRender() {
	switch t.ratingDistributions[0].(type) {
	case *model.FhrsSchemeRatingDistribution:
		t.createAndRenderFhrsTable()
		break
	case *model.FhisSchemeRatingDistribution:
		t.createAndRenderFhisTable()
		break
	default:
		fmt.Println("scheme type not found")
		break
	}
}

// createAndRenderFhrsTable creates and renders a FhrsTable
func (t *CompareTable) createAndRenderFhrsTable() {
	t.header = table.Row{"Rating"}
	t.subHeader = table.Row{""}
	fiveStarRow := table.Row{"Five Stars"}
	fourStarRow := table.Row{"Four Stars"}
	threeStarRow := table.Row{"Three Stars"}
	twoStarRow := table.Row{"Two Stars"}
	oneStarRow := table.Row{"One Stars"}
	zeroStarRow := table.Row{"Zero Stars"}
	exemptRow := table.Row{"Exempt"}
	awaitingInspectionRow := table.Row{"Awaiting Inspection"}

	for _, fsd := range t.ratingDistributions {
		if fhrsDistribution, ok := fsd.(*model.FhrsSchemeRatingDistribution); ok {
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
		awaitingInspectionRow,
	}

	t.Render()
}

// createAndRenderFhisTable creates and renders a FhisTable
func (t *CompareTable) createAndRenderFhisTable() {
	t.header = table.Row{"Rating"}
	t.subHeader = table.Row{""}
	passRow := table.Row{"Pass"}
	exemptRow := table.Row{"Exempt"}
	improvementRequiredRow := table.Row{"Improvement Required"}
	awaitingPublicationRow := table.Row{"Awaiting Publication"}
	awaitingInspectionRow := table.Row{"Awaiting Inspection"}
	passAndEatSafeRow := table.Row{"Pass And Eat Safe"}

	for _, fsd := range t.ratingDistributions {
		if f, ok := fsd.(*model.FhisSchemeRatingDistribution); ok {
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
