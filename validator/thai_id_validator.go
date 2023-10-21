package validator

import (
	"regexp"
	"strconv"
)

const maxLengthOfThaiCitizenID = 13

func ThaiIDValidator(citizenID string) bool {
	if citizenID == "" {
		return false
	}

	r, _ := regexp.Compile("!/^[0-9]{13}$/g")
	if r.MatchString(citizenID) {
		return false
	}

	sum := calculationSumCitizenID(citizenID)

	checkSum := (11 - sum%11) % 10
	lastNum, _ := strconv.Atoi(citizenID[len(citizenID)-1:])

	return checkSum == lastNum
}
func getCitizenIDIndex(index int32) int {
	asciiCodeCitizenID, _ := strconv.Atoi(string(index))
	return asciiCodeCitizenID
}
func calculationSumCitizenID(citizenID string) int {
	sum := 0
	for index, numberOfCitzenID := range citizenID[0 : len(citizenID)-1] {
		sum += getCitizenIDIndex(numberOfCitzenID) * (maxLengthOfThaiCitizenID - index)
	}
	return sum
}
