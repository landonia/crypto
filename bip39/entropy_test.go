// Copyright 2017 Landonia Ltd. All rights reserved.

package bip39

import "testing"

func TestOddBitSeed(t *testing.T) {
	if _, err := GenerateRandomEntropy(129); err == nil {
		t.Errorf("Should not be able to generate odd bit size")
		t.Fail()
	}
}

func TestLargeBitSeed(t *testing.T) {
	if _, err := GenerateRandomEntropy(257); err == nil {
		t.Errorf("Should not be able to generate odd bit size")
		t.Fail()
	}
}

func TestSmallBitSeed(t *testing.T) {
	if _, err := GenerateRandomEntropy(127); err == nil {
		t.Errorf("Should not be able to generate odd bit size")
		t.Fail()
	}
}

func TestSmallSeed(t *testing.T) {
	if r, err := GenerateRandomEntropy(128); err != nil {
		t.Errorf("A seed should have been generated")
		t.Fail()
	} else if len(r) != 16 {
		t.Errorf("Seed generated should be 16 bytes")
		t.Fail()
	}
}

func TestLargeSeed(t *testing.T) {
	if r, err := GenerateRandomEntropy(256); err != nil {
		t.Errorf("A seed should have been generated")
		t.Fail()
	} else if len(r) != 32 {
		t.Errorf("Seed generated should be 16 bytes")
		t.Fail()
	} else if r.String() != string(r.ToHex()) {
		t.Errorf("String() should be equal to ToHex()")
		t.Fail()
	}
}
