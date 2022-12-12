package main

import (
	"testing"
)

func Test_parseRange(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "1-5",
			args:  args{a: "1-5"},
			want:  1,
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseRange(tt.args.a)
			if got != tt.want {
				t.Errorf("parseRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isSuperSet(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-4, 6-8",
			args: args{
				a: "2-4",
				b: "6-8",
			},
			want: false,
		},
		{
			name: "2-3, 4-5",
			args: args{
				a: "2-3",
				b: "4-5",
			},
			want: false,
		},
		{
			name: "5-7, 7-9",
			args: args{
				a: "5-7",
				b: "7-9",
			},
			want: false,
		},
		{
			name: "2-8, 3-7",
			args: args{
				a: "2-8",
				b: "3-7",
			},
			want: true,
		},
		{
			name: "6-6, 4-6",
			args: args{
				a: "6-6",
				b: "4-6",
			},
			want: true,
		},
		{
			name: "2-6, 4-8",
			args: args{
				a: "2-6",
				b: "4-8",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSuperSet(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isSuperSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC example data",
			args: args{
				input: []string{
					"2-4",
					"6-8",
					"2-3",
					"4-5",
					"5-7",
					"7-9",
					"2-8",
					"3-7",
					"6-6",
					"4-6",
					"2-6",
					"4-8",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
