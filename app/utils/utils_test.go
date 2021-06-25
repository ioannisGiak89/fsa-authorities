package utils_test

import (
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePercentage(t *testing.T) {
	t.Run("should return a percentage in string", func(t *testing.T) {
		p := utils.CalculatePercentage(10.0, 20)
		assert.Equal(t, "50.00%", p)
	})

	t.Run("should return 0.00%", func(t *testing.T) {
		p := utils.CalculatePercentage(0, 30)
		assert.Equal(t, "0.00%", p)
	})

	t.Run("should return an empty string if totalEstablishments is 0", func(t *testing.T) {
		p := utils.CalculatePercentage(30, 0)
		assert.Equal(t, "", p)
	})
}
