package main

import (
	"fmt"
	"game-of-life/internal/io"
	"game-of-life/internal/rule"
)

const inputFile = "./examples/input1.xml"
const outputFile = "./out.xml"

var rules = rule.RuleCollection{
	Rules: []rule.EvolutionRule{
		rule.UnderpopulationRule{},
		rule.OverpopulationRule{},
		rule.ReproductionRule{},
	},
}

func main() {
	input := io.XmlGameReader{FileName: inputFile}
	game, err := input.ReadGame()
	if err != nil {
		fmt.Println(err)
		return
	}

	finalWorld := game.Run(rules)

	output := io.XmlWorldWriter{FileName: outputFile}
	err = output.WriteWorld(finalWorld)
	if err != nil {
		fmt.Println(err)
	}
}
