package rule

import "game-of-life/internal/cell"

type OverpopulationRule struct {}

const overpopulationThreshold = 3

func (overpopulationRule OverpopulationRule) CanBeApplied(oldCell cell.Cell, neighbours cell.Neighbours) bool {
	return !oldCell.IsEmpty() && neighbours.NeighboringSpeciesCount[oldCell.OrganismSpecies] > overpopulationThreshold
}

func (overpopulationRule OverpopulationRule) Apply(oldCell cell.Cell, neighbours cell.Neighbours) cell.Cell {
	return cell.CreateEmptyCell()
}
