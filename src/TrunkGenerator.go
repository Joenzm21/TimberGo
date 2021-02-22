package src

import (
	"TimberGo/assets"
	"math/rand"
	"time"
)

type TrunkGenerator struct{}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (TrunkGenerator) Init() []Trunk {
	res := make([]Trunk, TrunkCount)
	basePos := Point{
		X: WorldWidth / 2,
		Y: TrunkBasePosY,
	}
	res[0] = Trunk{
		GameObject: &GameObject{
			Position: basePos,
			Rotation: 0,
			RStep:    0,
			IWWidth:  TrunkWidth,
			IWHeight: TrunkHeight,
			Velocity: Vector{}.Zero(),
			Image:    assets.LoadImage(`mid.png`),
		},
		TrunkType: 0,
	}
	for i := 1; i < TrunkCount; i++ {
		basePos.Y -= TrunkHeight
		res[i] = TrunkGenerator{}.CreateTrunk(basePos)
	}
	return res
}

func (TrunkGenerator) CreateTrunk(position Point) Trunk {
	res := Trunk{
		GameObject: &GameObject{
			Position: position,
			Rotation: 0,
			RStep:    0,
			Flip:     false,
			IWWidth:  TrunkWidth,
			IWHeight: TrunkHeight,
			Velocity: Vector{}.Zero(),
		},
	}
	if r.Float32() < MidChange {
		res.TrunkType = 0
		res.Image =assets.LoadImage(`mid.png`)
	} else if r.Float32() < 0.5 {
		res.TrunkType = -1
		res.Image =assets.LoadImage(`left.png`)
	} else {
		res.TrunkType = 1
		res.Image =assets.LoadImage(`right.png`)
	}
	return res
}
