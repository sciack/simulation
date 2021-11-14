package montecarlo

type Player struct {
	game *Roulette
}

type Bet struct {
	Amount int
	Value  int
}

func CreatePlayer(roulette *Roulette) *Player {
	return &Player{
		game: roulette,
	}
}

func (p *Player) Play(bet Bet, round int) int {
	outcome := 0
	for i := 0; i < round; i++ {
		outcome += p.game.spin(bet) - bet.Amount
	}
	return outcome
}
