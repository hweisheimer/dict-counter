package main

import (
	"strings"
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
	expected := "angstrome"
	actual := normalize("Ångström\u00e9")
	assert.Equal(t, expected, actual)
}

func TestNormalize_RemovesCursedAccents(t *testing.T) {
	expected := "hecomes"
	actual := normalize("h̺̼̞̼͇̮̖̭̗̳̳̣̜̦̬̟̻̄͐͗̎͂ͤ̄̌͆͂ͩ͑̿͛̏͂̇̚e͓͖̰̹̯̬͙̼͇̊ͯͫ̈̊ͩ̔ͣͤ̾͂c͔̪̣͊͋͑̆ͪͯ̍ͩ̎͌͛͋̆͑͗ͅo͍̭̟͎͓̹̖͔̱̼͉̪̪͕͖̭͐̇ͤͯ͛͂͛̅̔̓̋͒̊̐ͩm̯̭͖͚͇̯̠̫͔̼͔̟̯̪̲͛͐̈̃̀̈́́ͨ̽̔̏ͪ̅͐͐͗̂ͮ̔ê͎͚͎͇̣̟̺͇̲͉̱̫ͬ̒̐̉ͥ̐ͭͭͫ̔͐̈́ͨ͑s͉̫̥̬̠̤̭̙̿̑̃̾͒̌ͧ͛̍̚")
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

func TestCountCharacters_AccumulatesAllKeys(t *testing.T) {
	expectedKeys := []rune{'a', 'b', 'c'}
	actualCounts := countCharacters("bbbaacccc")
	actualKeys := []rune{}
	for key := range actualCounts {
		actualKeys = append(actualKeys, key)
	}
	assert.ElementsMatch(t, expectedKeys, actualKeys, "Character count keys do not match")
}

func TestCountCharacters_AccumulatesAllOccurences(t *testing.T) {
	expectedCounts := map[rune]uint32{
		'a': 2,
		'b': 3,
		'c': 4,
	}
	actualCounts := countCharacters("bbbaacccc")
	assert.InDeltaMapValues(t, expectedCounts, actualCounts, 0, "Character count values do not match")
}

func TestNormalizeAndCountCharacters_CountsEmbellishedCharactesSanely(t *testing.T) {
	expectedCounts := map[rune]uint32{
		'a': 4,
		'n': 2,
		'd': 1,
		'w': 1,
		'e': 2,
		'b': 2,
		'y': 1,
		'g': 1,
		's': 1,
		't': 1,
		'r': 1,
		'o': 1,
		'm': 1,
	}
	actualCounts := countCharacters(normalize("and a w\u00e9e baby Ångström"))
	assert.InDeltaMapValues(t, expectedCounts, actualCounts, 0, "Character count values do not match")
}

func TestBuildHistogram_SortsKeys(t *testing.T) {
	counts := map[rune]uint32{
		'b': 12345,
		'a': 54321,
		'c': 200,
	}
	rawPlot := buildHistogram(counts, 10)
	plotRows := strings.Split(rawPlot, "\n")

	countPassed := assert.Len(t, plotRows, 3)
	if countPassed {
		assert.Equal(t, 'a', rune(plotRows[0][0]))
		assert.Equal(t, 'b', rune(plotRows[1][0]))
		assert.Equal(t, 'c', rune(plotRows[2][0]))
	}
}

func TestBuildHistogram_ScalesToWidth(t *testing.T) {
	counts := map[rune]uint32{
		'b': 200,
		'o': 300,
		'p': 600,
	}
	expectedPlot := `b |>>> (200)
o |>>>>> (300)
p |>>>>>>>>>> (600)`
	rawPlot := buildHistogram(counts, 10)
	assert.Equal(t, expectedPlot, rawPlot)
}
