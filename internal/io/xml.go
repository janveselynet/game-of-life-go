package io

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"game-of-life/internal/game"
	"game-of-life/internal/environment"
	"game-of-life/internal/cell"
)

type XmlLife struct {
	XMLName xml.Name `xml:"life"`
	World XmlWorld `xml:"world"`
	Organisms XmlOrganisms `xml:"organisms"`
}

type XmlWorld struct {
	XMLName xml.Name `xml:"world"`
	Size int `xml:"cells"`
	Iterations int `xml:"iterations"`
}

type XmlOrganisms struct {
	XMLName xml.Name `xml:"organisms"`
	Organisms []XmlOrganism `xml:"organism"`
}

type XmlOrganism struct {
	XMLName xml.Name `xml:"organism"`
	XPos int `xml:"x_pos"`
	YPos int `xml:"y_pos"`
	Species int `xml:"species"`
}

type XmlGameReader struct {
	FileName string
}

func (xmlGameReader XmlGameReader) ReadGame() (game.Game, error) {
	xmlFile, err := os.Open(xmlGameReader.FileName)
	if err != nil {
		return game.Game{}, err
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return game.Game{}, err
	}

	var life XmlLife
	err = xml.Unmarshal(byteValue, &life)
	if err != nil {
		return game.Game{}, err
	}

	worldSize := life.World.Size
	worldCells := environment.CellsMap{}
	for y := 0; y < worldSize; y++ {
		worldCells[y] = environment.CellsMapRow{}
		for x := 0; x < worldSize; x++ {
			worldCells[y][x] = cell.Cell{OrganismSpecies: cell.EmptyOrganismSpecies}
		}
	}
	for _ , organism := range life.Organisms.Organisms {
		worldCells[organism.YPos][organism.XPos] = cell.Cell{
			OrganismSpecies: cell.OrganismSpecies(organism.Species + 1),
		}
	}

	return game.Game{
		InitialWorld: environment.World{
			Size: worldSize,
			Cells: worldCells,
		},
		IterationsCount: life.World.Iterations,
	}, nil
}

type XmlWorldWriter struct {
	FileName string
}

func (xmlWorldWriter XmlWorldWriter) WriteWorld(world environment.World) error {
	organisms := []XmlOrganism{}
	for y , cellRow := range world.Cells {
		for x , cell := range cellRow {
			if !cell.IsEmpty() {
				organisms = append(organisms, XmlOrganism{
					XPos: x,
					YPos: y,
					Species: int(cell.OrganismSpecies) - 1,
				})
			}
		}
	}

	life := XmlLife{
		World: XmlWorld{ Size: world.Size },
		Organisms: XmlOrganisms{ Organisms: organisms },
	}

	file, err := xml.MarshalIndent(life, "", " ")
	if err != nil {
		return err
	}
 
	err = ioutil.WriteFile(xmlWorldWriter.FileName, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
