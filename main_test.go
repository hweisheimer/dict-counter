package main

import (
	"strings"
	"testing"
)

func expect(expected, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected '%s', Got '%s'", expected, actual)
	}
}

func TestNormalize_RemovesPunctuation(t *testing.T) {
	expected := "eatsshootsleaves"
	actual := normalize("eats,shoots&leaves-")
	expect(expected, actual, t)
}

func TestNormalize_LowersCase(t *testing.T) {
	expected := "lowercase"
	actual := normalize("LowerCase")
	expect(expected, actual, t)
}

func TestNormalize_RemovesSimpleAccents(t *testing.T) {
	expected := "e"
	actual := normalize("\u00e9")
	expect(expected, actual, t)
}

func TestNormalize_RemovesCursedAccents(t *testing.T) {
	expected := "he comes"
	actual := normalize("h̺̼̞̼͇̮̖̭̗̳̳̣̜̦̬̟̻̄͐͗̎͂ͤ̄̌͆͂ͩ͑̿͛̏͂̇̚e͓͖̰̹̯̬͙̼͇̊ͯͫ̈̊ͩ̔ͣͤ̾͂ ̮̭̙̂ͪ̏̿ͫ̇̐̆͗̐͂ͮͣ̂c͔̪̣͊͋͑̆ͪͯ̍ͩ̎͌͛͋̆͑͗ͅo͍̭̟͎͓̹̖͔̱̼͉̪̪͕͖̭͐̇ͤͯ͛͂͛̅̔̓̋͒̊̐ͩm̯̭͖͚͇̯̠̫͔̼͔̟̯̪̲͛͐̈̃̀̈́́ͨ̽̔̏ͪ̅͐͐͗̂ͮ̔ê͎͚͎͇̣̟̺͇̲͉̱̫ͬ̒̐̉ͥ̐ͭͭͫ̔͐̈́ͨ͑s͉̫̥̬̠̤̭̙̿̑̃̾͒̌ͧ͛̍̚")
	expect(expected, actual, t)
}

func TestNormalize_RemovesPossessives(t *testing.T) {
	actual := normalize("parson's\nperson\nperson's\r\npersonal")
	if strings.Contains(actual, "parson's") {
		t.Error("Result should not contain \"parson's\"")
	}
	if strings.Contains(actual, "person's") {
		t.Error("Result should not contain \"person's\"")
	}
}

func TestNormalize_RemovesWhitespace(t *testing.T) {
	expected := "correcthorsebatterystaple"
	actual := normalize(" correct\rhorse\n battery\r\n\tstaple")
	expect(expected, actual, t)
}
