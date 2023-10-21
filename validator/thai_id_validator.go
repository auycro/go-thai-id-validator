package validator

import (
	"regexp"
	"strconv"
)

func ThaiIDValidator(citizenID string) bool {
	if citizenID == "" {
		return false
	}

	r, _ := regexp.Compile("!/^[0-9]{13}$/g")
	if r.MatchString(citizenID) {
		return false
	}

	sum := 0
	for i, n := range citizenID[0 : len(citizenID)-1] {
		xint, _ := strconv.Atoi(string(n))
		sum += xint * (13 - i)
	}

	checkSum := (11 - sum%11) % 10
	last_num, _ := strconv.Atoi(citizenID[len(citizenID)-1:])

	return checkSum == last_num
}
