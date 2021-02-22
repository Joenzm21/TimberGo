package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type HealthBar struct {
	Border              *GameObject
	Color               color.Color
	Value, MaximumValue float64
}

func (h *HealthBar) Loss(g *Game) bool {
	if g.Score == 0 {
		return false
	}
	h.Value = math.Max(0, h.Value-HealthLossSpeed*DeltaTime*g.CurrentSpeed)
	return h.Value == 0
}

func (h *HealthBar) Increase(v float64) {
	h.Value = math.Min(h.MaximumValue, h.Value+v)
}

func (h *HealthBar) DrawOnScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	img := ebiten.NewImage(1, 1)
	img.Fill(h.Color)
	ScrWidth, ScrHeight := screen.Size()
	op.GeoM.Scale(float64(ScrWidth)*HealthBarWidth/WorldWidth*h.Value/h.MaximumValue*0.905, float64(ScrHeight)*HealthBarHeight*17/20/WorldHeight)
	op.GeoM.Translate((h.Border.Position.X-HealthBarWidth/2*0.83)*float64(ScrWidth)/WorldWidth, (h.Border.Position.Y-HealthBarHeight*17/40)*float64(ScrHeight)/WorldHeight)
	screen.DrawImage(img, op)
	h.Border.DrawOnScreen(screen)
}
