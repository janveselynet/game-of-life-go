package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"game-of-life/internal/cell"
	"game-of-life/internal/rule"
)

var emptyCell = cell.CreateEmptyCell()
var fullCell = cell.Cell{OrganismSpecies: cell.OrganismSpecies(1)}

var cellsToEvolveProvider = []struct {
	title string
	cellsToEvolve CellsMap
	expectedCellsAfterEvolution CellsMap
}{
	{
		title: "all cells die",
		cellsToEvolve: CellsMap{
			0: CellsMapRow{0: fullCell, 1: emptyCell, 2: fullCell},
			1: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
			2: CellsMapRow{0: fullCell, 1: emptyCell, 2: fullCell},
		},
		expectedCellsAfterEvolution: CellsMap{
			0: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
			1: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
		},
	},
	{
		title: "two full rows",
		cellsToEvolve: CellsMap{
			0: CellsMapRow{0: fullCell, 1: fullCell, 2: fullCell},
			1: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
			2: CellsMapRow{0: fullCell, 1: fullCell, 2: fullCell},
		},
		expectedCellsAfterEvolution: CellsMap{
			0: CellsMapRow{0: emptyCell, 1: fullCell, 2: emptyCell},
			1: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: fullCell, 2: emptyCell},
		},
	},
	{
		title: "new cell is evolved",
		cellsToEvolve: CellsMap{
			0: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			1: CellsMapRow{0: fullCell, 1: emptyCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
		},
		expectedCellsAfterEvolution: CellsMap{
			0: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			1: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
		},
	},
	{
		title: "nothing happen",
		cellsToEvolve: CellsMap{
			0: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			1: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
		},
		expectedCellsAfterEvolution: CellsMap{
			0: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			1: CellsMapRow{0: fullCell, 1: fullCell, 2: emptyCell},
			2: CellsMapRow{0: emptyCell, 1: emptyCell, 2: emptyCell},
		},
	},
}

const worldSize = 3

func TestWorldIsEvolved(t *testing.T) {
	for _, testData := range cellsToEvolveProvider {
		t.Run(testData.title, func(t *testing.T) {
			world := World{size: worldSize, cells: testData.cellsToEvolve}
			rules := rule.RuleCollection{
				Rules: []rule.EvolutionRule{
					rule.UnderpopulationRule{},
					rule.OverpopulationRule{},
					rule.ReproductionRule{},
				},
			}

			worldAfterEvolution := world.Evolve(rules)

			assert.Equal(t, testData.expectedCellsAfterEvolution, worldAfterEvolution.cells)
			assert.Equal(t, worldSize, worldAfterEvolution.size)
		})
	}
}
