package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize_RemovesPunctuation(t *testing.T) {
	expected := "eatsshootsleaves"
	actual := normalize("eats,shoots&leaves-")
	assert.Equal(t, expected, actual)
}

func TestNormalize_LowersCase(t *testing.T) {
	expected := "lowercase"
	actual := normalize("LowerCase")
	assert.Equal(t, expected, actual)
}

func TestNormalize_RemovesSimpleAccents(t *testing.T) {
	expected := "e"
	actual := normalize("\u00e9")
	assert.Equal(t, expected, actual)
}

func TestNormalize_RemovesCursedAccents(t *testing.T) {
	expected := "he comes"
	actual := normalize("h̺̼̞̼͇̮̖̭̗̳̳̣̜̦̬̟̻̄͐͗̎͂ͤ̄̌͆͂ͩ͑̿͛̏͂̇̚e͓͖̰̹̯̬͙̼͇̊ͯͫ̈̊ͩ̔ͣͤ̾͂ ̮̭̙̂ͪ̏̿ͫ̇̐̆͗̐͂ͮͣ̂c͔̪̣͊͋͑̆ͪͯ̍ͩ̎͌͛͋̆͑͗ͅo͍̭̟͎͓̹̖͔̱̼͉̪̪͕͖̭͐̇ͤͯ͛͂͛̅̔̓̋͒̊̐ͩm̯̭͖͚͇̯̠̫͔̼͔̟̯̪̲͛͐̈̃̀̈́́ͨ̽̔̏ͪ̅͐͐͗̂ͮ̔ê͎͚͎͇̣̟̺͇̲͉̱̫ͬ̒̐̉ͥ̐ͭͭͫ̔͐̈́ͨ͑s͉̫̥̬̠̤̭̙̿̑̃̾͒̌ͧ͛̍̚")
	assert.Equal(t, expected, actual)
}

func TestNormalize_RemovesPossessives(t *testing.T) {
	actual := normalize("parson's\nperson\nperson's\r\npersonal")
	assert.NotContains(t, actual, "parson's")
	assert.NotContains(t, actual, "person's")
}

func TestNormalize_RemovesWhitespace(t *testing.T) {
	expected := "correcthorsebatterystaple"
	actual := normalize(" correct\rhorse\n battery\r\n\tstaple")
	assert.Equal(t, expected, actual)
}

func TestCountCharacters_ContainsAllKeys(t *testing.T) {
	expectedKeys := []byte{'a', 'b', 'c'}
	actualCounts := countCharacters("bbbaacccc")
	actualKeys := []byte{}
	for key := range actualCounts {
		actualKeys = append(actualKeys, key)
	}
	assert.ElementsMatch(t, expectedKeys, actualKeys, "Character count keys do not match")
}
