package rule

import (
  "math/rand"
	"time"
	
	"game-of-life/internal/cell"
)

type ReproductionRule struct {}

const reproductionThreshold = 3

func (reproductionRule ReproductionRule) CanBeApplied(oldCell cell.Cell, neighbours cell.Neighbours) bool {
	return oldCell.IsEmpty()
}

func (reproductionRule ReproductionRule) Apply(oldCell cell.Cell, neighbours cell.Neighbours) cell.Cell {
	speciesThatCanReproduce := neighbours.GetSpeciesWithCount(reproductionThreshold)

	if (len(speciesThatCanReproduce) == 0) {
		return cell.CreateEmptyCell()
	}

	rand.Seed(time.Now().Unix())
	speciesToReproduce := speciesThatCanReproduce[rand.Intn(len(speciesThatCanReproduce))]

	return cell.Cell{OrganismSpecies: speciesToReproduce}
}
