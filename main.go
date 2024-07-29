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
	Coop() int
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

var greenHexCode = "#00FF00"
var redHexCode = "#FF0000"
var yellowHexCode = "#FFFF00"

func CheckValidMoves(opponentHistory []Move, myHistory []Move) bool {
	// Checks if the game state is valid
	if math.Abs(float64(len(opponentHistory)-len(myHistory))) > 0 {
		return false
	} else {
		return true
	}

}

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

func (t *TitForTat) Coop() int {
	return 1
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

func (t *GenerousTitForTat) Coop() int {
	return 1
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

func (t *Random) Coop() int {
	return 0
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

func (t *AlwaysCooperate) Coop() int {
	return 1
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

func (t *Joss) Coop() int {
	return 1
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

func (t *AlwaysDefect) Coop() int {
	return -1
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

func (t *Grudger) Coop() int {
	return -1
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

func (t *Pavlov) Coop() int {
	return 1
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

func (t *TesterStrat) Coop() int {
	return 0
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

func (t *SoftMajority) Coop() int {
	return 1
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

func (t *HardMajority) Coop() int {
	return 0
}

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
	// list of strategies to run on this simulation

	rounds := 200
	results := make(map[string]int)

	matchChan := make(chan struct {
		p1, p2         Strategy
		score1, score2 int
	})
	// to collect head to head performance. Help from ChatGPT

	roundWiseResults := make(chan struct {
		p1, p2                   Strategy
		scoresList1, scoresList2 []Move
	}, len(SubmittedStrategies)*len(SubmittedStrategies)) // Buffer size for channel. Code idea by ChatGPT

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

	// Wait for all matches to complete
	time.Sleep(6 * time.Second)
	// Ensure this is closed after all matches are done or else causes errors
	close(matchChan)
	close(roundWiseResults)

	// Print overall results
	fmt.Println("\n\nOverall Results:")
	for strategy, score := range results {
		fmt.Printf("%s: %d\n", strategy, score)
	}

	// to group strategies
	strategyClassification := make(map[string]int)
	for i := range SubmittedStrategies {
		strategyClassification[SubmittedStrategies[i].Name()] = SubmittedStrategies[i].Coop()
	}
	// to color code strategies as per cooperating, defecting and neutral
	colorResults := colorMapper(strategyClassification)
	// Visualize the results of the simulation
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

func colorMapper(strats map[string]int) map[string]string {
	// assuming cooperating strategies have green, neutral have yellow and defecting have red

	colorResults := make(map[string]string)
	for k, v := range strats {
		if v == 1 {
			colorResults[k] = greenHexCode

		} else if v == -1 {
			colorResults[k] = redHexCode
		} else {
			colorResults[k] = yellowHexCode
		}
	}
	// return the color map
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
	dataLabels := make([]string, 0)

	// Define the color categories
	greenSeries := make([]opts.BarData, 0)
	redSeries := make([]opts.BarData, 0)
	yellowSeries := make([]opts.BarData, 0)

	for strategy, score := range results {
		color, exists := colorResults[strategy]
		if !exists {
			color = "#000000"
			// Default to black if no color is defined
		}

		dataLabels = append(dataLabels, strategy)

		// Assign data points to the appropriate series based on color
		barData := opts.BarData{
			Value: score,
			ItemStyle: &opts.ItemStyle{
				Color: color,
			},
		}
		switch color {
		case "#00FF00": // Green
			greenSeries = append(greenSeries, barData)
			redSeries = append(redSeries, opts.BarData{}) // Add empty data to maintain alignment. ChatGPT idea
			yellowSeries = append(yellowSeries, opts.BarData{})
		case "#FF0000": // Red
			redSeries = append(redSeries, barData)
			greenSeries = append(greenSeries, opts.BarData{}) // Add empty data to maintain alignment. ChatGPT idea
			yellowSeries = append(yellowSeries, opts.BarData{})
		case "#FFFF00": // Yellow
			yellowSeries = append(yellowSeries, barData)
			greenSeries = append(greenSeries, opts.BarData{}) // Add empty data to maintain alignment. ChatGPT idea
			redSeries = append(redSeries, opts.BarData{})
		default:
			greenSeries = append(greenSeries, opts.BarData{})
			redSeries = append(redSeries, opts.BarData{})
			yellowSeries = append(yellowSeries, opts.BarData{})
		}
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Prisoner's Dilemma",
			Subtitle: "Strategy Wise Performance",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Rotate: 45,
			},
		}),
	)
	// Rotate the x-axis labels by 45 degrees to ensure they are completely visible

	// Add series to the bar chart
	bar.SetXAxis(dataLabels).
		AddSeries("Cooperating Strategies", greenSeries, charts.WithItemStyleOpts(opts.ItemStyle{Color: greenHexCode})).
		AddSeries("Defecting Strategies", redSeries, charts.WithItemStyleOpts(opts.ItemStyle{Color: redHexCode})).
		AddSeries("Neutral Strategies", yellowSeries, charts.WithItemStyleOpts(opts.ItemStyle{Color: yellowHexCode}))
	// WithItemStyleOpts is important for the legend colors to match. ChatGPT idea

	f, err := os.Create("myVis.html")
	if err != nil {
		panic(err)
	}
	bar.Render(f)

	fmt.Println("Dummy Prints")
	// to satisfy go compiler
	fmt.Println(matchChan)
	fmt.Println(roundWiseResults)
}
