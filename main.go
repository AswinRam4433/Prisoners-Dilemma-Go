package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Move int

const (
	Cooperate Move = iota
	Defect
)

type Strategy interface {
	Name() string
	Play(oppHistory []Move, myHistory []Move) Move
}
type TitForTat struct{}
type GenerousTitForTat struct{}
type Random struct{}
type AlwaysCooperate struct{}

func (t *TitForTat) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}

	if len(opponentHistory) == 0 {
		return Cooperate
	}
	return opponentHistory[len(opponentHistory)-1]
}

func (t *TitForTat) Name() string {
	return "TitForTat"
}

func (t *GenerousTitForTat) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	if len(opponentHistory) == 0 || len(opponentHistory) == 1 {
		return Cooperate
	} else {
		l := len(opponentHistory)
		if opponentHistory[l-1] == Defect && opponentHistory[l-2] == Defect {
			return Defect
		} else {
			return Cooperate
		}
	}
}

func (t *GenerousTitForTat) Name() string {
	return "Generous Tit For Tat"
}

func (t *Random) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	if rand.Float32() < 0.5 {
		return Defect
	} else {
		return Cooperate
	}
}

func (t *Random) Name() string {
	return "Random"
}

func (t *AlwaysCooperate) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}

	return Cooperate
}
func (t *AlwaysCooperate) Name() string {
	return "Always Cooperate"
}

func CheckValidMoves(opponentHistory []Move, myHistory []Move) bool {
	if math.Abs(float64(len(opponentHistory)-len(myHistory))) > 1 {
		return false
	} else {
		return true
	}

}

func main() {
	fmt.Println("Hello World New")

}
