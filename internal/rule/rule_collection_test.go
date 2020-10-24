package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"game-of-life/internal/cell"
)

var neighboursToApplyEvolutionOnRuleCollectionDataProvider = []struct {
	title string
	neighboringSpeciesCount cell.NeighboringSpeciesCountMap
	expectedCellOrganismSpecies cell.OrganismSpecies
}{
	{
		title: "rule is not applied",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{
			cell.OrganismSpecies(2): 2,
		},
		expectedCellOrganismSpecies: cell.OrganismSpecies(2),
	},
	{
		title: "rule is applied",
		neighboringSpeciesCount: cell.NeighboringSpeciesCountMap{
			cell.OrganismSpecies(2): 1,
		},
		expectedCellOrganismSpecies: cell.EmptyOrganismSpecies,
	},
}

var sampleRules = RuleCollection {
	Rules: []EvolutionRule{UnderpopulationRule{}},
}

func TestRuleCollectionIsApplied(t *testing.T) {
	for _, testData := range neighboursToApplyEvolutionOnRuleCollectionDataProvider {
		t.Run(testData.title, func(t *testing.T) {
			cellToEvolve := cell.Cell{OrganismSpecies: cell.OrganismSpecies(2)}
			neighbours := cell.Neighbours{NeighboringSpeciesCount: testData.neighboringSpeciesCount}

			cellAfterApplication := sampleRules.Apply(cellToEvolve, neighbours)

			assert.Equal(t, testData.expectedCellOrganismSpecies, cellAfterApplication.OrganismSpecies)
		})
	}
}
