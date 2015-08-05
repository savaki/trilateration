//	Copyright 2015 Matt Ho
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

package trilateration

import (
	"fmt"
	"math"
)

var (
	ErrNoSolution = fmt.Errorf("no solution found")
)

type Point struct {
	X float64
	Y float64
	Z float64
	R float64
}

type Solution []Point

func (s Solution) First() Point {
	return s[0]
}

func square(v float64) float64 {
	return v * v
}

func normalize(p Point) float64 {
	return math.Sqrt(square(p.X) + square(p.Y) + square(p.Z))
}

func dot(p1, p2 Point) float64 {
	return p1.X*p2.X + p1.Y*p2.Y + p1.Z*p2.Z
}

func subtract(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
		Z: p1.Z - p2.Z,
	}
}

func add(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func divide(p Point, v float64) Point {
	return Point{
		X: p.X / v,
		Y: p.Y / v,
		Z: p.Z / v,
	}
}

func multiply(p Point, v float64) Point {
	return Point{
		X: p.X * v,
		Y: p.Y * v,
		Z: p.Z * v,
	}
}

func cross(p1, p2 Point) Point {
	return Point{
		X: p1.Y*p2.Z - p1.Z*p2.Y,
		Y: p1.Z*p2.X - p1.X*p2.Z,
		Z: p1.X*p2.Y - p1.Y*p2.X,
	}
}

func Solve(p1, p2, p3 Point) (Solution, error) {
	ex := divide(subtract(p2, p1), normalize(subtract(p2, p1)))
	i := dot(ex, subtract(p3, p1))
	a := subtract(subtract(p3, p1), multiply(ex, i))
	ey := divide(a, normalize(a))
	d := normalize(subtract(p2, p1))
	j := dot(ey, subtract(p3, p1))

	x := (square(p1.R) - square(p2.R) + square(d)) / (2 * d)
	y := (square(p1.R)-square(p3.R)+square(i)+square(j))/(2*j) - (i/j)*x
	z := math.Sqrt(square(p1.R) - square(x) - square(y))

	if math.IsNaN(z) {
		return nil, ErrNoSolution
	}

	p4 := add(p1, add(multiply(ex, x), multiply(ey, y)))
	if z == 0 {
		return []Point{p4}, nil
	}

	ez := cross(ex, ey)
	p4a := add(a, multiply(ez, z))
	p4b := subtract(a, multiply(ez, z))
	return Solution{p4, p4a, p4b}, nil
}
