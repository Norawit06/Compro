package main

import (
	"testing"
)

func Test_math2(t *testing.T) {
	a := math2(7, 9)

	if a != 16 {
		t.Errorf("Want result : 16, but it was %v instead", a)
	}

}

func Test_name(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name(tt.args.str)
		})
	}
}
