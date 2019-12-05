package main

import "testing"

func TestHello(t *testing.T) {

	result := Hello("Per") // Should return "hello Per!"

	if result != "heloo Per!"{
		t.Errorf("hello(\"Per\") failed, expected %v, got %v, ""hello man!", result) 
	} else {
		t.Logf("hello(\"hello Per!\") success,
	}
}