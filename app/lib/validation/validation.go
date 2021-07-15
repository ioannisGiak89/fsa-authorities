package validation

import (
	"strconv"
	"strings"

	"github.com/ioannisGiak89/fsa-authorities/app/model"
)

// IsSchemeValid check if the Schema is supported
func IsSchemeValid(schemeType string) bool {
	schemeType = strings.ToUpper(schemeType)

	if schemeType == model.FHRS.String() || schemeType == model.FHIS.String() {
		return true
	}

	return false
}

// IsIDValid checks if the provided ID is a number
func IsIDValid(id string) bool {
	_, err := strconv.Atoi(id)

	if err != nil {
		return false
	}

	return true
}
