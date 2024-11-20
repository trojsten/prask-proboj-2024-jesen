package main

import "encoding/json"

type observerMap struct {
	Radius float64 `json:"radius"`
	Walls  []*Wall `json:"walls"`
}

type observerState struct {
	Radius  float64   `json:"radius"`
	Items   []*Item   `json:"items"`
	Players []*Player `json:"players"`
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
		Radius:  g.Map.Radius,
		Items:   g.Map.Items,
		Players: g.Map.Players,
	}

	data, err := json.Marshal(oState)
	if err != nil {
		panic(err)
	}

	g.Runner.ToObserver(string(data) + "\n")
}
