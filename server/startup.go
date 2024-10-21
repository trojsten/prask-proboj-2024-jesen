package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
)

type startupMap struct {
	Radius float64 `json:"radius"`
	Walls  []*Wall `json:"walls"`
}

func (g *Game) GreetPlayers() {
	startupState := startupMap{
		Radius: g.Map.Radius,
		Walls:  g.Map.Walls,
	}

	greeting, err := json.Marshal(startupState)
	if err != nil {
		panic(err)
	}

	for _, player := range g.Map.Players {
		r := g.Runner.ToPlayer(player.Name, "greeting", string(greeting))
		if r != client.Ok {
			panic(fmt.Sprintf("error while greeting players, got %v", r))
		}
	}
}
