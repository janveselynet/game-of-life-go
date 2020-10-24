package cell

type NeighboringSpeciesCountMap map[OrganismSpecies]int

type Neighbours struct {
	NeighboringSpeciesCount NeighboringSpeciesCountMap
}

func (neighbours Neighbours) GetSpeciesWithCount(countToFilterBy int) []OrganismSpecies {
	speciesToReturn := []OrganismSpecies{}

	for organismSpecies, speciesCount := range neighbours.NeighboringSpeciesCount {
		if speciesCount == countToFilterBy {
			speciesToReturn = append(speciesToReturn, organismSpecies)
		}
	}

	return speciesToReturn
}

func CreateNeighboursFromCells(cells []Cell) Neighbours {
	neighboringSpeciesCount := NeighboringSpeciesCountMap{}

	for _ , neighbour := range cells {
		currentCount := neighboringSpeciesCount[neighbour.OrganismSpecies]
		neighboringSpeciesCount[neighbour.OrganismSpecies] = currentCount + 1
	}

	return Neighbours{NeighboringSpeciesCount: neighboringSpeciesCount}
}
