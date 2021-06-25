package utils

import (
	"fmt"
)

// CalculatePercentage does the maths and converts float to string
func CalculatePercentage(ratingCounter float64, totalEstablishments int) string {
	if totalEstablishments == 0 {
		return ""
	}

	return fmt.Sprintf("%.2f%s", float32(100.0*ratingCounter/float64(totalEstablishments)), "%")
}
