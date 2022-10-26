package utils

import "testing"

func TestHexToInt(t *testing.T) {
	b1 := [2]byte{0x00, 0x08}
	b2 := [2]byte{0x00, 0x21}
	b3 := [2]byte{0xaf, 0x10}
	b4 := [2]byte{0x10, 0x09}
	exVal1 := 8
	exVal2 := 33
	exVal3 := 44816
	exVal4 := 4105

	actualVal1 := HexToInt(b1)
	if actualVal1 != exVal1 {
		t.Fatal("Values do not match: ", exVal1, " != ", actualVal1)
	}

	actualVal2 := HexToInt(b2)
	if actualVal2 != exVal2 {
		t.Fatal("Values do not match: ", exVal2, " != ", actualVal2)
	}

	actualVal3 := HexToInt(b3)
	if actualVal3 != exVal3 {
		t.Fatal("Values do not match: ", exVal3, " != ", actualVal3)
	}

	actualVal4 := HexToInt(b4)
	if actualVal4 != exVal4 {
		t.Fatal("Values do not match: ", exVal4, " != ", actualVal4)
	}
}
