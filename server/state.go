package main

import (
	"encoding/json"
	"github.com/trojsten/ksp-proboj/client"
	"sort"
)

type state struct {
	Radius         float64       `json:"radius"`
	Player         Player        `json:"player"`
	VisibleItems   []Item        `json:"visible_items"`
	VisiblePlayers []statePlayer `json:"visible_players"`
}

type statePlayer struct {
	Position
	Id     int    `json:"id"`
	Weapon Weapon `json:"weapon"`
}

func (g *Game) closestWallInTheWay(p *Player, pos Position) (*Wall, Position) {
	var sortedWalls []*Wall
	for _, wall := range g.Map.Walls {
		sortedWalls = append(sortedWalls, wall)
	}
	sort.Slice(sortedWalls, func(i, j int) bool {
		return sortedWalls[i].Center().SquaredDistance(p.Position) < sortedWalls[j].Center().SquaredDistance(p.Position)
	})

	for _, wall := range sortedWalls {
		if collisionAt, collides := Intesect(p.Position, pos, wall.A, wall.B); collides {
			return wall, collisionAt
		}
	}
	return nil, Position{}
}

func (g *Game) stateForPlayer(p *Player) string {
	playerState := state{Player: *p, VisibleItems: []Item{}, VisiblePlayers: []statePlayer{}, Radius: g.Map.Radius}

	for _, item := range g.Map.Items {
		if wall, _ := g.closestWallInTheWay(p, item.Position); wall == nil {
			playerState.VisibleItems = append(playerState.VisibleItems, *item)
		}
	}
	for _, otherPlayer := range g.Map.Players {
		if otherPlayer.Id == p.Id {
			continue
		}

		if otherPlayer.Health <= 0 {
			continue
		}

		if wall, _ := g.closestWallInTheWay(p, otherPlayer.Position); wall == nil {
			playerState.VisiblePlayers = append(playerState.VisiblePlayers, statePlayer{
				Position: otherPlayer.Position,
				Id:       otherPlayer.Id,
				Weapon:   otherPlayer.Weapon,
			})
		}
	}

	stateString, err := json.Marshal(playerState)
	if err != nil {
		panic(err)
	}
	return string(stateString)
}

func (g *Game) ShouldContinue() bool {
	alive := 0
	for _, player := range g.Map.Players {
		if player.Health > 0 {
			alive++
		}
	}
	return alive > 1
}

func (g *Game) SendScores() {
	scores := client.Scores{}
	for _, player := range g.Map.Players {
		scores[player.Name] = player.Score
	}
	g.Runner.Scores(scores)
}
