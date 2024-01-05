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
	lastCitizenDigit, _ := strconv.Atoi(citizenID[len(citizenID)-1:])

	return validateCheckSumCitizenID(checkSum, lastCitizenDigit)
}
func calculationSumCitizenID(citizenID string) int {
	sum := 0
	for index, numberOfCitizenID := range citizenID[0 : len(citizenID)-1] {
		sum += getCitizenIDIndexValue(numberOfCitizenID) * (maxLengthOfThaiCitizenID - index)
	}
	return sum
}

func getCitizenIDIndexValue(index int32) int {
	citizenIDIndexValue, _ := strconv.Atoi(string(index))
	return citizenIDIndexValue
}
func validateCheckSumCitizenID(checkSum, lastCitizenDigit int) bool {
	return checkSum == lastCitizenDigit
}
