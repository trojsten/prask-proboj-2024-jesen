package main

import (
	"github.com/trojsten/ksp-proboj/client"
	"reflect"
	"testing"
)

func TestGame_closestWallInTheWay(t *testing.T) {
	type fields struct {
		Runner *client.Runner
		Map    *Map
		Turn   int
	}
	type args struct {
		p   *Player
		pos Position
	}
	tests := []struct {
		name         string
		walls        []*Wall
		playerPos    Position
		pos          Position
		wallId       int
		intersection Position
	}{
		{"no walls", []*Wall{}, Position{0, 0}, Position{100, 0}, -1, Position{}},
		{"single wall", []*Wall{{Position{50, -10}, Position{50, 10}}}, Position{0, 0}, Position{100, 0}, 0, Position{50, 0}},
		{"two walls", []*Wall{{Position{100, -10}, Position{100, 10}}, {Position{50, -10}, Position{50, 10}}}, Position{0, 0}, Position{200, 0}, 1, Position{50, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Map: &Map{
					Walls: tt.walls,
				},
			}
			got, got1 := g.closestWallInTheWay(&Player{Position: tt.playerPos}, tt.pos)
			if tt.wallId == -1 {
				if got != nil {
					t.Errorf("closestWallInTheWay() returned wall %v, want nil", *got)
				}
			} else {
				if got == nil {
					t.Errorf("closestWallInTheWay() returned nil, want %v", *tt.walls[tt.wallId])
				} else if !reflect.DeepEqual(got, tt.walls[tt.wallId]) {
					t.Errorf("closestWallInTheWay() returned wall %v, want %v", *got, *tt.walls[tt.wallId])
				}
			}
			if !reflect.DeepEqual(got1, tt.intersection) {
				t.Errorf("closestWallInTheWay() got1 = %v, want %v", got1, tt.intersection)
			}
		})
	}
}
