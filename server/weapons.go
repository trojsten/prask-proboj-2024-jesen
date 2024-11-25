package main

type Weapon int

const (
	WeaponNone Weapon = iota
	WeaponKnife
	WeaponPistol
	WeaponTommy
)

type WeaponStats struct {
	Range        float64
	Damage       int
	ReloadTime   int
	AmmoCapacity int
}

var WeaponStatsMap = map[Weapon]WeaponStats{
	WeaponNone:   {},
	WeaponKnife:  {10, 34, 0, 1},
	WeaponPistol: {PlayerMovementRange * 5, 5, 2, 10},
	WeaponTommy:  {PlayerMovementRange * 10, 8, 4, 25},
}
