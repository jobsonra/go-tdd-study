package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola()
	expected := "Olá Mundo!"

	if result != expected {
		t.Errorf("the result %s is different from the expected %s", result, expected)
	}
}
