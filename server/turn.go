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
	if movementVector.Length() > PlayerMovementRange {
		movementVector = movementVector.Normalize().Mul(PlayerMovementRange)
	}

	target = p.Position.Add(movementVector)

	wall, obstaclePosition := g.closestWallInTheWay(p, target)
	if wall != nil {
		angle := movementVector.Angle(wall.Vector())
		forcefieldVector := movementVector.Normalize().Mul(-1).Mul(PlayerRadius / math.Sin(angle))
		target = obstaclePosition.Add(forcefieldVector)
	}

	return target
}

func (g *Game) shoot(shooter *Player, target *Player) error {
	if shooter.ReloadCooldown > 0 {
		return fmt.Errorf("still reloading: %v", shooter.ReloadCooldown)
	}

	weapon := WeaponStatsMap[shooter.Weapon]
	dist := shooter.Position.SquaredDistance(target.Position)
	if dist > weapon.Range {
		return fmt.Errorf("target is out of range")
	}

	wall, pos := g.closestWallInTheWay(shooter, target.Position)
	if wall != nil {
		return fmt.Errorf("wall in the way @ %v", pos)
	}

	target.Health -= weapon.Damage
	shooter.LoadedAmmo--
	if shooter.LoadedAmmo <= 0 {
		shooter.ReloadCooldown = weapon.ReloadTime
		shooter.LoadedAmmo = weapon.AmmoCapacity
	}
	return nil
}

func (g *Game) closestItem(pos Position) *Item {
	var closest *Item = nil
	for _, item := range g.Map.Items {
		if item.Distance(pos) > PlayerPickupRadius {
			continue
		}

		if closest == nil || item.Distance(pos) < closest.Distance(pos) {
			closest = item
		}
	}
	return closest
}

func (g *Game) removeItem(item *Item) {
	newItems := []*Item{}

	for _, i := range g.Map.Items {
		if i == item {
			continue
		}
		newItems = append(newItems, i)
	}

	g.Map.Items = newItems
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
	case "SHOOT":
		targetPlayerId, err := strconv.Atoi(args[0])
		if err != nil {
			g.Runner.Log(fmt.Sprintf("rejecting SHOOT from %v: %v", p.Name, err))
			return false
		}

		if targetPlayerId < 0 || targetPlayerId >= len(g.Map.Players) {
			g.Runner.Log(fmt.Sprintf("rejecting SHOOT from %v: unknown player %v", p.Name, targetPlayerId))
			return false
		}
		targetPlayer := g.Map.Players[targetPlayerId]

		err = g.shoot(p, targetPlayer)
		if err != nil {
			g.Runner.Log(fmt.Sprintf("rejecting SHOOT from %v: %v", p.Name, err))
			return false
		}
	case "RELOAD":
		if p.ReloadCooldown != 0 {
			g.Runner.Log(fmt.Sprintf("rejecting RELOAD from %v: still reloading, %v", p.Name, p.ReloadCooldown))
			return false
		}

		weapon := WeaponStatsMap[p.Weapon]
		p.ReloadCooldown = weapon.ReloadTime
		p.LoadedAmmo = weapon.AmmoCapacity
	case "NOOP":
	case "DROP":
		if p.Weapon == WeaponNone {
			g.Runner.Log(fmt.Sprintf("rejecting DROP from %v: no weapon to drop", p.Name))
			return false
		}
		g.Map.Items = append(g.Map.Items, &Item{
			Position: p.Position,
			Type:     ItemWeapon,
			Weapon:   p.Weapon,
		})
		p.Weapon = WeaponNone
		p.LoadedAmmo = 0
	case "PICKUP":
		item := g.closestItem(p.Position)
		if item == nil {
			g.Runner.Log(fmt.Sprintf("rejecting PICKUP from %v: no item in range", p.Name))
			return false
		}

		switch item.Type {
		case ItemWeapon:
			weapon := item.Weapon
			if p.Weapon != WeaponNone {
				item.Weapon = p.Weapon
			} else {
				g.removeItem(item)
			}
			p.Weapon = weapon
			p.LoadedAmmo = WeaponStatsMap[p.Weapon].AmmoCapacity
			p.ReloadCooldown = WeaponStatsMap[p.Weapon].ReloadTime
		case ItemHealth:
			p.Health = min(p.Health+HealthItemRegeneration, PlayerFullHealth)
			g.removeItem(item)
		}
	}

	return true
}
