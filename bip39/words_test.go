// Copyright 2017 Landonia Ltd. All rights reserved.

package bip39

import "testing"

func TestGetWrongIndexWord(t *testing.T) {
	if _, err := GetWord(English, 2049); err == nil {
		t.Errorf("The index should be within range: 0 <= index < 2048")
		t.Fail()
	}
}

func TestGetAllLanguageWord(t *testing.T) {

	// Try to get the correct word for each language
	if word, _ := GetWord(English, 1); word != "ability" {
		t.Errorf("The expected English word is ability")
		t.Fail()
	}

	if word, _ := GetWord(Spanish, 2); word != "abeja" {
		t.Errorf("The expected Spanish word is abeja")
		t.Fail()
	}

	if word, _ := GetWord(French, 3); word != "abeille" {
		t.Errorf("The expected French word is abeille")
		t.Fail()
	}

	if word, _ := GetWord(Italian, 4); word != "abisso" {
		t.Errorf("The expected Italian word is abisso")
		t.Fail()
	}

	if word, _ := GetWord(Japanese, 5); word != "あきる" {
		t.Errorf("The expected Japanese word is あきる")
		t.Fail()
	}

	if word, _ := GetWord(Korean, 6); word != "가뭄" {
		t.Errorf("The expected Korean word is 가뭄")
		t.Fail()
	}

	if word, _ := GetWord(ChineseSimple, 7); word != "和" {
		t.Errorf("The expected ChineseSimple word is 和")
		t.Fail()
	}

	if word, _ := GetWord(ChineseTraditional, 8); word != "人" {
		t.Errorf("The expected ChineseTraditional word is 人")
		t.Fail()
	}
}

func TestGetMissingLanguageWord(t *testing.T) {

	// Try a different language tat doesnt exist
	if _, err := GetWord(Language("german"), 2047); err == nil {
		t.Errorf("The german language should not be accepted")
		t.Fail()
	}
}
