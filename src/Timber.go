package src

import (
	"TimberGo/assets"
)

type Timber struct {
	*GameObject
	*Animation
	Direction int8
}

func (t *Timber) CutDown(g *Game) {
	if t.Direction == g.Trunks[1].TrunkType {
		g.GameOver()
	} else {
		t.Start()
	}
	g.Score += int(BonusScorePerTrunk * g.CurrentSpeed)
	g.HealthBar.Increase(g.CurrentSpeed * HealthBonusPerTrunk)
	if t.Direction == 1 {
		g.Trunks[0].Velocity = Vector{
			X: -TrunkThrowingSpeedX * g.CurrentSpeed,
			Y: TrunkThrowingSpeedY,
		}
		g.Trunks[0].RStep = -TrunkThrowingRStep
	} else {
		g.Trunks[0].Velocity = Vector{
			X: TrunkThrowingSpeedX * g.CurrentSpeed,
			Y: TrunkThrowingSpeedY,
		}
		g.Trunks[0].RStep = TrunkThrowingRStep * g.CurrentSpeed
	}
	g.ThrowingTrunks = append(g.ThrowingTrunks, g.Trunks[0].GameObject)
	lastPosY := g.Trunks[len(g.Trunks)-1].Position.Y
	g.Trunks = append(g.Trunks[1:], TrunkGenerator{}.CreateTrunk(Point{
		X: WorldWidth / 2,
		Y: lastPosY,
	}))
	g.Trunks[0].TrunkType = 0
	g.Trunks[0].Image = assets.LoadImage(`mid.png`)
	for _, trunk := range g.Trunks {
		trunk.Velocity.Y = TrunkDroppingSpeedY
	}
}
