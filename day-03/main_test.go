package main

import (
	"reflect"
	"testing"
)

func Test_getHalfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "even length",
			args: args{
				s: "abcdef",
			},
			want: []string{"abc", "def"},
		},
		{
			name: "odd length",
			args: args{
				s: "abcde",
			},
			want: []string{"ab", "cde"},
		},
		{
			name: "aoc1",
			args: args{
				s: "vJrwpWtwJgWrhcsFMMfFFhFp",
			},
			want: []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		},
		{
			name: "aoc2",
			args: args{
				s: "PmmdzqPrVvPwwTWBwg",
			},
			want: []string{"PmmdzqPrV", "vPwwTWBwg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getHalfString(tt.args.s)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHalfString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToRuneSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: []rune{},
		},
		{
			name: "one",
			args: args{
				s: "a",
			},
			want: []rune{'a'},
		},
		{
			name: "two",
			args: args{
				s: "ab",
			},
			want: []rune{'a', 'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToRuneSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToRuneSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCommonCharacter(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "aoc1",
			args: args{
				s1: "vJrwpWtwJgWr",
				s2: "hcsFMMfFFhFp",
			},
			want: "p",
		},
		{
			name: "aoc2",
			args: args{
				s1: "jqHRNqRjqzjGDLGL",
				s2: "rsFMfFZSrLrFZsSL",
			},
			want: "L",
		},
		{
			name: "aoc3",
			args: args{
				s1: "PmmdzqPrV",
				s2: "vPwwTWBwg",
			},
			want: "P",
		},
		{
			name: "aoc4",
			args: args{
				s1: "wMqvLMZHhHMvwL",
				s2: "HjbvcjnnSBnvTQFn",
			},
			want: "v",
		},
		{
			name: "aoc5",
			args: args{
				s1: "ttgJtRGJQ",
				s2: "ctTZtZT",
			},
			want: "t",
		},
		{
			name: "aoc6",
			args: args{
				s1: "CrZsJsPPZsGz",
				s2: "wwsLwLmpwMDw",
			},
			want: "s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonCharacter(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("getCommonCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumPriorities(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				r: []rune{},
			},
			want: 0,
		},
		{
			name: "a, 1",
			args: args{
				r: []rune{'a'},
			},
			want: 1,
		},
		{
			name: "ab, 1+2",
			args: args{
				r: []rune{'a', 'b'},
			},
			want: 3,
		},
		{
			name: "A, 27",
			args: args{
				r: []rune{'A'},
			},
			want: 27,
		},
		{
			name: "aoc",
			args: args{
				r: []rune{'p', 'L', 'P', 'v', 't', 's'},
			},
			want: 157,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPrioities(tt.args.r); got != tt.want {
				t.Errorf("sumPriorities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupByThree(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "6",
			args: args{
				input: []string{"a", "b", "c", "d", "e", "f"},
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupByThree(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupByThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCommonCharacters(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "aoc1",
			args: args{
				s1: "vJrwpWtwJgWrhcsFMMfFFhFp",
				s2: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				s3: "PmmdzqPrVvPwwTWBwg",
			},
			want: "r",
		},

		{
			name: "aoc2",
			args: args{
				s1: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				s2: "ttgJtRGJQctTZtZT",
				s3: "CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonCharacters(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("getCommonCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
