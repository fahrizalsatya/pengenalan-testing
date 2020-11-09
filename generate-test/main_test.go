package main

import "testing"

func TestAvg(t *testing.T) {
	type args struct {
		nos []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.nos...); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}
