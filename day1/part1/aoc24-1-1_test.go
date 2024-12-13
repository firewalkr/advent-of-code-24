package main

import "testing"

func Test_sumAbsDiffs(t *testing.T) {
	type args struct {
		listOne []int
		listTwo []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"aoc sample", args{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAbsDiffs(tt.args.listOne, tt.args.listTwo); got != tt.want {
				t.Errorf("sumAbsDiffs() = %v, want %v", got, tt.want)
			}
		})
	}
}