package validator

import (
	"regexp"
	"strconv"
)

func ThaiIDValidator(id string) bool {
	if id == "" {
		return false
	}

	r, _ := regexp.Compile("!/^[0-9]{13}$/g")
	if r.MatchString(id) {
		return false
	}

	sum := 0
	for i, n := range id[0 : len(id)-1] {
		xint, _ := strconv.Atoi(string(n))
		sum += xint * (13 - i)
	}

	checkSum := (11 - sum%11) % 10
	last_num, _ := strconv.Atoi(id[len(id)-1:])

	return checkSum == last_num
}
