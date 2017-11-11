// Copyright 2017 Landonia Ltd. All rights reserved.

package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// RandomEntropy wraps a random series of bytes
type RandomEntropy []byte

// toBitSlice will convert the byte slice into an individual bit slice
// with each element holding one bit
func (r RandomEntropy) toBitSlice() BitSlice {

	// Create a new bitslice of the correct length
	return convertToBitSlice(r)
}

// String will print the Entropy as a string
func (r RandomEntropy) String() string {
	return string(r.ToHex())
}

// ToHex will create a HEX string representation
func (r RandomEntropy) ToHex() RandomEntropyHex {
	return RandomEntropyHex(hex.EncodeToString(r))
}

// GenerateMnemonics will return the matching words for the provided entropy
// The number of words depends on the entropy bit size
func (r RandomEntropy) GenerateMnemonics(lang Language) (Mnemonics, error) {
	return GenerateMnemonics(r, lang)
}

// generateChecksum will SHA256 the entropy and add the first bytes of the checksum to
// the return slice. The number of bits added = ENT / 32
// 128 = 4 bits = 12 words
// 160 = 5 bits = 15 words
// 192 = 6 bits = 18 words
// 224 = 7 bits = 21 words
// 256 = 8 bits = 24 words
func (r RandomEntropy) generateChecksum() (BitSlice, error) {

	// Test that the data is valid
	if len(r) < 16 || len(r) > 32 {
		return nil, fmt.Errorf("16 <= entropy <= 256")
	} else if len(r)*8%32 != 0 {
		return nil, fmt.Errorf("entropy needs to be divisible by 32")
	}

	// Use the SHA256 hash to get a checksum for the current entropy
	hash := Checksum(sha256.Sum256(r))

	// Convert to get slices of the bits
	seedBits := r.toBitSlice()
	checksumBits := hash.ToBitSlice()

	for i := 0; i < len(seedBits)/32; i++ {
		seedBits = append(seedBits, checksumBits[i])
	}
	return seedBits, nil
}

// GenerateRandomEntropy will generate a random number of bytes matching the length
// of bits required. The length of bits must be divisible by 32 or else
// a corresponding error will be returned.
func GenerateRandomEntropy(bits int) (RandomEntropy, error) {
	if bits < 128 || bits > 256 {
		return nil, errors.New("expected bit size 128 <= bits <= 256")
	} else if bits%32 != 0 {
		return nil, errors.New("bits count must be in divisible by 32")
	}

	// Generate a new byte array to generate some random bytes
	b := make([]byte, bits/8)
	_, err := rand.Read(b)
	return b, err
}

// Seed generated from the mnemonics
type Seed []byte

// GenerateSeed will use the PBKDF2 to generate a hash (using the SHA512 hash)
// and performing 2048 iterations.
func GenerateSeed(password, salt string) Seed {
	return pbkdf2.Key([]byte(password), []byte("mnemonic"+salt), 2048, 64, sha512.New)
}

// String will print the Entropy as a string
func (seed Seed) String() string {
	return string(seed.ToHex())
}

// ToHex will create a Hex string representation
func (seed Seed) ToHex() SeedHex {
	return SeedHex(hex.EncodeToString(seed))
}

// RandomEntropyHex will hold a Random hex value
type RandomEntropyHex string

// ToRandomEntropy will create the Random entropy from the Hex value
func (h RandomEntropyHex) ToRandomEntropy() (RandomEntropy, error) {
	bytes, err := hex.DecodeString(string(h))
	if err != nil {
		return nil, err
	}
	return RandomEntropy(bytes), nil
}

// SeedHex will hold a Seed hex value
type SeedHex string

// ToSeed will recreate the Seed from the HEX value
func (h SeedHex) ToSeed() (Seed, error) {
	bytes, err := hex.DecodeString(string(h))
	if err != nil {
		return nil, err
	}
	return Seed(bytes), nil
}

// BitSlice will hold the individual bits for a slice
type BitSlice []string

// convertToBitSlice will convert a byte slice
func convertToBitSlice(conv []byte) BitSlice {
	// Create a new bitslice of the correct length
	bits := make(BitSlice, 0, len(conv)*8)
	for _, b := range conv {
		for j := 0; j < 8; j++ {
			if b&(1<<byte(7-j)) != 0 {
				bits = append(bits, "1")
			} else {
				bits = append(bits, "0")
			}
		}
	}
	return bits
}
