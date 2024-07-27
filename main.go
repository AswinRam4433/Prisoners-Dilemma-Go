package main

import (
	"fmt"
	"math"
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

func (t *TitForTat) Play(opponentHistory []Move, myHistory []Move) Move {
	if len(opponentHistory) == 0 {
		return Cooperate
	}
	return opponentHistory[len(opponentHistory)-1]
}

func (t *TitForTat) Name() string {
	return "TitForTat"
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
