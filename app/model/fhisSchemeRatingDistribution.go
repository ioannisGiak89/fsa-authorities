package model

import (
	"fmt"
	"strings"

	"github.com/ioannisGiak89/fsa-authorities/app/utils"
)

// FHISSchemeRatingDistribution represents a rating distribution for authorities in FHIS scheme
type FHISSchemeRatingDistribution struct {
	Authority           *Authority
	Establishments      []Establishment
	Pass                DistributionResults
	AwaitingPublication DistributionResults
	ImprovementRequired DistributionResults
	AwaitingInspection  DistributionResults
	PassAndEatSafe      DistributionResults
	Exempt              DistributionResults
}

// NewFHISSchemeRatingDistribution creates a new FHISSchemeRatingDistribution
func NewFHISSchemeRatingDistribution(authority *Authority, establishments []Establishment) *FHISSchemeRatingDistribution {
	return &FHISSchemeRatingDistribution{Authority: authority, Establishments: establishments}
}

func (f *FHISSchemeRatingDistribution) CalculatePercentages() {
	for _, establishment := range f.Establishments {
		switch strings.TrimSpace(establishment.RatingValue) {
		case "Pass":
			f.Pass.Total += 1
		case "Improvement Required":
			f.ImprovementRequired.Total += 1
		case "Awaiting Inspection":
			f.AwaitingInspection.Total += 1
		case "Exempt":
			f.Exempt.Total += 1
		case "Awaiting Publication":
			f.AwaitingPublication.Total += 1
		case "Pass and Eat Safe":
			f.PassAndEatSafe.Total += 1
		default:
			fmt.Println("FHIS: Unknown rating value " + establishment.RatingValue)
		}
	}

	total := len(f.Establishments)
	f.Exempt.Percentage = utils.CalculatePercentage(f.Exempt.Total, total)
	f.AwaitingInspection.Percentage = utils.CalculatePercentage(f.AwaitingInspection.Total, total)
	f.Pass.Percentage = utils.CalculatePercentage(f.Pass.Total, total)
	f.ImprovementRequired.Percentage = utils.CalculatePercentage(f.ImprovementRequired.Total, total)
	f.AwaitingPublication.Percentage = utils.CalculatePercentage(f.AwaitingPublication.Total, total)
	f.PassAndEatSafe.Percentage = utils.CalculatePercentage(f.PassAndEatSafe.Total, total)
}
