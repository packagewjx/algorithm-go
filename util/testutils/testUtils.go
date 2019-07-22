package testutils

import "testing"

func AssertTrue(expression bool, errorMessage string, t *testing.T) {
	if !expression {
		t.Error(errorMessage)
	}
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}, msg ...string) {
	if expected != actual {
		if len(msg) == 0 {
			t.Error("Expected:", expected, "Actual:", actual)
		} else {
			t.Error(msg)
		}
	}
}

func AssertNoErr(err error, errorMessage string, t *testing.T) {
	if err != nil {
		t.Error(errorMessage, ". Error is ", err)
	}
}
