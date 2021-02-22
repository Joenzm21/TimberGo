package src

type Point struct {
	X, Y float64
}

func (Point) Zero() Point {
	return Point{
		X: 0,
		Y: 0,
	}
}
type Vector struct {
	X, Y float64
}

func (Vector) Zero() Vector {
	return Vector{
		X: 0,
		Y: 0,
	}
}

type Trunk struct {
	*GameObject
	TrunkType int8
}
