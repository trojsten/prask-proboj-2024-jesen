package main

type Weapon int

const (
	WEAPON_NONE = iota
)

type Player struct {
	Position
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Health         int    `json:"health"`
	Weapon         Weapon `json:"weapon"`
	LoadedAmmo     int    `json:"loaded_ammo"`
	ReloadCooldown int    `json:"reload_cooldown"`
}

type Stats struct{}

type ItemType int

const (
	ITEM_WEAPON = iota
	ITEM_HEALTH
)

type Item struct {
	Position
	Type   ItemType `json:"type"`
	Weapon Weapon   `json:"weapon,omitempty"`
}

type Wall struct {
	A Position `json:"a"`
	B Position `json:"b"`
}

func (w *Wall) Center() Position {
	return Position{
		X: (w.A.X + w.B.X) / 2,
		Y: (w.A.Y + w.B.Y) / 2,
	}
}

func (w *Wall) Vector() Vector {
	return w.A.VectorTo(w.B)
}

type Map struct {
	Radius  float64
	Walls   []*Wall
	Items   []*Item
	Players []*Player
}
