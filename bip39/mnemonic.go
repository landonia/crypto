// Copyright 2017 Landonia Ltd. All rights reserved.

package bip39

import (
	"crypto/sha256"
	"strconv"
	"strings"
)

// Mnemonic holds a word string
type Mnemonic string

// Mnemonics is a simple wrapper for a string slice
type Mnemonics []Mnemonic

// JoinWords output the Mnemonics as a HEX string
func (m Mnemonics) JoinWords() string {
	str := ""
	for i, mnemonic := range m {
		str += string(mnemonic)
		if i < len(m)-1 {
			str += " "
		}
	}
	return str
}

// String will default to the HEX value (see #ToHex)
func (m Mnemonics) String() string {
	return m.JoinWords()
}

// GenerateSeed will create a new Seed to use within a HD wallet
func (m Mnemonics) GenerateSeed(passphrase string) Seed {
	return GenerateSeed(m.JoinWords(), passphrase)
}

// Checksum holds SHA256 hash of the entropy
type Checksum [sha256.Size]byte

// ToBitSlice will convert the byte slice into an individual bit slice
// with each element holding one bit
func (c Checksum) ToBitSlice() BitSlice {

	// Create a new bitslice of the correct length
	return convertToBitSlice(c[:])
}

// GenerateMnemonics will generate the set of mnemonic words (length dependent
// upon bitsize). An error will be returned if the RandomEntropy or Language
// are invalid
// Word Count:
// 128 = 12 words
// 160 = 15 words
// 192 = 18 words
// 224 = 21 words
// 256 = 24 words
func GenerateMnemonics(r RandomEntropy, lang Language) (Mnemonics, error) {

	// Generate the checkSum
	checksum, err := r.generateChecksum()
	if err != nil {
		return nil, err
	}
	const blockLength = 11

	// We need to split the word count into sections of 11 to get the index for the word
	mnemonicCount := len(checksum) / blockLength
	mnemonics := make([]Mnemonic, 0, mnemonicCount)

	// Now for each word we need to get the index (by creating a bit string)
	for i := 0; i < mnemonicCount; i++ {
		start := i * blockLength

		bitString := strings.Join(checksum[start:start+blockLength], "")
		index, err := strconv.ParseInt(bitString, 2, 16)
		if err != nil {
			return nil, err
		}

		var mnemonic Mnemonic
		mnemonic, err = GetWord(lang, int(index))
		if err != nil {
			return nil, err
		}
		mnemonics = append(mnemonics, mnemonic)
	}
	return mnemonics, nil
}
