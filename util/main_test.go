package util

import (
	"reflect"
	"testing"
)

func TestGroupByN(t *testing.T) {
	type args struct {
		input     []string
		groupSize int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "group size 3",
			args: args{
				input:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				groupSize: 3,
			},
			want: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
		},
		{
			name: "group size 2",
			args: args{
				input:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				groupSize: 2,
			},
			want: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"5", "6"},
				{"7", "8"},
				{"9"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByN(tt.args.input, tt.args.groupSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByN() = %v, want %v", got, tt.want)
			}
		})
	}
}
