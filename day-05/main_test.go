package main

import (
	"reflect"
	"testing"

	"github.com/bbriggs/aoc-2022/util"
)

func Test_popN(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "pop 1",
			args: args{
				s: []string{"A", "B", "C"},
				n: 1,
			},
			want:  []string{"C"},
			want1: []string{"A", "B"},
		},
		{
			name: "pop 2",
			args: args{
				s: []string{"A", "B", "C"},
				n: 2,
			},
			want:  []string{"B", "C"},
			want1: []string{"A"},
		},
		{
			name: "pop 3",
			args: args{
				s: []string{"A", "B", "C"},
				n: 3,
			},
			want:  []string{"A", "B", "C"},
			want1: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := popN(tt.args.s, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("popN() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("popN() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestState_MoveChunk(t *testing.T) {
	type args struct {
		m Move
	}
	tests := []struct {
		name string
		s    State
		args args
		want State
	}{
		{
			name: "move 1 element",
			s: State{
				{"Z", "N"},
				{"M", "C", "D"},
				{"P"},
			},
			args: args{
				m: Move{
					from:  1, // stack 2 (index 1)
					to:    0, // stack 1 (index 0)
					count: 1,
				},
			},
			want: State{
				{"Z", "N", "D"},
				{"M", "C"},
				{"P"},
			},
		},

		{
			name: "move 3 elements",
			// Start with previous state
			s: State{
				{"Z", "N", "D"},
				{"M", "C"},
				{"P"},
			},
			args: args{
				m: Move{
					from:  0, // stack 1 (index 0)
					to:    2, // stack 3 (index 2)
					count: 3,
				},
			},
			want: State{
				{},
				{"M", "C"},
				{"P", "Z", "N", "D"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MoveChunk(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("State.MoveChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pushN(t *testing.T) {
	type args struct {
		s []string
		e []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "push 1",
			args: args{
				s: []string{"A", "B"},
				e: []string{"C"},
			},
			want: []string{"A", "B", "C"},
		},
		{
			name: "push 2",
			args: args{
				s: []string{"A", "B"},
				e: []string{"C", "D"},
			},
			want: []string{"A", "B", "C", "D"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushN(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pushN() = %v, want %v", got, tt.want)
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
		want string
	}{
		{
			// this test ensures that I don't monkey around with known good functions
			// if it fails, I've broken something that caused the original known good solution to fail
			name: "integirity check",
			args: args{
				input: util.GetInput(),
			},
			want: "VGBBJCRMN",
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

func Test_part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// test against my hand-calculated solution for the first 5 steps
			name: "integirity check",
			args: args{
				input: util.GetInput()[0:6], // only the first 6 steps to get through some challenging indexing issues
			},
			want: "SRJFZCMJL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
