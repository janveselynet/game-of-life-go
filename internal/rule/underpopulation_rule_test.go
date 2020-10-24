package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"game-of-life/internal/cell"
)

var canUnderpopulationRuleBeAppliedDataProvider = []struct {
	title string
	cellOrganismSpecies cell.OrganismSpecies
	neighboringSpeciesCount cell.NeighboringSpeciesCountMap
	expectedCanBeApplied bool
}{
	{
		title: "cell is empty",
		cellOrganismSpecies: cell.EmptyOrganismSpecies,
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{},
		expectedCanBeApplied: false,
	},
	{
		title: "cell has no neighbour with same type",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{},
		expectedCanBeApplied: true,
	},
	{
		title: "cell has one neighbour with same type",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{cell.OrganismSpecies(1): 1},
		expectedCanBeApplied: true,
	},
	{
		title: "cell has two neighbour with same type",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{cell.OrganismSpecies(1): 2},
		expectedCanBeApplied: false,
	},
	{
		title: "cell has three neighbour with same type",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{cell.OrganismSpecies(1): 3},
		expectedCanBeApplied: false,
	},
	{
		title: "cell has many neighbour with same type",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{cell.OrganismSpecies(1): 8},
		expectedCanBeApplied: false,
	},
}

func TestUnderpopulationRuleCanBeAppliedCorrectly(t *testing.T) {
	for _, testData := range canUnderpopulationRuleBeAppliedDataProvider {
		t.Run(testData.title, func(t *testing.T) {
			cellToEvolve := cell.Cell{OrganismSpecies: testData.cellOrganismSpecies}
			neighbours := cell.Neighbours{NeighboringSpeciesCount: testData.neighboringSpeciesCount}

			actualCanBeApplied := UnderpopulationRule{}.CanBeApplied(cellToEvolve, neighbours)

			assert.Equal(t, testData.expectedCanBeApplied, actualCanBeApplied)
		})
	}
}

func TestUnderpopulationRuleIsApplied(t *testing.T) {
	cellToEvolve := cell.Cell{OrganismSpecies: 10}
	neighbours := cell.Neighbours{}

	result := UnderpopulationRule{}.Apply(cellToEvolve, neighbours)

	assert.Equal(t, cell.EmptyOrganismSpecies, result.OrganismSpecies)
}
