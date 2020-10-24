package rule

import "game-of-life/internal/cell"

type RuleCollection struct {
	Rules []EvolutionRule
}

func (ruleCollection RuleCollection) Apply(cellToApplyTo cell.Cell, neighbours cell.Neighbours) cell.Cell {
	for _ , rule := range ruleCollection.Rules {
		if rule.CanBeApplied(cellToApplyTo, neighbours) {
			return rule.Apply(cellToApplyTo, neighbours)
		}
	}

	return cellToApplyTo
}
