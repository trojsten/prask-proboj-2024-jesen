package main

import (
	"errors"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"math"
	"strconv"
	"strings"
)

func loadPosition(data []string) (Position, error) {
	if len(data) != 2 {
		return Position{}, errors.New("wrong number of arguments")
	}

	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return Position{}, fmt.Errorf("x is invalid: %w", err)
	}

	y, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return Position{}, fmt.Errorf("y is invalid: %w", err)
	}

	return Position{x, y}, nil
}

func (g *Game) whereToMove(p *Player, target Position) Position {
	movementVector := p.Position.VectorTo(target)
	if movementVector.Length() > PLAYER_MOVEMENT_RANGE {
		movementVector = movementVector.Normalize().Mul(PLAYER_MOVEMENT_RANGE)
	}

	target = p.Position.Add(movementVector)

	wall, obstaclePosition := g.closestWallInTheWay(p, target)
	if wall != nil {
		angle := movementVector.Angle(wall.Vector())
		forcefieldVector := movementVector.Normalize().Mul(-1).Mul(PLAYER_RADIUS / math.Sin(angle))
		target = obstaclePosition.Add(forcefieldVector)
	}

	return target
}

func (g *Game) processTurn(p *Player) bool {
	resp, strData := g.Runner.ReadPlayer(p.Name)
	if resp != client.Ok {
		return false
	}

	data := strings.Split(strData, " ")
	command := data[0]
	args := data[1:]

	switch command {
	case "MOVE":
		target, err := loadPosition(args)
		if err != nil {
			g.Runner.Log(fmt.Sprintf("rejecting MOVE from %v: %v", p.Name, err))
			return false
		}

		target = g.whereToMove(p, target)
		p.Position = target
		g.Runner.Log(fmt.Sprintf("moving player %v to %v", p.Name, target))
	case "":
	}

	return true
}
