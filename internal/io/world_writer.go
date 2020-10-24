package io

import "game-of-life/internal/environment"

type WorldWriter interface {
	WriteWorld(world environment.World) error
}
