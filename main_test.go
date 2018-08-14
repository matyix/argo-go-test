package main

import "testing"

func TestHelloArgo(t *testing.T) {
	if helloargo() != "Hello Argo!!" {
		t.Fatal("Test fail")
	}
}
