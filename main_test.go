package main

import "testing"

func TestHistory(t *testing.T) {
	_, err := histories()
	if err != nil {
		t.Fatal(err)
	}
}

func TestExcd(t *testing.T) {
	err := excd("C:\\Windows\\")
	if err != nil {
		t.Fatal(err)
	}
}
