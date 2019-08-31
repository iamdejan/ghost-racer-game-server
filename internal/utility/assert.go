package utility

import (
	"log"
	"strings"
	"testing"
)

func AssertTrue(condition bool, message string, t *testing.T) {
	if condition != true {
		t.Fatal("Assertion fails! Condition is not true! Message:", message)
	}
}

func AssertEquals(object1 interface{}, object2 interface{}, message string, t *testing.T) {
	if object1 != nil {
		log.Printf("assertEquals -> object1: %v, object2: %v", object1, object2)
		AssertTrue(object1 == object2, message, t)
	} else {
		AssertTrue(object2 == nil, message, t)
	}
}

func AssertNotEquals(object1 interface{}, object2 interface{}, message string, t *testing.T) {
	AssertTrue(object1 != object2, message, t)
}

func AssertNotNil(object interface{}, message string, t *testing.T) {
	AssertTrue(object != nil, message, t)
}

func AssertNotBlank(s string, message string, t *testing.T) {
	AssertNotNil(s, message, t)
	s = strings.Trim(s, " ")
	AssertTrue(len(s) > 0, message, t)
}