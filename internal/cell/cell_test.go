package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const isEmpty = true
const isNotEmpty = false

var isCellEmptyTestData = []struct {
	title string
	cellOrganismSpecies OrganismSpecies
	expectedIsEmpty bool
}{
	{
		title: "is empty",
		cellOrganismSpecies: EmptyOrganismSpecies,
		expectedIsEmpty: isEmpty,
	},
	{
		title: "is not empty #1",
		cellOrganismSpecies: OrganismSpecies(1),
		expectedIsEmpty: isNotEmpty,
	},
	{
		title: "is not empty #2",
		cellOrganismSpecies: OrganismSpecies(153),
		expectedIsEmpty: isNotEmpty,
	},
}

func TestIsCellEmpty(t *testing.T) {
	for _, testData := range isCellEmptyTestData {
		t.Run(testData.title, func(t *testing.T) {
			cell := Cell{testData.cellOrganismSpecies}

			actualIsEmpty := cell.IsEmpty()

			assert.Equal(t, testData.expectedIsEmpty, actualIsEmpty)
		})
	}
}

func TestCreatingEmptyCell(t *testing.T) {
	emptyCell := CreateEmptyCell()

	assert.Equal(t, emptyCell.OrganismSpecies, EmptyOrganismSpecies)
}
