package validation_test

import (
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidator_IsValidScheme(t *testing.T) {
	t.Run("should return true if the scheme is FHRS or FHIS", func(t *testing.T) {
		assert.Equal(t, true, validation.IsSchemeValid("fhrs"))
		assert.Equal(t, true, validation.IsSchemeValid("FHRS"))
		assert.Equal(t, true, validation.IsSchemeValid("FHIS"))
		assert.Equal(t, true, validation.IsSchemeValid("fhis"))
	})

	t.Run("should return false if the scheme is neither FHRS nor FHIS", func(t *testing.T) {
		assert.Equal(t, false, validation.IsSchemeValid("invalid"))
		assert.Equal(t, false, validation.IsSchemeValid("Invalid"))
		assert.Equal(t, false, validation.IsSchemeValid(""))
	})
}

func TestValidator_IsValidId(t *testing.T) {
	t.Run("should return true if the id is a number", func(t *testing.T) {
		assert.Equal(t, true, validation.IsIDValid("2"))
		assert.Equal(t, true, validation.IsIDValid("123"))
		assert.Equal(t, true, validation.IsIDValid("1231231414214"))
		assert.Equal(t, true, validation.IsIDValid("0"))
	})

	t.Run("should return false if the id is not a number", func(t *testing.T) {
		assert.Equal(t, false, validation.IsIDValid("invalid"))
		assert.Equal(t, false, validation.IsIDValid("*&^"))
		assert.Equal(t, false, validation.IsIDValid(""))
	})
}
