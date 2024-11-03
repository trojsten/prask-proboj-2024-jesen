package main

import "math"

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// SquaredDistance calculates the squared distance between two Position points
func (p Position) SquaredDistance(p2 Position) float64 {
	var dx = p.X - p2.X
	var dy = p.Y - p2.Y
	return dx*dx + dy*dy
}

func (p Position) Distance(p2 Position) float64 {
	return math.Sqrt(p.SquaredDistance(p2))
}

func (p Position) VectorTo(target Position) Vector {
	return Vector{target.X - p.X, target.Y - p.Y}
}

func (p Position) Add(v2 Vector) Position {
	return Position{X: p.X + v2.X, Y: p.Y + v2.Y}
}

type Vector Position

func (v Vector) Normalize() Vector {
	vectorLen := v.Length()
	return Vector{X: v.X / vectorLen, Y: v.Y / vectorLen}
}

func (v Vector) Mul(k float64) Vector {
	return Vector{X: v.X * k, Y: v.Y * k}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) DotProduct(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector) Angle(target Vector) float64 {
	return math.Acos(v.DotProduct(target) / (v.Length() * target.Length()))
}

func (v Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func Intesect(a1, a2, b1, b2 Position) (Position, bool) {
	t1 := (a1.X-b1.X)*(b2.Y-b1.Y) - (a1.Y-b1.Y)*(b1.X-b2.X)
	t2 := (a1.X-a2.X)*(b2.Y-b1.Y) - (a1.Y-a2.Y)*(b1.X-b2.X)
	t := t1 / t2

	u1 := (a1.X-a2.X)*(a1.Y-b1.Y) - (a1.Y-a2.Y)*(a1.X-b1.X)
	u2 := (a1.X-a2.X)*(b2.Y-b1.Y) - (a1.Y-a2.Y)*(b1.X-b2.X)
	u := u1 / u2

	intersects := 0 <= t && t <= 1 && 0 <= u && u <= 1
	point := Position{
		X: a1.X + t*(a2.X-a1.X),
		Y: a1.Y + t*(a2.Y-a1.Y),
	}
	if intersects {
		return point, true
	}
	return Position{}, false
}
