package cell

type Cell struct {
	OrganismSpecies OrganismSpecies
}

func (cell Cell) IsEmpty() bool {
	return cell.OrganismSpecies == EmptyOrganismSpecies
}

func CreateEmptyCell() Cell {
	return Cell{OrganismSpecies: EmptyOrganismSpecies}
}
