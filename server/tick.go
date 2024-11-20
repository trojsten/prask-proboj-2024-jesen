package main

func (g *Game) Tick() {
	g.Map.Radius *= WorldBorderSpeed
	center := Position{}

	for _, player := range g.Map.Players {
		if player.Health <= 0 {
			continue
		}
		player.ReloadCooldown = max(player.ReloadCooldown-1, 0)

		if center.Distance(player.Position) > g.Map.Radius-PlayerRadius {
			player.Health -= WorldBorderDamage
		}
	}

	newItems := []*Item{}
	for _, item := range g.Map.Items {
		if item.Distance(Position{}) >= g.Map.Radius {
			continue
		}
		newItems = append(newItems, item)
	}
	g.Map.Items = newItems
}
