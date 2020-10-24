package cell

type NeighboringSpeciesCountMap map[OrganismSpecies]uint

type Neighbours struct {
	NeighboringSpeciesCount NeighboringSpeciesCountMap
}

func (neighbours Neighbours) GetSpeciesWithCount(countToFilterBy uint) []OrganismSpecies {
	speciesToReturn := []OrganismSpecies{}

	for organismSpecies, speciesCount := range neighbours.NeighboringSpeciesCount {
		if speciesCount == countToFilterBy {
			speciesToReturn = append(speciesToReturn, organismSpecies)
		}
	}

	return speciesToReturn
}

func createNeighboursFromCells(cells []Cell) Neighbours {
	neighboringSpeciesCount := NeighboringSpeciesCountMap{}

	for _ , neighbour := range cells {
		currentCount := neighboringSpeciesCount[neighbour.OrganismSpecies]
		neighboringSpeciesCount[neighbour.OrganismSpecies] = currentCount + 1
	}

	return Neighbours{NeighboringSpeciesCount: neighboringSpeciesCount}
}
