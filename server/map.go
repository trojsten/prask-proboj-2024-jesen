package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type WallJson [4]float64
type MapJson []WallJson

func (g *Game) LoadMap(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Errorf("could not open map file: %w", err))
	}

	mapData := MapJson{}
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		panic(fmt.Errorf("could not read map file: %w", err))
	}

	for _, wallData := range mapData {
		wall := Wall{
			A: Position{wallData[0], wallData[1]},
			B: Position{wallData[2], wallData[3]},
		}
		g.Map.Walls = append(g.Map.Walls, &wall)
	}
}
