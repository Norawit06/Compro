package main

import "testing"

func TestHello(t *testing.T) {

	result := Hello("Per") // Should return "hello Per!"

	if result != "heloo Per!"
}

