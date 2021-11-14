package montecarlo_test

import (
	"testing"

	"github.com/sciack/simulation/montecarlo"
)

func TestWinning(t *testing.T) {
	rnd := func(_ int) int { return 0 }
	p := montecarlo.CreatePlayer(montecarlo.FairRoulette(rnd))
	result := p.Play(montecarlo.Bet{Amount: 1, Value: 1}, 1)
	if result != 35 {
		t.Errorf("Expecting 35 got %d", result)
	}
}

func TestWinning5Times(t *testing.T) {
	rnd := func(_ int) int { return 0 }
	p := montecarlo.CreatePlayer(montecarlo.FairRoulette(rnd))
	result := p.Play(montecarlo.Bet{Amount: 1, Value: 1}, 5)
	if result != 175 {
		t.Errorf("Expecting 175 got %d", result)
	}
}

func TestLoosing(t *testing.T) {
	rnd := func(_ int) int { return 10 }
	p := montecarlo.CreatePlayer(montecarlo.FairRoulette(rnd))
	result := p.Play(montecarlo.Bet{Amount: 1, Value: 1}, 5)
	if result != -5 {
		t.Errorf("Expecting -5 got %d", result)
	}
}
