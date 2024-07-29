package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"math/rand"
	"os"
	"time"
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

//func main() {
////	SubmittedStrategies := []Strategy{
////		&TitForTat{},
////		&GenerousTitForTat{},
////		&Random{},
////		&AlwaysCooperate{},
////		&AlwaysDefect{},
////		&Joss{},
////		&Grudger{},
////		&Pavlov{},
////		&TesterStrat{},
////		&SoftMajority{},
////		&HardMajority{},
////	}
////
////	rounds := 100
////	results := make(map[string]int)
////
////	matchChan := make(chan struct {
////		p1, p2         Strategy
////		score1, score2 int
////	})
////
////	roundWiseResults := make(chan struct {
////		p1, p2                   Strategy
////		scoresList1, scoresList2 []Move
////	})
////
////	// GENAI START
////	// Needed help to run the go routines
////
////	var wg sync.WaitGroup
////
////	go func() {
////		for result := range matchChan {
////			results[result.p1.Name()] += result.score1
////			results[result.p2.Name()] += result.score2
////		}
////	}()
////
////	for i := 0; i < len(SubmittedStrategies); i++ {
////		for j := i; j < len(SubmittedStrategies); j++ {
////			wg.Add(1)
////			go func(s1, s2 Strategy) {
////				defer wg.Done()
////				playMatch(s1, s2, rounds, matchChan, roundWiseResults)
////			}(SubmittedStrategies[i], SubmittedStrategies[j])
////		}
////	}
////
////	wg.Wait()
////	close(matchChan)
////
////	// GENAI END
////
////	//for i := 0; i < len(SubmittedStrategies); i++ {
////	//	for j := i; j < len(SubmittedStrategies); j++ {
////	//		go playMatch(SubmittedStrategies[i], SubmittedStrategies[j], rounds, matchChan)
////	//	}
////	//}
////
////	fmt.Println("Printing Results")
////	for strategy, score := range results {
////		fmt.Printf("%s: %d\n", strategy, score)
////	}
////	fmt.Println("End Of Results")
////
////}

//func main() {
//	SubmittedStrategies := []Strategy{
//		&TitForTat{},
//		&GenerousTitForTat{},
//		&Random{},
//		&AlwaysCooperate{},
//		&AlwaysDefect{},
//		&Joss{},
//		&Grudger{},
//		&Pavlov{},
//		&TesterStrat{},
//		&SoftMajority{},
//		&HardMajority{},
//	}
//
//	rounds := 200
//	results := make(map[string]int)
//	matchChan := make(chan struct {
//		p1, p2         Strategy
//		score1, score2 int
//	})
//	roundWiseResults := make(chan struct {
//		p1, p2                   Strategy
//		scoresList1, scoresList2 []Move
//	})
//
//	go func() {
//		for result := range matchChan {
//			results[result.p1.Name()] += result.score1
//			results[result.p2.Name()] += result.score2
//		}
//	}()
//
//	for i := 0; i < len(SubmittedStrategies); i++ {
//		for j := i; j < len(SubmittedStrategies); j++ {
//			go playMatch(SubmittedStrategies[i], SubmittedStrategies[j], rounds, matchChan, roundWiseResults)
//		}
//	}
//
//	time.Sleep(2 * time.Second) // Wait for goroutines to finish
//	close(matchChan)
//	close(roundWiseResults)
//
//	for strategy, score := range results {
//		fmt.Printf("%s: %d\n", strategy, score)
//	}
//
//	// Process and print round-wise results
//	for result := range roundWiseResults {
//		fmt.Printf("Match: %s VS %s\n", result.p1.Name(), result.p2.Name())
//		fmt.Printf("Scores: %v VS %v\n", result.scoresList1, result.scoresList2)
//	}
//}
//
//func playMatch(strategy1 Strategy, strategy2 Strategy, rounds int, matchChan chan<- struct {
//	p1, p2         Strategy
//	score1, score2 int
//}, roundWiseResults chan<- struct {
//	p1, p2                   Strategy
//	scoresList1, scoresList2 []Move
//}) {
//	var strat1History, strat2History []Move
//	var strat1Score, strat2Score int
//
//	for i := 0; i < rounds; i++ {
//		//fmt.Printf("In round %d of match %s VS %s\n", i, strategy1.Name(), strategy2.Name())
//
//		move1 := strategy1.Play(strat2History, strat1History)
//		move2 := strategy2.Play(strat1History, strat2History)
//		// function signature and calls have to match semantically
//
//		strat1History = append(strat1History, move1)
//		strat2History = append(strat2History, move2)
//
//		if move1 == move2 {
//			if move1 == Defect {
//				// both parties Defected
//				// both get 1 point each
//				strat1Score = strat1Score + 1
//				strat2Score = strat2Score + 1
//			} else {
//				// both parties Cooperated
//				// both get 3 points each
//				strat1Score = strat1Score + 3
//				strat2Score = strat2Score + 3
//
//			}
//		} else {
//			// whoever Defects receives 5 points
//			// whoever Cooperated receives 0 points
//			if move1 == Defect {
//				strat1Score = strat1Score + 5
//			} else {
//				strat2Score = strat2Score + 5
//			}
//		}
//
//	}
//	roundWiseResults <- struct {
//		p1, p2                   Strategy
//		scoresList1, scoresList2 []Move
//	}{strategy1, strategy2, strat1History, strat2History}
//
//	fmt.Printf("In %s VS %s: %d VS %d\n", strategy1.Name(), strategy2.Name(), strat1Score, strat2Score)
//	matchChan <- struct {
//		p1, p2         Strategy
//		score1, score2 int
//	}{strategy1, strategy2, strat1Score, strat2Score}
//
//}

func main() {
	SubmittedStrategies := []Strategy{
		&TitForTat{},
		&GenerousTitForTat{},
		&Random{},
		&AlwaysCooperate{},
		&AlwaysDefect{},
		&Joss{},
		&Grudger{},
		&Pavlov{},
		&TesterStrat{},
		&SoftMajority{},
		&HardMajority{},
	}

	rounds := 200
	results := make(map[string]int)
	matchChan := make(chan struct {
		p1, p2         Strategy
		score1, score2 int
	})
	roundWiseResults := make(chan struct {
		p1, p2                   Strategy
		scoresList1, scoresList2 []Move
	}, len(SubmittedStrategies)*len(SubmittedStrategies)) // Buffer size for channel

	go func() {
		for result := range matchChan {
			results[result.p1.Name()] += result.score1
			results[result.p2.Name()] += result.score2
		}
	}()

	go func() {
		for result := range roundWiseResults {
			fmt.Printf("\nRound Wise Results For Match: %s VS %s\n", result.p1.Name(), result.p2.Name())
			fmt.Printf("%v VS \n%v", result.scoresList1, result.scoresList2)
		}
	}()

	for i := 0; i < len(SubmittedStrategies); i++ {
		for j := i; j < len(SubmittedStrategies); j++ {
			go playMatch(SubmittedStrategies[i], SubmittedStrategies[j], rounds, matchChan, roundWiseResults)
		}
	}

	// Wait for all matches to complete (adjust sleep time as needed)
	time.Sleep(10 * time.Second)
	close(matchChan)
	close(roundWiseResults) // Ensure this is closed after all matches are done

	// Print overall results
	fmt.Println("\n\nOverall Results:")
	for strategy, score := range results {
		fmt.Printf("%s: %d\n", strategy, score)
	}
	strategyClassification := make(map[string]bool)
	for i := range SubmittedStrategies {
		strategyClassification[SubmittedStrategies[i].Name()] = true
	}
	colorResults := colorMapper(strategyClassification)
	visResults(matchChan, roundWiseResults, results, colorResults)
}

func playMatch(strategy1, strategy2 Strategy, rounds int, matchChan chan<- struct {
	p1, p2         Strategy
	score1, score2 int
}, roundWiseResults chan<- struct {
	p1, p2                   Strategy
	scoresList1, scoresList2 []Move
}) {
	var strat1History, strat2History []Move
	var strat1Score, strat2Score int
	var scoresList1, scoresList2 []int

	for i := 0; i < rounds; i++ {
		move1 := strategy1.Play(strat2History, strat1History)
		move2 := strategy2.Play(strat1History, strat2History)

		strat1History = append(strat1History, move1)
		strat2History = append(strat2History, move2)

		if move1 == move2 {
			if move1 == Defect {
				strat1Score++
				strat2Score++
			} else {
				strat1Score += 3
				strat2Score += 3
			}
		} else {
			if move1 == Defect {
				strat1Score += 5
			} else {
				strat2Score += 5
			}
		}

		// Collecting scores for each round
		scoresList1 = append(scoresList1, strat1Score)
		scoresList2 = append(scoresList2, strat2Score)
	}

	roundWiseResults <- struct {
		p1, p2                   Strategy
		scoresList1, scoresList2 []Move
	}{strategy1, strategy2, strat1History, strat2History}

	matchChan <- struct {
		p1, p2         Strategy
		score1, score2 int
	}{strategy1, strategy2, strat1Score, strat2Score}
}

func colorMapper(strats map[string]bool) map[string]string {
	// assuming cooperating strategies have True
	greenHexCode := "#00FF00"
	redHexCode := "#FF0000"
	colorResults := make(map[string]string)
	for k, v := range strats {
		if v == true {
			colorResults[k] = greenHexCode

		} else {
			colorResults[k] = redHexCode
		}
	}
	return colorResults

}
func visResults(
	matchChan chan<- struct {
		p1, p2         Strategy
		score1, score2 int
	},
	roundWiseResults chan<- struct {
		p1, p2                   Strategy
		scoresList1, scoresList2 []Move
	},
	results map[string]int,
	colorResults map[string]string,
) {
	fmt.Println("Visualizing Results")
	dataPoints := make([]opts.BarData, 0)
	dataLabels := make([]string, 0)

	//for strategy, score := range results {
	//	dataPoints = append(dataPoints, opts.BarData{Value: int32(score)})
	//	dataLabels = append(dataLabels, strategy)
	//}

	for strategy, score := range results {
		color, exists := colorResults[strategy]
		if !exists {
			color = "#000000" // Default to black if no color is defined
		}

		dataPoints = append(dataPoints, opts.BarData{
			Value: score,
			ItemStyle: &opts.ItemStyle{
				Color: color,
			},
		})
		dataLabels = append(dataLabels, strategy)
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Prisoner's Dilemma",
		Subtitle: "Strategy Wise Performance",
	}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Rotate: 45,
			},
		}))

	bar.SetXAxis(dataLabels).AddSeries("Strategies", dataPoints)

	f, err := os.Create("myVis.html")
	if err != nil {
		panic(err)
	}
	bar.Render(f)

	fmt.Println("Dummy Prints")
	fmt.Println(matchChan)
	fmt.Println(roundWiseResults)
}
