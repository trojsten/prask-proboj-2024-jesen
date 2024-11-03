package main

type Weapon int

const (
	WeaponNone = iota
)

type WeaponStats struct {
	Range        float64
	Damage       int
	ReloadTime   int
	AmmoCapacity int
}

var WeaponStatsMap = map[Weapon]WeaponStats{
	WeaponNone: {},
}
