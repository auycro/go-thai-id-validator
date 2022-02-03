package validator

import (
	"testing"
)

func TestCorrectNumber_ThaiIDValidator(t *testing.T) {
	actual := ThaiIDValidator("1112034563562")
	expected := true
	if actual != expected {
		t.Fatalf(`ThaiIDValidator("1131890334671") got %v, want %v`, actual, expected)
	}
}

func TestIncorrectNumber_ThaiIDValidator(t *testing.T) {
	actual := ThaiIDValidator("1131812674671")
	expected := false
	if actual != expected {
		t.Fatalf(`ThaiIDValidator("1131812674671") got %v, want %v`, actual, expected)
	}
}

func TestEmpty_ThaiIDValidator(t *testing.T) {
	actual := ThaiIDValidator("")
	expected := false
	if actual != expected {
		t.Fatalf(`ThaiIDValidator("") got %v, want %v`, actual, expected)
	}
}
