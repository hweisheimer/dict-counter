package main

import (
	"testing"
)

func expect(expected, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected '%s', Got '%s'", expected, actual)
	}
}

func TestNormalize_RemovesPunctuation(t *testing.T) {
	expected := "passiveaggressive"
	actual := normalize("passive-aggressive")
	expect(expected, actual, t)
}

func TestNormalize_LowersCase(t *testing.T) {
	expected := "lowercase"
	actual := normalize("LowerCase")
	expect(expected, actual, t)
}
