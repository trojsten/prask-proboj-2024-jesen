package main

import (
	"math"
)

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

func (v Vector) CP(w Vector) float64 {
	return v.X*w.Y - v.Y*w.X
}

func Intesect(a1, a2, b1, b2 Position) (Position, bool) {
	// https://stackoverflow.com/a/565282
	vecA := a1.VectorTo(a2)
	vecB := b1.VectorTo(b2)

	if vecA.Length() == 0 {
		return Position{}, false
	}
	vecX := a1.VectorTo(b1)

	// b1 = q, a1 = p
	// s = B
	// r = A

	if vecA.CP(vecB) == 0 {
		if vecX.CP(vecA) == 0 {
			// t0 = (q − p) · r / (r · r)
			t0 := vecX.DotProduct(vecA) / vecA.DotProduct(vecA)
			// t1 = (q + s − p) · r / (r · r) = t0 + s · r / (r · r)
			t1 := a1.VectorTo(b2).DotProduct(vecA) / vecA.DotProduct(vecA)

			i0 := a1.Add(vecA.Mul(t0))
			i1 := a1.Add(vecA.Mul(t1))

			intersect := i0
			if i0.Distance(a1) > i1.Distance(a1) {
				intersect = i1
			}

			return intersect, true
		}
		return Position{}, false
	}

	t := vecX.CP(vecB) / vecA.CP(vecB)
	u := vecX.CP(vecA) / vecA.CP(vecB)

	if 0 <= t && t <= 1 && 0 < u && u < 1 {
		return a1.Add(vecA.Mul(t)), true
	}

	return Position{}, false
}
