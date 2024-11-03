package main

func (g *Game) Tick() {
	g.Map.Radius *= WorldBorderSpeed
	center := Position{}

	for _, player := range g.Map.Players {
		if player.Health <= 0 {
			continue
		}
		player.ReloadCooldown = max(player.ReloadCooldown-1, 0)

		if center.Distance(player.Position) > g.Map.Radius {
			player.Health -= WorldBorderDamage
		}
	}
}
