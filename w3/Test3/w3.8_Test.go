package main

import "testing"

func TestFib(t *testing.T) {
     type args struct{
		 n int
	 }

	 tests := []struct {
		 name string
		 args args
		 want int
	 }
	for _, tt := range tests {
		t.run(tt.name, func(t *testing.T)){
			if got := fib(tt.args.n); got != tt.want{
				t.Errorf("fib()= %v, want %v", got, tt.want)
			}
		}
	}
	

}
