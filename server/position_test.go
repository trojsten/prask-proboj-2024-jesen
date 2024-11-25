package main

import (
	"math"
	"reflect"
	"testing"
)

func TestPosition_Add(t *testing.T) {
	tests := []struct {
		name string
		p1   Position
		v    Vector
		want Position
	}{
		// TODO: Add test cases.
		{
			name: "add posit",
			p1:   Position{X: 0, Y: 0},
			v:    Vector{X: 10, Y: 10},
			want: Position{X: 10, Y: 10},
		},
		{
			name: "add posit",
			p1:   Position{X: 7, Y: 2},
			v:    Vector{X: -10, Y: 50},
			want: Position{X: -3, Y: 52},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p1.Add(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_SquaredDistance(t *testing.T) {
	tests := []struct {
		name string
		p1   Position
		p2   Position
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "classic",
			p1:   Position{X: 0, Y: 0},
			p2:   Position{X: 10, Y: 10},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p1.SquaredDistance(tt.p2); got != tt.want {
				t.Errorf("SquaredDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_VectorTo(t *testing.T) {
	tests := []struct {
		name string
		p1   Position
		p2   Position
		want Vector
	}{
		// TODO: Add test cases.
		{
			name: "classic",
			p1:   Position{0, 0},
			p2:   Position{3, 2},
			want: Vector{3, 2},
		}, {
			name: "classic reversed",
			p1:   Position{4, -1},
			p2:   Position{0, 0},
			want: Vector{-4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p1.VectorTo(tt.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Angle(t *testing.T) {
	tests := []struct {
		name string
		v1   Vector
		v2   Vector
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "right angle",
			v1:   Vector{X: 0, Y: 10},
			v2:   Vector{X: 10, Y: 0},
			want: math.Pi / 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Angle(tt.v2); got != tt.want {
				t.Errorf("Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_DotProduct(t *testing.T) {
	tests := []struct {
		name string
		v1   Vector
		v2   Vector
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "classic",
			v1:   Vector{1, 1},
			v2:   Vector{5, -6},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.DotProduct(tt.v2); got != tt.want {
				t.Errorf("DotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_IsZero(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want bool
	}{
		{"zero", Vector{0, 0}, true},
		{"non-zero", Vector{12, -5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Length(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
		{"1", Vector{0, 0}, 0},
		{"2", Vector{100, 0}, 100},
		{"3", Vector{2, 2}, math.Sqrt(8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Mul(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		k    float64
		want Vector
	}{
		{"1", Vector{1, 2}, -1, Vector{-1, -2}},
		{"2", Vector{10, 20}, 1.5, Vector{15, 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mul(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Normalize(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want Vector
	}{
		{"normal", Vector{1, 0}, Vector{1, 0}},
		{"1", Vector{1, 1}, Vector{math.Sqrt(2) / 2, math.Sqrt(2) / 2}},
		{"2", Vector{100, 0}, Vector{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Normalize()
			eps := math.Pow(10, -5)
			if math.Abs(got.X-tt.want.X) > eps || math.Abs(got.Y-tt.want.Y) > eps {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntesect(t *testing.T) {
	type args struct {
		a1 Position
		a2 Position
		b1 Position
		b2 Position
	}
	tests := []struct {
		name  string
		args  args
		want  Position
		want1 bool
	}{
		{"none", args{
			Position{0, 0},
			Position{100, 0},
			Position{10, 20},
			Position{110, 20},
		}, Position{}, false},
		{"1", args{
			Position{0, 0},
			Position{100, 0},
			Position{50, -20},
			Position{50, 20},
		}, Position{50, 0}, true},
		{"bug", args{
			Position{-10, 8},
			Position{3.6463650072002203, 3.6463650072002203},
			Position{0, 0},
			Position{250, 250},
		}, Position{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Intesect(tt.args.a1, tt.args.a2, tt.args.b1, tt.args.b2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intesect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Intesect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
