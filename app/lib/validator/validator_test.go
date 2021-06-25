package validator_test

import (
	"testing"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidator_IsValidScheme(t *testing.T) {
	t.Run("should return true if the scheme is FHRS or FHIS", func(t *testing.T) {
		v := validator.NewValidator()
		assert.Equal(t, true, v.IsSchemeValid("fhrs"))
		assert.Equal(t, true, v.IsSchemeValid("FHRS"))
		assert.Equal(t, true, v.IsSchemeValid("FHIS"))
		assert.Equal(t, true, v.IsSchemeValid("fhis"))
	})

	t.Run("should return false if the scheme is neither FHRS nor FHIS", func(t *testing.T) {
		v := validator.NewValidator()
		assert.Equal(t, false, v.IsSchemeValid("invalid"))
		assert.Equal(t, false, v.IsSchemeValid("Invalid"))
		assert.Equal(t, false, v.IsSchemeValid(""))
	})
}

func TestValidator_IsValidId(t *testing.T) {
	t.Run("should return true if the id is a number", func(t *testing.T) {
		v := validator.NewValidator()
		assert.Equal(t, true, v.IsIdValid("2"))
		assert.Equal(t, true, v.IsIdValid("123"))
		assert.Equal(t, true, v.IsIdValid("1231231414214"))
		assert.Equal(t, true, v.IsIdValid("0"))
	})

	t.Run("should return false if the id is not a number", func(t *testing.T) {
		v := validator.NewValidator()
		assert.Equal(t, false, v.IsIdValid("invalid"))
		assert.Equal(t, false, v.IsIdValid("*&^"))
		assert.Equal(t, false, v.IsIdValid(""))
	})
}
