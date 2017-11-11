// Copyright 2017 Landonia Ltd. All rights reserved.

package bip39

import (
	"bytes"
	"testing"
)

func TestGeneratedEntropy(t *testing.T) {

	// Test the successful generation of an RandomEntropy
	testGenerateRandomEntropy(128, t)
	testGenerateRandomEntropy(160, t)
	testGenerateRandomEntropy(192, t)
	testGenerateRandomEntropy(224, t)
	testGenerateRandomEntropy(256, t)
}

func testGenerateRandomEntropy(bits int, t *testing.T) {

	// Test the successful generation of an RandomEntropy
	r, err := GenerateRandomEntropy(bits)
	if err != nil {
		t.Errorf("The GenerateRandomEntropy with length of %d bits has failed", bits)
		t.Fail()
	}

	m, err := r.GenerateMnemonics(English)
	if err != nil {
		t.Errorf("Could not generate Mnemonics for %d bits entropy", bits)
		t.Fail()
	} else if len(m) != (bits+(bits/32))/11 {
		t.Errorf("Incorrect number of words generated. Expecting %d and got %d", (bits+(bits/32))/11, len(m))
		t.Fail()
	}
}

func TestGenerateSuccessfulChecksum(t *testing.T) {
	r := RandomEntropy([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	// Check that the checksum generated is correct for 128 bits
	checksumTest(r, 132, false, t)

	r = append(r, 17, 18, 19, 20)
	checksumTest(r, 165, false, t)

	r = append(r, 21, 22, 23, 24)
	checksumTest(r, 198, false, t)

	r = append(r, 25, 26, 27, 28)
	checksumTest(r, 231, false, t)

	r = append(r, 29, 30, 31, 32)
	checksumTest(r, 264, false, t)
}

func TestGenerateFailedChecksum(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	r := RandomEntropy(data)

	// Check that the checksum fails as it has been passed an incorrect length
	checksumTest(r, 112, true, t)

	r = append(r, 15, 16, 17, 18)
	checksumTest(r, 144, true, t)
}

func checksumTest(r RandomEntropy, expectedSize int, expectedFail bool, t *testing.T) {
	if checksum, err := r.generateChecksum(); err != nil && !expectedFail {
		t.Errorf("The checksum has returned an unexpected error: %s", err.Error())
		t.Fail()
	} else if !expectedFail && expectedSize != len(checksum) {
		t.Errorf("Expected %d bits and got %d bits", expectedSize, len(checksum))
		t.Fail()
	}
}

func TestGenerateMnemonics(t *testing.T) {
	r := RandomEntropy([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	mnemonicTest(English, r, 12, false, t)

	r = append(r, 17, 18, 19, 20)
	mnemonicTest(English, r, 15, false, t)

	r = append(r, 21, 22, 23, 24)
	mnemonicTest(English, r, 18, false, t)

	r = append(r, 25, 26, 27, 28)
	mnemonicTest(Spanish, r, 21, false, t)

	r = append(r, 29, 30, 31, 32)
	mnemonicTest(Spanish, r, 24, false, t)
}

func TestCorrectMnemonics(t *testing.T) {
	hex := RandomEntropyHex("000000000000000000000000000000000")
	r, _ := hex.ToRandomEntropy()
	_, err := mnemonicTest(English, r, 12, true, t)
	if err == nil {
		t.Fail()
	}

	hex = RandomEntropyHex("00000000000000000000000000000000")
	r, _ = hex.ToRandomEntropy()
	words, _ := mnemonicTest(English, r, 12, false, t)
	expected := Mnemonics{"abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "about"}
	compareWords(words, expected, t)
	compareSeeds(words, expected, t)

	hex = RandomEntropyHex("ffffffffffffffffffffffffffffffff")
	r, _ = hex.ToRandomEntropy()
	words, _ = mnemonicTest(English, r, 12, false, t)
	expected = Mnemonics{"zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "wrong"}
	compareWords(words, expected, t)
	compareSeeds(words, expected, t)

	hex = RandomEntropyHex("0c1e24e5917779d297e14d45f14e1a1a")
	r, _ = hex.ToRandomEntropy()
	words, _ = mnemonicTest(English, r, 12, false, t)
	expected = Mnemonics{"army", "van", "defense", "carry", "jealous", "true", "garbage", "claim", "echo", "media", "make", "crunch"}
	compareWords(words, expected, t)
	compareSeeds(words, expected, t)
}

func mnemonicTest(lang Language, r RandomEntropy, expectedSize int, expectedFail bool, t *testing.T) (Mnemonics, error) {
	mnemonics, err := r.GenerateMnemonics(lang)
	if err != nil && !expectedFail {
		t.Errorf("The function has returned an unexpected error: %s", err.Error())
		t.Fail()
	} else if !expectedFail && expectedSize != len(mnemonics) {
		t.Errorf("Expected %d words but got %d words", expectedSize, len(mnemonics))
		t.Fail()
	}
	return mnemonics, err
}

func compareWords(words, expected Mnemonics, t *testing.T) {
	if len(words) != len(expected) {
		t.Errorf("Expected %d words but got %d words", len(words), len(expected))
		t.Fail()
	}
	for i := 0; i < len(words); i++ {
		if words[i] != expected[i] {
			t.Errorf("Expected %s but got %s", words[i], expected[i])
			t.Fail()
		}
	}

	if words.String() != expected.String() {
		t.Errorf("Expected '%s' but got '%s'", expected.String(), words.String())
		t.Fail()
	}
}

// Ensure that the seeds match the expected type
func compareSeeds(words, expected Mnemonics, t *testing.T) {
	got := words.GenerateSeed("")
	want := expected.GenerateSeed("")
	if !bytes.Equal(got, want) {
		t.Errorf("Expected seed '%x' but got '%x'", want, got)
		t.Fail()
	}

	// Make sure that the Hex values match
	gotHex := got.ToHex()
	wantHex := want.ToHex()
	if gotHex != wantHex {
		t.Errorf("Expected seed '%x' but got '%x'", want, got)
		t.Fail()
	}

	// Ensure that they convert back to the correct seeds
	gotSeed, err := gotHex.ToSeed()
	if err != nil {
		t.Fail()
	} else if !bytes.Equal(gotSeed, got) {
		t.Errorf("Expected seed '%x' but got '%x'", got, gotSeed)
		t.Fail()
	} else if gotSeed.String() != got.String() {
		t.Errorf("Expected seed string '%s' but got '%s'", got.String(), gotSeed.String())
		t.Fail()
	}

	wantSeed, err := wantHex.ToSeed()
	if err != nil {
		t.Fail()
	} else if !bytes.Equal(wantSeed, want) {
		t.Errorf("Expected seed '%x' but got '%x'", want, wantSeed)
		t.Fail()
	} else if wantSeed.String() != want.String() {
		t.Errorf("Expected seed string '%s' but got '%s'", want.String(), wantSeed.String())
		t.Fail()
	}
}
