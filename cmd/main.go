package main

import (
	"fmt"

	thai_id_validator "github.com/auycro/go-thai-id-validator/validator"
)

func main() {
	test_id := []string{
		"1112034563562",
		"1101700230705",
		"110170023073",
		"11017002070d3",
		"rytege54fsfsf",
		"",
	}

	for _, id := range test_id {
		valid := thai_id_validator.ThaiIDValidator(id)
		if valid {
			fmt.Printf("%v is valid!", id)
		} else {
			fmt.Printf("%v is invalid!", id)
		}
	}

}
