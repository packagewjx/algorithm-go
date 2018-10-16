package study

import (
	"testing"
)

func TestAlphabetChange(t *testing.T) {
	testCases := []struct {
		text   string
		result string
	}{
		{"ABC", "ABC"},
		{"abc", "ABC"},
		{"a", "A"},
		{"A", "A"},
		{"", ""},
		{"abcd", "ABCD"},
		{"abcde", "ABCDE"},
		{"abCD", "ABCD"},
		{"@ABCDEFGHIJKLMNOPQRSTUVWXYZ_`abcdefghijklmnopqrstuvwxyz{|}", "@ABCDEFGHIJKLMNOPQRSTUVWXYZ_`ABCDEFGHIJKLMNOPQRSTUVWXYZ{|}"}}

	for _, val := range testCases {
		upperCase := toUpperCase(val.text)
		if upperCase != val.result {
			t.Error(upperCase, val.result)
		}
	}
}

func TestByteToUint32Array(t *testing.T) {

	testCase := []struct {
		array  []byte
		expect []uint32
	}{
		{[]byte{1, 2, 3, 4}, []uint32{0x01020304}},
		{[]byte{1, 2, 3}, []uint32{0x01020300}},
		{[]byte{1, 2}, []uint32{0x01020000}},
		{[]byte{1}, []uint32{0x01000000}},
		{[]byte{1, 2, 3, 4, 5}, []uint32{0x01020304, 0x05000000}},
		{[]byte{1, 2, 3, 4, 5, 6, 7, 8}, []uint32{0x01020304, 0x05060708}},
		{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []uint32{0x01020304, 0x05060708, 0x090A0000}},
		{[]byte{1, 2, 3, 4, 5, 6}, []uint32{0x01020304, 0x05060000}},
		{[]byte{1, 2, 3, 4, 5, 6, 7}, []uint32{0x01020304, 0x05060700}},
	}

	for _, val := range testCase {
		result := byteArrayToIntArray(val.array)

		if len(result) != len(val.expect) {
			t.Error("长度不对")
		}

		for i := 0; i < len(result); i++ {
			if val.expect[i] != result[i] {
				t.Error("第", i, "位不对")
			}
		}
	}
}
