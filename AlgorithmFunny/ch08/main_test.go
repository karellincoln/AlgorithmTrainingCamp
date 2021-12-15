package main

import "testing"

func Test_checkValidGraph(t *testing.T) {
	type args struct {
		graph *[GROUPS_ITEMS][GROUPS_COUNT]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "validTest", args: args{
			&[GROUPS_ITEMS][GROUPS_COUNT]int{
				{4,1,2,3,5},
				{1,2,4,5,3},
				{2,1,5,3,4},
				{2,1,3,4,5},
				{2,1,4,3,5},
			},
		},
		want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidGraph(tt.args.graph); got != tt.want {
				t.Errorf("checkValidGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
