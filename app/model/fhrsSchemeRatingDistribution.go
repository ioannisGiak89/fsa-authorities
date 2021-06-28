package model

import (
	"fmt"

	"github.com/ioannisGiak89/fsa-authorities/app/utils"
)

// FSASchemeRatingDistribution interface for FSA schemes distribution
type FSASchemeRatingDistribution interface {
	// CalculatePercentages calculates the percentages and the totals for authorities in FSA scheme
	CalculatePercentages()
}

// DistributionResults stores the results from the calculations
type DistributionResults struct {
	Total      float64
	Percentage string
}

// FHRSSchemeRatingDistribution represents a rating distribution for authorities in FHRS scheme
type FHRSSchemeRatingDistribution struct {
	Authority          *Authority
	Establishments     []Establishment
	ZeroStar           DistributionResults
	FiveStar           DistributionResults
	FourStar           DistributionResults
	ThreeStar          DistributionResults
	TwoStar            DistributionResults
	OneStar            DistributionResults
	Exempt             DistributionResults
	Pass               DistributionResults
	AwaitingInspection DistributionResults
}

// NewFHRSSchemeRatingDistribution creates a new FHRSSchemeRatingDistribution
func NewFHRSSchemeRatingDistribution(authority *Authority, establishments []Establishment) *FHRSSchemeRatingDistribution {
	return &FHRSSchemeRatingDistribution{Establishments: establishments, Authority: authority}
}

func (f *FHRSSchemeRatingDistribution) CalculatePercentages() {
	for _, establishment := range f.Establishments {
		switch rv := establishment.RatingValue; rv {
		case "Pass":
			f.Pass.Total += 1
		case "Exempt":
			f.Exempt.Total += 1
		case "AwaitingInspection":
			f.AwaitingInspection.Total += 1
		case "0":
			f.ZeroStar.Total += 1
		case "1":
			f.OneStar.Total += 1
		case "2":
			f.TwoStar.Total += 1
		case "3":
			f.ThreeStar.Total += 1
		case "4":
			f.FourStar.Total += 1
		case "5":
			f.FiveStar.Total += 1
		default:
			fmt.Printf("%+v\n", establishment)
			fmt.Println("FHRS: Unknown rating value " + rv)
		}
	}

	total := len(f.Establishments)
	f.Pass.Percentage = utils.CalculatePercentage(f.Pass.Total, total)
	f.Exempt.Percentage = utils.CalculatePercentage(f.Exempt.Total, total)
	f.AwaitingInspection.Percentage = utils.CalculatePercentage(f.AwaitingInspection.Total, total)
	f.ZeroStar.Percentage = utils.CalculatePercentage(f.ZeroStar.Total, total)
	f.OneStar.Percentage = utils.CalculatePercentage(f.OneStar.Total, total)
	f.TwoStar.Percentage = utils.CalculatePercentage(f.TwoStar.Total, total)
	f.ThreeStar.Percentage = utils.CalculatePercentage(f.ThreeStar.Total, total)
	f.FourStar.Percentage = utils.CalculatePercentage(f.FourStar.Total, total)
	f.FiveStar.Percentage = utils.CalculatePercentage(f.FiveStar.Total, total)
}
