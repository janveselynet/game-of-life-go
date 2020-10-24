package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"game-of-life/internal/cell"
)

var canReproductionRuleBeAppliedDataProvider = []struct {
	title string
	cellOrganismSpecies cell.OrganismSpecies
	expectedCanBeApplied bool
}{
	{
		title: "cell is empty",
		cellOrganismSpecies: cell.EmptyOrganismSpecies,
		expectedCanBeApplied: true,
	},
	{
		title: "cell is not empty",
		cellOrganismSpecies: cell.OrganismSpecies(1),
		expectedCanBeApplied: false,
	},
}

func TestReproductionRuleCanBeAppliedCorrectly(t *testing.T) {
	for _, testData := range canReproductionRuleBeAppliedDataProvider {
		t.Run(testData.title, func(t *testing.T) {
			cellToEvolve := cell.Cell{OrganismSpecies: testData.cellOrganismSpecies}
			neighbours := cell.Neighbours{}

			actualCanBeApplied := ReproductionRule{}.CanBeApplied(cellToEvolve, neighbours)

			assert.Equal(t, testData.expectedCanBeApplied, actualCanBeApplied)
		})
	}
}

var neighboursToApplyReproductionRuleDataProvider = []struct {
	title string
	neighboringSpeciesCount cell.NeighboringSpeciesCountMap
	expectedCellOrganismSpecies []cell.OrganismSpecies
}{
	{
		title: "no species to reproduce #1",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{},
		expectedCellOrganismSpecies: []cell.OrganismSpecies{cell.EmptyOrganismSpecies},
	},
	{
		title: "no species to reproduce #2",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{
			cell.OrganismSpecies(1): 2,
			cell.OrganismSpecies(2): 1,
			cell.OrganismSpecies(3): 4,
		},
		expectedCellOrganismSpecies: []cell.OrganismSpecies{cell.EmptyOrganismSpecies},
	},
	{
		title: "one species to reproduce",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{
			cell.OrganismSpecies(1): 2,
			cell.OrganismSpecies(2): 3,
			cell.OrganismSpecies(3): 4,
		},
		expectedCellOrganismSpecies: []cell.OrganismSpecies{cell.OrganismSpecies(2)},
	},
	{
		title: "multiple species to reproduce",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{
			cell.OrganismSpecies(1): 3,
			cell.OrganismSpecies(2): 3,
			cell.OrganismSpecies(3): 3,
		},
		expectedCellOrganismSpecies: []cell.OrganismSpecies{
			cell.OrganismSpecies(1),
			cell.OrganismSpecies(2),
			cell.OrganismSpecies(3),
		},
	},
}

func TestReproductionRuleIsApplied(t *testing.T) {
	for _, testData := range neighboursToApplyReproductionRuleDataProvider {
		t.Run(testData.title, func(t *testing.T) {
			cellToEvolve := cell.Cell{}
			neighbours := cell.Neighbours{NeighboringSpeciesCount: testData.neighboringSpeciesCount}

			cellAfterApplication := ReproductionRule{}.Apply(cellToEvolve, neighbours)

			assert.Subset(
				t,
				testData.expectedCellOrganismSpecies,
				[]cell.OrganismSpecies{cellAfterApplication.OrganismSpecies},
			)
		})
	}
}
