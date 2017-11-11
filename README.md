# crypto

[![Go Report Card](https://goreportcard.com/badge/github.com/landonia/crypto)](https://goreportcard.com/report/github.com/landonia/crypto)
[![GoDoc](https://godoc.org/github.com/landonia/crypto?status.svg)](https://godoc.org/github.com/landonia/crypto)

Provides crypto packages allowing you to create crypto wallets.

## Overview

Provides implementations of BIP standards for creating crypto wallets/applications.
It will also contain other code that doesnt follow any BIP standard, but can be useful.
I'm implementing my own crypto apps and I thought that implementing each part myself
helps me to learn how it all works.

### BIP-39

An implementation of a mnemonic code or mnemonic sentence for the generation of deterministic wallets (BIP-39). It will create (or recreate from a group of easy to remember words.

It consists two main parts:

1. Generate the mnemonic words using the provided key and convert it into a binary seed. This seed can be later used to generate deterministic wallets using BIP-0032 or similar methods. The accepted range of entropy is 128 - 256 bits.
2. Provide the seed generator that takes the mnemonics (and optional passphrase) to use within deterministic wallets.

#### BIP-39 Example
```go
  package main

  import (
    "flag"
    "github.com/landonia/crypto/bip39"
  )

  func main() {

    // You can generate a new random entropy - acceptable bits 128 - 256 increments of 32
    entropy, err := bip39.GenerateRandomEntropy(bits)

    // RandomEntropy is a wrapper that provides some useful helpers
    entropy.ToHex() // will return a Hex representation of the entropy
    entropy.String() // calls the ToHex() function and is there to print pretty strings

    // You can also regenerate the random entropy from an existing Hexadecimal string
    hex := RandomEntropyHex("0c1e24e5917779d297e14d45f14e1a1a")
    entropy, err = hex.ToRandomEntropy()

    // Using a valid entropy you can generate the mnemonic words using one of the supplied
    // languages (English, Spanish, French, Italian, Japanese, Korean, ChineseSimple, ChineseTraditional)
    mnemonics, err := entropy.GenerateMnemonics(bip39.English)

    // mnemonics is a wrapper for an array of strings. It provides some handy functions..
    mnemonics.JoinWords() // will create the real mnemonics string
    mnemonics.String() // calls the JoinWords() function and is there to print pretty string

    // From here you can generate the Seed to use in a deterministic wallet directly from the Mnemonic
    seed := mnemonics.GenerateSeed("passphrase is optional")

    // If you just have the mnemonic string of words and passphrase you can generate the Seed
    seed = bip39.GenerateSeed("army van defense carry jealous true garbage claim echo media make crunch", "optional passphrase")

    // Seed is wrapper around a byte array that contains the 512bit seed
    seed.ToHex() // will output the seed as a Hexadecimal string
    seed.String() // calls the ToHex() function and is there to print pretty strings

    // If you already have the Hexadecimal seed string you can regenerate the Seed easily
    hex = SeedHex("5b56c417303faa3fcba7e57400e120a0ca83ec5a4fc9ffba757fbe63fbd77a89a1a3be4c67196f57c39a88b76373733891bfaba16ed27a813ceed498804c0570")

    // Then easily recreate the Seed (which is a []byte wrapper)
    seed = hex.ToSeed()
  }
```

## Installation

With a healthy Go Language installed, simply run `go get github.com/landonia/crypto`

## About

goproxwas written by [Landon Wainwright](http://www.landotube.com) | [GitHub](https://github.com/landonia).

Follow me on [Twitter @landotube](http://www.twitter.com/landotube)! Although I don't really tweet much tbh.
