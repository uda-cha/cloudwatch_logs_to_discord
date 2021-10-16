package main

import (
	"strings"
	"testing"
)

func TestReconstructSlicesforDiscordLimit1(t *testing.T) {
	msgSlice := []string{"aaa", "bbb"}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{"aaa\rbbb"}

	if len(actual) != len(expected) {
		t.Errorf("shortMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("shortMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
}

func TestReconstructSlicesforDiscordLimit2(t *testing.T) {
	str := "a"
	msgSlice := []string{strings.Repeat(str, 1900), strings.Repeat(str, 200), "bbb"}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{strings.Repeat(str, 1900), strings.Repeat(str, 200) + "\rbbb"}

	if len(actual) != len(expected) {
		t.Errorf("longMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
	if actual[1] != expected[1] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[1], expected[1])
	}
}

func TestReconstructSlicesforDiscordLimit3(t *testing.T) {
	str := "a"
	msgSlice := []string{strings.Repeat(str, 200), strings.Repeat(str, 1900), "bbb"}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{strings.Repeat(str, 200), strings.Repeat(str, 1900) + "\rbbb"}

	if len(actual) != len(expected) {
		t.Errorf("longMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
	if actual[1] != expected[1] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[1], expected[1])
	}
}

func TestReconstructSlicesforDiscordLimit4(t *testing.T) {
	str := "a"
	msgSlice := []string{strings.Repeat(str, 1900), strings.Repeat(str, 1900), strings.Repeat(str, 1900)}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{strings.Repeat(str, 1900), strings.Repeat(str, 1900), strings.Repeat(str, 1900)}

	if len(actual) != len(expected) {
		t.Errorf("longMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
	if actual[1] != expected[1] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[1], expected[1])
	}
}

func TestReconstructSlicesforDiscordLimit5(t *testing.T) {
	str := "a"
	msgSlice := []string{strings.Repeat(str, 2200), strings.Repeat(str, 1900), strings.Repeat(str, 1900)}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{strings.Repeat(str, 2200), strings.Repeat(str, 1900), strings.Repeat(str, 1900)}

	if len(actual) != len(expected) {
		t.Errorf("longMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
	if actual[1] != expected[1] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[1], expected[1])
	}
}

func TestReconstructSlicesforDiscordLimit6(t *testing.T) {
	str := "a"
	msgSlice := []string{strings.Repeat(str, 1900), strings.Repeat(str, 200), strings.Repeat(str, 1900)}
	actual := ReconstructSlicesforDiscordLimit(msgSlice)
	expected := []string{strings.Repeat(str, 1900), strings.Repeat(str, 200), strings.Repeat(str, 1900)}

	if len(actual) != len(expected) {
		t.Errorf("longMsg length got: %v\nwant: %v", len(actual), len(expected))
	}
	if actual[0] != expected[0] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[0], expected[0])
	}
	if actual[1] != expected[1] {
		t.Errorf("longMsg first got: %v\nwant: %v", actual[1], expected[1])
	}
}
