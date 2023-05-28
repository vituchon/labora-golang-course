package fighters

import (
	"math/rand"
)

type Paladin struct {
	BaseFighter
}

const PaladinInitialLife = 20

func NewPaladin() Paladin {
	return Paladin{
		BaseFighter: BaseFighter{
			Life: PaladinInitialLife,
		},
	}
}

func (p *Paladin) ThrowAttack() int {
	attack := rand.Intn(10)
	proportion := p.Life / PaladinInitialLife
	return attack * proportion
}

func (p *Paladin) RecieveAttack(intensity int) {
	p.Life -= intensity
}

func (p *Paladin) GetName() string {
	return "Paladin"
}
