package main

import (
	"testing"
)

func Test_calculateScore(t *testing.T) {
	type args struct {
		player1 string
		player2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Player 1 wins with rock
		{
			name: "Player 1 rock, Player 2 scissors",
			args: args{
				player1: "A", // rock
				player2: "Z", // scissors
			},
			want: 3, // 0 for losing, 3 for scissors
		},

		// Player 2 wins with rock
		{
			name: "Player 1 scissors, Player 2 rock",
			args: args{
				player1: "C", // scissors
				player2: "X", // rock
			},
			want: 7, // 6 for winning, 1 for rock
		},

		// Players tie with rock
		{
			name: "Player 1 rock, Player 2 rock",
			args: args{
				player1: "A", // rock
				player2: "X", // rock
			},
			want: 4, // 3 for draw, 1 for rock
		},

		// Player 1 plays rock, Player 2 plays paper
		{
			name: "Player 1 rock, Player 2 paper",
			args: args{
				player1: "A", // rock
				player2: "Y", // paper
			},
			want: 8, // 6 for winning, 2 for paper
		},

		// Player 1 plays paper, Player 2 plays rock
		{
			name: "Player 1 paper, Player 2 rock",
			args: args{
				player1: "B", // paper
				player2: "X", // rock
			},
			want: 1, // 0 for losing, 1 for rock
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScore(tt.args.player1, tt.args.player2); got != tt.want {
				t.Errorf("calculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMove(t *testing.T) {
	type args struct {
		player1        string
		desiredOutcome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rock, lose",
			args: args{
				player1:        "A",
				desiredOutcome: "X",
			},
			want: "Z", // scissors loses to rock
		},
		{
			name: "rock, draw",
			args: args{
				player1:        "A",
				desiredOutcome: "Y",
			},
			want: "X", // rock draws with rock
		},
		{
			name: "rock, win",
			args: args{
				player1:        "A",
				desiredOutcome: "Z",
			},
			want: "Y", // paper wins against rock
		},
		{
			name: "paper, lose",
			args: args{
				player1:        "B",
				desiredOutcome: "X",
			},
			want: "X", // rock loses to paper
		},
		{
			name: "paper, draw",
			args: args{
				player1:        "B",
				desiredOutcome: "Y",
			},
			want: "Y", // scissors draws with scissors
		},
		{
			name: "paper, win",
			args: args{
				player1:        "B",
				desiredOutcome: "Z",
			},
			want: "Z", // scissors wins against paper
		},
		{
			name: "scissors, lose",
			args: args{
				player1:        "C",
				desiredOutcome: "X",
			},
			want: "Y", // paper loses to scissors
		},
		{
			name: "scissors, draw",
			args: args{
				player1:        "C",
				desiredOutcome: "Y",
			},
			want: "Z", // scissors draws with scissors
		},
		{
			name: "scissors, win",
			args: args{
				player1:        "C",
				desiredOutcome: "Z",
			},
			want: "X", // rock wins against scissors
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMove(tt.args.player1, tt.args.desiredOutcome); got != tt.want {
				t.Errorf("calculateMove()\tgot %v, want %v", got, tt.want)
			}
		})
	}
}
