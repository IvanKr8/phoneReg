package phoneRegPackage

import (
	"regexp"
)

var (
	phoneReg = regexp.MustCompile(`(\+\d{1,2}\s?)?(\(\d{3}\)|\d{3})[\s\.-]?\d{3}[\s\.-]?\d{4}`)
)

// PhoneReg takes a slice of strings and checks each string for phone numbers.
// Returns a new string slice containing only the phone numbers found.

func PhoneReg(slice []string) []string {
	var phonesSlice []string

	for _, val := range slice {
		isValid := phoneReg.MatchString(val)
		if isValid {
			phonesSlice = append(phonesSlice, val)
		}
	}
	return phonesSlice
}
