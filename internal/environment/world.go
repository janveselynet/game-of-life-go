package environment

import (
	"math"

	"game-of-life/internal/cell"
	"game-of-life/internal/rule"
)

type CellsMapRow map[int]cell.Cell
type CellsMap map[int]CellsMapRow

type World struct {
	Size int
	Cells CellsMap
}

func (world World) Evolve(rules rule.RuleCollection) World {
	evolvedCells := CellsMap{}

	for y := 0; y < world.Size; y++ {
		evolvedCells[y] = CellsMapRow{}
		for x := 0; x < world.Size; x++ {
			neighbours := world.getNeighboringCells(x, y)
			evolvedCell := rules.Apply(world.Cells[y][x], neighbours)
			evolvedCells[y][x] = evolvedCell
		}
	}

	return World{
		Size: world.Size,
		Cells: evolvedCells,
	}
}

func (world World) getNeighboringCells(x, y int) cell.Neighbours {
	minX := int(math.Max(float64(x - 1), 0))
	maxX := int(math.Min(float64(x + 1), float64(world.Size - 1)))
	minY := int(math.Max(float64(y - 1), 0))
	maxY := int(math.Min(float64(y + 1), float64(world.Size - 1)))

	neighboringCells := []cell.Cell{}
	for neighbourY := minY; neighbourY <= maxY; neighbourY++ {
		for neighbourX := minX; neighbourX <= maxX; neighbourX++ {
			if !(neighbourX == x && neighbourY == y) && !world.Cells[neighbourY][neighbourX].IsEmpty() {
				neighboringCells = append(neighboringCells, world.Cells[neighbourY][neighbourX])
			}
		}
	}

	return cell.CreateNeighboursFromCells(neighboringCells)
}
