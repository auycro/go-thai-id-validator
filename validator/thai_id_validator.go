package validator

import (
	"regexp"
	"strconv"
	"sync"
)

var thaiIDRegexPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return regexp.MustCompile("!/^[0-9]{13}$/g")
	},
}

func ThaiIDValidator(id string) bool {
	if len(id) != 13 {
		return false
	}

	thaiIDRegex := thaiIDRegexPool.Get().(*regexp.Regexp)
	defer thaiIDRegexPool.Put(thaiIDRegex)
	if thaiIDRegex.MatchString(id) {
		return false
	}

	sum := 0
	for i := 0; i < len(id)-1; i++ {
		xint, _ := strconv.Atoi(string(id[i]))
		sum += xint * (13 - i)
	}

	checkSum := (11 - sum%11) % 10
	lastNum, _ := strconv.Atoi(string(id[len(id)-1]))

	return checkSum == lastNum
}
