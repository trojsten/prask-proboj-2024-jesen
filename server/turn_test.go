package main

import (
	"testing"
)

func TestGame_whereToMove(t *testing.T) {
	tests := []struct {
		name   string
		walls  []*Wall
		start  Position
		target Position
		want   Position
	}{
		{
			name:   "in range no walls",
			walls:  nil,
			target: Position{1, 1},
			want:   Position{1, 1},
		},
		{
			name:   "too far",
			walls:  nil,
			target: Position{PlayerMovementRange * 100, 0},
			want:   Position{PlayerMovementRange, 0},
		},
		{
			name: "wall away",
			walls: []*Wall{
				{Position{100, -100}, Position{100, 100}},
			},
			target: Position{10, 0},
			want:   Position{10, 0},
		},
		{
			name: "wall behind",
			walls: []*Wall{
				{Position{10, -20}, Position{10, 20}},
			},
			target: Position{20, 0},
			want:   Position{10 - PlayerRadius, 0},
		},
		{
			name: "right at wall",
			walls: []*Wall{
				{Position{PlayerRadius, -20}, Position{PlayerRadius, 20}},
			},
			target: Position{20, 0},
			want:   Position{0, 0},
		},
		{
			name: "inside wall",
			walls: []*Wall{
				{Position{PlayerRadius - 2, -20}, Position{PlayerRadius - 2, 20}},
			},
			target: Position{20, 0},
			want:   Position{-2, 0},
		},
		{
			name: "wall end",
			walls: []*Wall{
				{Position{10, 0}, Position{10, 20}},
			},
			target: Position{20, 0},
			want:   Position{10 - PlayerRadius, 0},
		},
		{
			name: "bug",
			walls: []*Wall{
				{Position{0, 0}, Position{250, 250}},
			},
			start:  Position{-10, 8},
			target: Position{46, 68},
			want:   Position{3.6463650072002203, 22.621105364857378},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Map: &Map{
					Radius:  10000,
					Walls:   tt.walls,
					Items:   []*Item{},
					Players: []*Player{},
				},
			}

			if got := g.whereToMove(&Player{Position: tt.start}, tt.target); got.Distance(tt.want) >= 0.0001 {
				t.Errorf("whereToMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
