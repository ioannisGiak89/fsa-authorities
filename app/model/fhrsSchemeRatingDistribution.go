package model

import (
	"fmt"

	"github.com/ioannisGiak89/fsa-authorities/app/utils"
)

// FsaSchemeRatingDistribution interface for FSA schemes distribution
type FsaSchemeRatingDistribution interface {
	// CalculatePercentages calculates the percentages and the totals for authorities in FSA scheme
	CalculatePercentages()
}

// DistributionResults stores the results from the calculations
type DistributionResults struct {
	Total      float64
	Percentage string
}

// FhrsSchemeRatingDistribution represents a rating distribution for authorities in FHRS scheme
type FhrsSchemeRatingDistribution struct {
	Authority          *Authority
	Establishments     []Establishment
	ZeroStar           DistributionResults
	FiveStar           DistributionResults
	FourStar           DistributionResults
	ThreeStar          DistributionResults
	TwoStar            DistributionResults
	OneStar            DistributionResults
	Exempt             DistributionResults
	AwaitingInspection DistributionResults
}

// NewFhrsSchemeRatingDistribution creates a new FhrsSchemeRatingDistribution
func NewFhrsSchemeRatingDistribution(authority *Authority, establishments []Establishment) *FhrsSchemeRatingDistribution {
	return &FhrsSchemeRatingDistribution{Establishments: establishments, Authority: authority}
}

func (f *FhrsSchemeRatingDistribution) CalculatePercentages() {
	for _, establishment := range f.Establishments {
		switch rv := establishment.RatingValue; rv {
		case "Exempt":
			f.Exempt.Total += 1
			break
		case "AwaitingInspection":
			f.AwaitingInspection.Total += 1
			break
		case "0":
			f.ZeroStar.Total += 1
			break
		case "1":
			f.OneStar.Total += 1
			break
		case "2":
			f.TwoStar.Total += 1
			break
		case "3":
			f.ThreeStar.Total += 1
			break
		case "4":
			f.FourStar.Total += 1
			break
		case "5":
			f.FiveStar.Total += 1
			break
		default:
			fmt.Printf("%+v\n", establishment)
			fmt.Println("Fhrs: Unknown rating value " + rv)
			break
		}
	}

	total := len(f.Establishments)
	f.Exempt.Percentage = utils.CalculatePercentage(f.Exempt.Total, total)
	f.AwaitingInspection.Percentage = utils.CalculatePercentage(f.AwaitingInspection.Total, total)
	f.ZeroStar.Percentage = utils.CalculatePercentage(f.ZeroStar.Total, total)
	f.OneStar.Percentage = utils.CalculatePercentage(f.OneStar.Total, total)
	f.TwoStar.Percentage = utils.CalculatePercentage(f.TwoStar.Total, total)
	f.ThreeStar.Percentage = utils.CalculatePercentage(f.ThreeStar.Total, total)
	f.FourStar.Percentage = utils.CalculatePercentage(f.FourStar.Total, total)
	f.FiveStar.Percentage = utils.CalculatePercentage(f.FiveStar.Total, total)
}
