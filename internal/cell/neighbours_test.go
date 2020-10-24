package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleNeighbours = Neighbours{
	NeighboringSpeciesCount: NeighboringSpeciesCountMap{
		OrganismSpecies(1): 3,
		OrganismSpecies(2): 2,
		OrganismSpecies(3): 3,
	},
}

var gettingSpeciesWithCertainCountTestData = []struct {
	title string
	countToFilterBy int
	expectedSpecies []OrganismSpecies
}{
	{
		title: "species with count 3",
		countToFilterBy: 3,
		expectedSpecies: []OrganismSpecies{1, 3},
	},
	{
		title: "species with count 2",
		countToFilterBy: 2,
		expectedSpecies: []OrganismSpecies{2},
	},
	{
		title: "species with count 1",
		countToFilterBy: 1,
		expectedSpecies: []OrganismSpecies{},
	},
}

func TestGettingSpeciesWithCertainCount(t *testing.T) {
	for _, testData := range gettingSpeciesWithCertainCountTestData {
		t.Run(testData.title, func(t *testing.T) {
			actualSpecies := sampleNeighbours.GetSpeciesWithCount(testData.countToFilterBy)

			assert.Equal(t, testData.expectedSpecies, actualSpecies)
		})
	}
}

func TestCreatingNeighboursFromCells(t *testing.T) {
	cells := []Cell{
		{OrganismSpecies: 1},
		{OrganismSpecies: 3},
		{OrganismSpecies: 1},
		{OrganismSpecies: 2},
		{OrganismSpecies: 2},
		{OrganismSpecies: 2},
	}

	neighbours := CreateNeighboursFromCells(cells)

	assert.Equal(t, 2, neighbours.NeighboringSpeciesCount[1])
	assert.Equal(t, 3, neighbours.NeighboringSpeciesCount[2])
	assert.Equal(t, 1, neighbours.NeighboringSpeciesCount[3])
}
