package main

import (
	"testing"
)

func Test_getStartofMessage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aoc1",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "aoc2",
			args: args{
				input: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			name: "aoc3",
			args: args{
				input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			name: "aoc4",
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStartofMessage(tt.args.input); got != tt.want {
				t.Errorf("getStartofMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasDuplicates(t *testing.T) {
	type args struct {
		input []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all 4 same",
			args: args{
				input: []rune{'a', 'a', 'a', 'a'},
			},
			want: true,
		},
		{
			name: "all 4 different",
			args: args{
				input: []rune{'a', 'b', 'c', 'd'},
			},
			want: false,
		},
		{
			name: "3 same",
			args: args{
				input: []rune{'a', 'a', 'a', 'b'},
			},
			want: true,
		},
		{
			name: "2 same",
			args: args{
				input: []rune{'a', 'a', 'b', 'c'},
			},
			want: true,
		},
		{
			name: "2 same, first and last",
			args: args{
				input: []rune{'a', 'b', 'c', 'a'},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDuplicates(tt.args.input); got != tt.want {
				t.Errorf("hasDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessages(t *testing.T) {
	type args struct {
		input string
		start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "aoc1",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
				start: 5,
			},
			want: "plbgvbhsrlpgdmjqwftvncz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessages(tt.args.input, tt.args.start); got != tt.want {
				t.Errorf("getMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
