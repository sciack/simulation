package montecarlo

const (
	WIN      = 36
	FAIR     = 36
	EUROPEAN = 37
	AMERICAN = 38
)

type Randomizer func(int) int

type Roulette struct {
	pockets int
	rnd     Randomizer
}

func FairRoulette(rnd Randomizer) *Roulette {
	return &Roulette{
		pockets: FAIR,
		rnd:     rnd,
	}
}

func EuropeanRoulette(rnd Randomizer) *Roulette {
	return &Roulette{
		pockets: EUROPEAN,
		rnd:     rnd,
	}
}

func AmericanRoulette(rnd Randomizer) *Roulette {
	return &Roulette{
		pockets: AMERICAN,
		rnd:     rnd,
	}
}

func (r *Roulette) spin(bet Bet) int {
	res := r.rnd(r.pockets) + 1
	if res == bet.Value {
		return bet.Amount * WIN
	} else {
		return 0
	}
}
