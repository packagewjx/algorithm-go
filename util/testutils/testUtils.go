package testutils

import "testing"

func AssertTrue(expression bool, errorMessage string, t *testing.T) {
	if !expression {
		t.Error(errorMessage)
	}
}

func AssertNoErr(err error, errorMessage string, t *testing.T) {
	if err != nil {
		t.Error(errorMessage, ". Error is ", err)
	}
}
