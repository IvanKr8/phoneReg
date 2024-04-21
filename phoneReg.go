package phoneRegPackage

import (
	"regexp"
	"strconv"
)

var (
	phoneReg = regexp.MustCompile(`(\+\d{1,2}\s?)?(\(\d{3}\)|\d{3})[\s\.-]?\d{3}[\s\.-]?\d{4}`)
)

// Function accepts a slice of numbers, among then it looks for phone numbers
// Returns also a slice from an integer

func PhoneReg(slice []int) []int {
	var phonesSlice []int

	for _, val := range slice {
		strVal := strconv.Itoa(val)
		isValid := phoneReg.MatchString(strVal)
		if isValid {
			phonesSlice = append(phonesSlice, val)
		}
	}
	return phonesSlice
}
