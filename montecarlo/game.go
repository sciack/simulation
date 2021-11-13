package montecarlo

const (
	WIN      = 36
	FAIR     = 36
	EUROPEAN = 37
	AMERICAN = 38
)

type Randomizer func(int) int

func FairRoulette(bet Bet, rnd Randomizer) int {
	return spin(bet, FAIR, rnd)
}

func EuropeanRoulette(bet Bet, rnd Randomizer) int {
	return spin(bet, EUROPEAN, rnd)
}

func AmericanRoulette(bet Bet, rnd Randomizer) int {
	return spin(bet, AMERICAN, rnd)
}

func spin(bet Bet, roulette int, rnd Randomizer) int {
	res := rnd(roulette) + 1
	if res == bet.Value {
		return bet.Amount * WIN
	} else {
		return 0
	}
}
