package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sciack/simulation/analysis"
	"github.com/sciack/simulation/montecarlo"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomizer := func(limit int) int {
		return rand.Intn(limit)
	}
	player := montecarlo.CreatePlayer(montecarlo.FairRoulette, randomizer)
	fmt.Println("Fair roulette")
	simulation(player, 100)
	simulation(player, 10_000)
	simulation(player, 1_000_000)
	simulation(player, 100_000_000)

	player = montecarlo.CreatePlayer(montecarlo.EuropeanRoulette, randomizer)
	fmt.Println("European roulette")
	simulation(player, 100)
	simulation(player, 10_000)
	simulation(player, 1_000_000)
	simulation(player, 100_000_000)

	player = montecarlo.CreatePlayer(montecarlo.AmericanRoulette, randomizer)
	fmt.Println("American roulette")
	simulation(player, 100)
	simulation(player, 10_000)
	simulation(player, 1_000_000)
	simulation(player, 100_000_000)
}

func simulation(player *montecarlo.Player, spins int) {
	outcomes := []int{}
	for i := 0; i <= spins; i++ {
		bet := montecarlo.Bet{Amount: 1, Value: 1}
		outcomes = append(outcomes, player.Play(bet, 1))
	}
	result := analysis.ComputeStats(&outcomes)
	fmt.Printf("Spin: %d, Avg: %f, Variance: %f, StdDev: %f", spins, result.Average, result.Variance, result.StdDev)
	fmt.Println()

	fmt.Printf("Confidence interval 95%%: [%f, %f] \n", result.Confidence95.Min, result.Confidence95.Max)

}
