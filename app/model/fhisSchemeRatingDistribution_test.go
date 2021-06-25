package model_test

import (
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
	"github.com/ioannisGiak89/fsa-authorities/testUtils"
	"github.com/stretchr/testify/assert"
)

func TestFhisSchemeRatingDistribution_CalculatePercentages(t *testing.T) {
	t.Run("should calculate the rating distribution for fhis scheme", func(t *testing.T) {
		fi := model.NewFhisSchemeRatingDistribution(testUtils.GetFakeAuthority(), testUtils.GetFakeFhisEstablishments())
		fi.CalculatePercentages()
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      3,
				Percentage: "30.00%",
			},
			fi.Pass,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      2,
				Percentage: "20.00%",
			},
			fi.PassAndEatSafe,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fi.AwaitingPublication,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      2,
				Percentage: "20.00%",
			},
			fi.AwaitingInspection,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fi.ImprovementRequired,
		)
		assert.Equal(
			t,
			model.DistributionResults{
				Total:      1,
				Percentage: "10.00%",
			},
			fi.Exempt,
		)
	})
}
