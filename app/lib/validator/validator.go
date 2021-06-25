package validator

import (
	"strconv"
	"strings"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
)

type Validator struct{}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{}
}

// IsSchemeValid check if the Schema is supported
func (v *Validator) IsSchemeValid(schemeType string) bool {
	schemeType = strings.ToUpper(schemeType)

	if schemeType == model.FHRS.String() || schemeType == model.FHIS.String() {
		return true
	}

	return false
}

// IsIdValid checks if the provided ID is a number
func (v *Validator) IsIdValid(id string) bool {
	_, err := strconv.Atoi(id)

	if err != nil {
		return false
	}

	return true
}
