# Prisoner's Dilemma Tournament

This project implements a simulation of the Prisoner's Dilemma game with multiple strategies. The goal is to compare and visualize various strategies based on their performance over a series of matches.
Inspired by [Veritasium Video](https://www.youtube.com/watch?v=mScpHTIi-kM).

## Overview

In this simulation, different strategies compete against each other in a series of rounds. The strategies used are:

- **TitForTat**: Cooperates initially, then mimics the opponent's previous move.
- **GenerousTitForTat**: Similar to TitForTat, but with a tendency to cooperate even after the opponent defects twice in a row.
- **Random**: Chooses between Cooperate and Defect randomly with equal probability.
- **AlwaysCooperate**: Always chooses to cooperate.
- **AlwaysDefect**: Always chooses to defect.
- **Joss**: Defects with a probability of 10%, otherwise mimics the opponent's previous move.
- **Grudger**: Starts by cooperating, but defects if the opponent defects at least once.
- **Pavlov**: Adapts its strategy based on the outcome of the previous round.
- **TesterStrat**: Cooperates initially, defects after the second round, and defects if the opponent has defected more than twice.
- **SoftMajority**: Chooses to cooperate if the majority of opponent's moves are cooperate, otherwise defects.
- **HardMajority**: Chooses to defect if the majority of opponent's moves are defect, otherwise cooperates.

## Installation

This project uses Go 1.21.5

Clone this repository and navigate to the project directory:

```sh
git clone https://github.com/AswinRam4433/Prisoners-Dilemma-Go.git
cd prisoners-dilemma
```

## Usage
You can run the simulation by executing the main function in main.go. This will:

Initialize a set of strategies.
Run a series of matches between these strategies.
Collect and print the results.
```sh
go run main.go
```

## Code Description
* Strategy Interface: Defines the methods that all strategies must implement:

  + Name(): Returns the name of the strategy.
  + Play(): Determines the move to play based on the opponent's history and the strategy's own history.

* Move Type: An enum representing the possible moves in the game: Cooperate and Defect

* CheckValidMoves Function: Checks if the histories of moves are valid (i.e., both histories are of the same length).

* playMatch Function: Simulates a match between two strategies for a given number of rounds. It collects and sends results to channels for further processing.

* Channels
  + matchChan: A channel for sending match results (total scores of strategies).
  + roundWiseResults: A channel for sending detailed round-wise results.
  + Goroutines: The simulation uses goroutines to handle concurrent matches between strategies and to process results concurrently.

## Results
After running the simulation, results for each strategy are printed, showing the total scores. Round-wise results for each match are also printed, showing the history of moves and scores for each round.

## Contributing
Feel free to contribute by creating issues, submitting pull requests, or suggesting improvements.


