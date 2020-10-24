package rule

import "game-of-life/internal/cell"

type UnderpopulationRule struct {}

const underpopulationThreshold = 2

func (underpopulationRule UnderpopulationRule) CanBeApplied(oldCell cell.Cell, neighbours cell.Neighbours) bool {
	return !oldCell.IsEmpty() && neighbours.NeighboringSpeciesCount[oldCell.OrganismSpecies] < underpopulationThreshold
}

func (underpopulationRule UnderpopulationRule) Apply(oldCell cell.Cell, neighbours cell.Neighbours) cell.Cell {
	return cell.CreateEmptyCell()
}
