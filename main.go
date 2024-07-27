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
type AlwaysDefect struct{}
type Joss struct{}
type Grudger struct{}
type Pavlov struct{}
type TesterStrat struct{}
type SoftMajority struct{}
type HardMajority struct{}

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

func (t *Joss) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	defectProb := rand.Float32()
	if defectProb < 0.1 {
		return Defect
	} else if len(opponentHistory) == 0 {
		return Cooperate

	} else {
		l := len(opponentHistory)
		return opponentHistory[l-1]
	}
}
func (t *Joss) Name() string {
	return "Joss"
}
func CheckValidMoves(opponentHistory []Move, myHistory []Move) bool {
	if math.Abs(float64(len(opponentHistory)-len(myHistory))) > 0 {
		return false
	} else {
		return true
	}

}

func (t *AlwaysDefect) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}

	return Defect
}

func (t *AlwaysDefect) Name() string {
	return "Always Defect"
}

func (t *Grudger) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	l1 := len(opponentHistory)
	l2 := len(myHistory)
	if l1 == 0 {
		return Cooperate
	}

	if (l1 > 0 && opponentHistory[l1-1] == Defect) || (l2 > 0 && myHistory[l2-1] == Defect) {
		return Defect
	} else {
		return Cooperate
	}
}

func (t *Grudger) Name() string {
	return "Grudger"
}

func (t *Pavlov) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	l1 := len(opponentHistory)
	l2 := len(myHistory)
	if l1 == 0 {
		return Cooperate
	} else {
		if (opponentHistory[l1-1] == Defect && myHistory[l2-1] == Defect) || (opponentHistory[l1-1] == Cooperate && myHistory[l2-1] == Cooperate) {
			return myHistory[l2-1]
		} else {
			return 1 - myHistory[l2-1]
		}

	}

}

func (t *Pavlov) Name() string {
	return "Pavlov"
}

func (t *TesterStrat) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	//l1 := len(opponentHistory)
	l2 := len(myHistory)
	if l2 == 0 || l2 == 1 {
		return Cooperate
	} else if l2 == 2 {
		return Defect
	} else {
		for i, move := range opponentHistory {
			if i > 2 && move == Defect {
				return Defect
			}
		}
		return Cooperate
	}

}

func (t *TesterStrat) Name() string {
	return "Testing Strategy"
}

func (t *SoftMajority) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	l := len(opponentHistory)
	c := 0
	if l == 0 {
		return Cooperate
	}
	for i := 0; i < l; i++ {
		if opponentHistory[i] == Cooperate {
			c++
		}
		if myHistory[i] == Defect {
			c--
		}
	}
	if c > 0 {
		return Cooperate
	} else {
		return Defect
	}
}

func (t *SoftMajority) Name() string {
	return "SoftMajority"
}

func (t *HardMajority) Play(opponentHistory []Move, myHistory []Move) Move {
	if !CheckValidMoves(opponentHistory, myHistory) {
		panic("Invalid moves")
	}
	l := len(opponentHistory)
	c := 0
	for i := 0; i < l; i++ {
		if opponentHistory[i] == Defect {
			c++
		}
		if myHistory[i] == Cooperate {
			c--
		}
	}
	if c > 0 {
		return Defect
	} else {
		return Cooperate
	}

}

func (t *HardMajority) Name() string {
	return "HardMajority"
}
func main() {
	fmt.Println("Hello World New")

}
