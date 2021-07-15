package model_test

import (
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/ioannisGiak89/fsa-authorities/testutils"
	"github.com/stretchr/testify/assert"
)

func TestFHRSSchemeRatingDistribution_CalculatePercentages(t *testing.T) {
	t.Run("should calculate the rating distribution for FHRS scheme", func(t *testing.T) {
		fr := model.NewFHRSSchemeRatingDistribution(testutils.GetFakeAuthority(), testutils.GetFakeFHRSEstablishments())
		fr.CalculatePercentages()
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.ZeroStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      2,
				Percentage: "20.00%",
			},
			fr.OneStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      2,
				Percentage: "20.00%",
			},
			fr.TwoStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.ThreeStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.FourStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.FiveStar,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.Exempt,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fr.AwaitingInspection,
		)
	})
}
