package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"math"
	"math/rand"
)

type startupMap struct {
	Radius float64 `json:"radius"`
	Walls  []*Wall `json:"walls"`
}

func (g *Game) spawnPoint() Position {
	spawnAngle := rand.Float64() * 2 * math.Pi
	spawnDist := rand.Float64() * g.Map.Radius

	return Position{
		X: math.Cos(spawnAngle) * spawnDist,
		Y: math.Sin(spawnAngle) * spawnDist,
	}
}

func (g *Game) SpawnPlayers(players []string) {
	for i, player := range players {
		spawn := g.spawnPoint()

		g.Map.Players = append(g.Map.Players, &Player{
			Position:       spawn,
			Id:             i,
			Name:           player,
			Health:         PlayerFullHealth,
			Weapon:         WeaponNone,
			LoadedAmmo:     0,
			ReloadCooldown: 0,
		})
	}
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
