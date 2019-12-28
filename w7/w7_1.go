package w7

import "fmt"

var flagtests = []struct {
	in  string
	out string
	}{
		{"%a", "[%a]"},
		{"%-a", "[%-a]"},
		{"%+a", "[%+a]"},
		{"%#a", "[%#a]"},
		{"% a", "[% a]"},
		{"%0a", "[%0a]"},
		{"%1.2a", "[%1.2a]"},
		{"%-1.2a", "[%-1.2a]"},
		{"%+1.2a", "[%+1.2a]"},
		{"%-+1.2a", "[%+-1.2a]"},
		{"%-+1.2abc", "[%+-1.2a]bc"},
		{"%-1.2abc", "[%-1.2a]bc"},	
