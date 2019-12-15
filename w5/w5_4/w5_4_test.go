package main

import "testing"

func Testmain(t *testing.T) {
	for _, test := range addResult {
		result := main(a + b)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}

}
