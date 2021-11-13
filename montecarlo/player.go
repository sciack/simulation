package montecarlo

type Player struct {
	game Roulette
	rnd  Randomizer
}

type Roulette func(Bet, Randomizer) int

type Bet struct {
	Amount int
	Value  int
}

func CreatePlayer(roulette Roulette, rnd Randomizer) *Player {
	return &Player{
		game: roulette,
		rnd:  rnd,
	}
}

func (p *Player) Play(bet Bet, round int) int {
	outcome := 0
	for i := 0; i < round; i++ {
		outcome += p.game(bet, p.rnd) - bet.Amount
	}
	return outcome
}
