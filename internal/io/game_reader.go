package io

import "game-of-life/internal/game"

type GameReader interface {
	ReadGame() (game.Game, error)
}
