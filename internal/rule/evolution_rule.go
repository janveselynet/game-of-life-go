package rule

import "game-of-life/internal/cell"

type EvolutionRule interface {
	CanBeApplied(cell cell.Cell, neighbours cell.Neighbours) bool
	Apply(cell cell.Cell, neighbours cell.Neighbours) cell.Cell
}
