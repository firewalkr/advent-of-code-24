package main

import (
	"reflect"
	_sort "sort"
	"testing"
)

var aocSample1 = StoneGroups{
	StoneGroup{0, 1},
	StoneGroup{1, 1},
	StoneGroup{10, 1},
	StoneGroup{99, 1},
	StoneGroup{999, 1},
}

var aocSample1Result = StoneGroups{
	StoneGroup{1, 2},
	StoneGroup{2024, 1},
	StoneGroup{0, 1},
	StoneGroup{9, 2},
	StoneGroup{2021976, 1},
}

var aocSample2 = StoneGroups{
	StoneGroup{1036288, 1},
	StoneGroup{7, 1},
	StoneGroup{2, 1},
	StoneGroup{20, 1},
	StoneGroup{24, 1},
	StoneGroup{4048, 2},
	StoneGroup{1, 1},
	StoneGroup{8096, 1},
	StoneGroup{28, 1},
	StoneGroup{67, 1},
	StoneGroup{60, 1},
	StoneGroup{32, 1},
}

var aocSample2Result = StoneGroups{
	StoneGroup{2097446912, 1},
	StoneGroup{14168, 1},
	StoneGroup{4048, 1},
	StoneGroup{2, 4},
	StoneGroup{0, 2},
	StoneGroup{4, 1},
	StoneGroup{40, 2},
	StoneGroup{48, 2},
	StoneGroup{2024, 1},
	StoneGroup{80, 1},
	StoneGroup{96, 1},
	StoneGroup{8, 1},
	StoneGroup{6, 2},
	StoneGroup{7, 1},
	StoneGroup{3, 1},
}

func Test_blink(t *testing.T) {
	tests := []struct {
		name        string
		stoneGroups StoneGroups
		output      StoneGroups
	}{
		{
			name:        "aoc sample 1",
			stoneGroups: aocSample1,
			output:      aocSample1Result,
		},
		{
			name:        "aoc sample 2",
			stoneGroups: aocSample2,
			output:      aocSample2Result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blink(tt.stoneGroups); !reflect.DeepEqual(sort(got), sort(tt.output)) {
				t.Errorf("blink() = %v, want %v", got, tt.output)
			}
		})
	}
}

func sort(sgs StoneGroups) StoneGroups {
	_sort.Slice(sgs, func(i, j int) bool {
		return sgs[i].stone < sgs[j].stone
	})

	return sgs
}
