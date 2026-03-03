package main

import "testing"

func TestHexConversion(t *testing.T) {
    got := hexToDec("1E")
    want := "30"
    if got != want {
        t.Errorf("hexToDec(1E) = %s; want %s", got, want)
    }
}

func TestBinConversion(t *testing.T) {
    got := binToDec("10")
    want := "2"
    if got != want {
        t.Errorf("binToDec(10) = %s; want %s", got, want)
    }
}

func TestUpperCase(t *testing.T) {
    words := []string{"hello"}
    got := applyCase(words, "up", "(up)")
    want := "HELLO"
    if got[0] != want {
        t.Errorf("applyCase(up) = %s; want %s", got[0], want)
    }
}

func TestLowerCase(t *testing.T) {
    words := []string{"HELLO"}
    got := applyCase(words, "low", "(low)")
    want := "hello"
    if got[0] != want {
        t.Errorf("applyCase(low) = %s; want %s", got[0], want)
    }
}

func TestCapitalize(t *testing.T) {
    words := []string{"bridge"}
    got := applyCase(words, "cap", "(cap)")
    want := "Bridge"
    if got[0] != want {
        t.Errorf("applyCase(cap) = %s; want %s", got[0], want)
    }
}

func TestArticles(t *testing.T) {
    text := "a apple a house a car"
    got := fixArticles(text)
    want := "an apple an house a car"
    if got != want {
        t.Errorf("fixArticles() = %s; want %s", got, want)
    }
}
