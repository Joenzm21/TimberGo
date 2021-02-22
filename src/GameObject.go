package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type GameObject struct {
	Position          Point
	IWWidth, IWHeight float64
	Velocity          Vector
	Rotation, RStep   float64
	Flip              bool
	Image             *ebiten.Image
}

func (gObj *GameObject) CalcPhysics() {
	gObj.Position.X += gObj.Velocity.X * DeltaTime
	gObj.Position.Y += gObj.Velocity.Y * DeltaTime
	gObj.Rotation += gObj.RStep * DeltaTime
}

func (gObj *GameObject) DrawOnScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	ScrWidth, ScrHeight := screen.Size()
	ImgWidth, ImgHeight := gObj.Image.Size()
	op.GeoM.Translate(-float64(ImgWidth)/2, -float64(ImgHeight)/2)
	op.GeoM.Rotate(gObj.Rotation * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(ImgWidth)/2, float64(ImgHeight)/2)
	if gObj.Flip {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(ImgWidth), 0)
	}
	op.GeoM.Scale(float64(ScrWidth)*gObj.IWWidth/float64(ImgWidth)/WorldWidth, float64(ScrHeight)*gObj.IWHeight/float64(ImgHeight)/WorldHeight)
	op.GeoM.Translate((gObj.Position.X-gObj.IWWidth/2)*float64(ScrWidth)/WorldWidth, (gObj.Position.Y-gObj.IWHeight/2)*float64(ScrHeight)/WorldHeight)
	screen.DrawImage(gObj.Image, op)
}
