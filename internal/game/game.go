package game

import (
	"game-of-life/internal/environment"
	"game-of-life/internal/rule"
)

type Game struct {
	InitialWorld environment.World
	IterationsCount int
}

func (game Game) Run(rules rule.RuleCollection) environment.World {
	world := game.InitialWorld

	iterationNumber := 1
	for ; iterationNumber <= game.IterationsCount; iterationNumber++ {
		world = world.Evolve(rules)
	}

	return world
}
