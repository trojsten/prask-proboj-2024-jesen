package main

import "encoding/json"

type observerMap struct {
	Radius float64 `json:"radius"`
	Walls  []*Wall `json:"walls"`
}

type observerState struct {
	Radius    float64            `json:"radius"`
	Items     []*Item            `json:"items"`
	Players   []*Player          `json:"players"`
	Shootings []observerShooting `json:"shootings"`
}

type observerShooting struct {
	Attacker Position `json:"attacker"`
	Target   Position `json:"target"`
}

func (g *Game) SendMapToObserver() {
	oMap := observerMap{
		Radius: g.Map.Radius,
		Walls:  g.Map.Walls,
	}

	data, err := json.Marshal(oMap)
	if err != nil {
		panic(err)
	}

	g.Runner.ToObserver(string(data) + "\n")
}

func (g *Game) SendStateToObserver() {
	oState := observerState{
		Radius:    g.Map.Radius,
		Items:     g.Map.Items,
		Players:   g.Map.Players,
		Shootings: []observerShooting{},
	}

	for _, shooting := range g.TurnShootings {
		oState.Shootings = append(oState.Shootings, observerShooting{Target: shooting.Target.Position, Attacker: shooting.Attacker.Position})
	}

	data, err := json.Marshal(oState)
	if err != nil {
		panic(err)
	}

	g.Runner.ToObserver(string(data) + "\n")
}
