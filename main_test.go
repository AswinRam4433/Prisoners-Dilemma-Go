package main

import "testing"

func TestTitForTat_Play(t1 *testing.T) {
	type args struct {
		opponentHistory []Move
		myHistory       []Move
	}
	tests := []struct {
		name string
		args args
		want Move
	}{
		{"Start Simulation", args{[]Move{}, []Move{}}, Cooperate},
		//{""},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TitForTat{}
			if got := t.Play(tt.args.opponentHistory, tt.args.myHistory); got != tt.want {
				t1.Errorf("Play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckValidMoves(t *testing.T) {
	type args struct {
		opponentHistory []Move
		myHistory       []Move
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases
		{"Same Size", args{[]Move{Defect}, []Move{Cooperate}}, true},
		{"Overflow In One", args{[]Move{Defect, Defect, Defect}, []Move{Cooperate}}, false},
		{"Overflow In Two", args{[]Move{Cooperate}, []Move{Defect, Defect, Defect}}, false},
		{"Empty Both", args{[]Move{}, []Move{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidMoves(tt.args.opponentHistory, tt.args.myHistory); got != tt.want {
				t.Errorf("CheckValidMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerousTitForTat_Play(t1 *testing.T) {
	type args struct {
		opponentHistory []Move
		myHistory       []Move
	}
	tests := []struct {
		name string
		args args
		want Move
	}{
		// TODO: Add test cases.
		{"Start", args{[]Move{}, []Move{}}, Cooperate},
		{"One Move Made-1", args{[]Move{Defect}, []Move{}}, Cooperate},
		{"One Move Made-2", args{[]Move{Cooperate}, []Move{}}, Cooperate},
		{"One Move Made-3", args{[]Move{Defect}, []Move{Cooperate}}, Cooperate},
		{"One Move Made-4", args{[]Move{Cooperate}, []Move{Cooperate}}, Cooperate},
		{"Two Moves Made", args{[]Move{Cooperate, Cooperate}, []Move{Cooperate}}, Cooperate},
		{"Two Moves Made", args{[]Move{Cooperate, Defect}, []Move{Cooperate}}, Cooperate},
		{"Two Moves Made", args{[]Move{Defect, Defect}, []Move{Cooperate}}, Defect},
		{"Two Moves Made", args{[]Move{Defect, Cooperate}, []Move{Cooperate}}, Cooperate},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &GenerousTitForTat{}
			if got := t.Play(tt.args.opponentHistory, tt.args.myHistory); got != tt.want {
				t1.Errorf("Play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom_Play(t1 *testing.T) {
	type args struct {
		opponentHistory []Move
		myHistory       []Move
	}
	tests := []struct {
		name string
		args args
		want Move
	}{
		{"Sample", args{[]Move{}, []Move{}}, Cooperate},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Random{}
			if got := t.Play(tt.args.opponentHistory, tt.args.myHistory); got != tt.want {
				t1.Errorf("Play() = %v, want %v", got, tt.want)
			}
		})
	}
}
