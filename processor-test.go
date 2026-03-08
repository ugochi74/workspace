package main

import "testing"

func TestConvertNumbers(t *testing.T) {

	input := "10 (bin)"
	expected := "2"

	result := ConvertNumbers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestArticles(t *testing.T) {

	input := "a apple"
	expected := "an apple"

	result := FixArticles(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}