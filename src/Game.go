package src

import (
	"TimberGO/assets"
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth         float64 = 600
	ScreenHeight        float64 = 720
	WorldWidth          float64 = 15
	WorldHeight         float64 = 18
	DeltaTime           float64 = 1.0 / 60
	MidChange           float32 = 0.5
	TrunkCount          int     = 5
	TrunkHeight         float64 = 5
	TrunkWidth          float64 = 193 / 89 * TrunkHeight
	TrunkBasePosY       float64 = 14 - TrunkHeight/2
	TimberHeight        float64 = 7
	TimberWidth         float64 = TimberHeight / 67 * 76
	TimberDeltaX        float64 = 3.5
	MaximumSpeed        float64 = 5
	TimberY             float64 = 14.3 - TrunkHeight/2
	Acceleration        float64 = 0.00091
	TrunkThrowingSpeedX float64 = 14
	TrunkThrowingSpeedY float64 = 3
	TrunkThrowingRStep  float64 = 130
	TrunkDroppingSpeedY float64 = 5 / 0.15
	HealthBarY          float64 = 1.25
	HealthBarWidth      float64 = 12
	HealthBarHeight     float64 = 40 * HealthBarWidth / 350
	HealthLossSpeed     float64 = 30
	HealthBonusPerTrunk float64 = 10
	BonusScorePerTrunk  float64 = 10
)

var (
	TimberAniDuration = []uint16{30, 30, 40, 20}
	font              = assets.LoadFont()
)

type Game struct {
	Background     *GameObject
	ThrowingTrunks []*GameObject
	Trunks         []Trunk
	Timber         *Timber
	HealthBar      *HealthBar
	Score          int
	CurrentSpeed   float64
	Failed         bool
}

func (g *Game) GameOver() {
	log.Println(`Game Over!`)
	g.Failed = true
}

func (g *Game) DestroyIfOver() {
	newArr := make([]*GameObject, 0)
	for _, gObj := range g.ThrowingTrunks {
		if gObj.Position.X-gObj.IWWidth/2 < WorldWidth && gObj.Position.X+gObj.IWWidth/2 >= 0 &&
			gObj.Position.Y-gObj.IWHeight/2 < WorldHeight && gObj.Position.Y+gObj.IWHeight/2 >= 0 {
			newArr = append(newArr, gObj)
		}
	}
	g.ThrowingTrunks = newArr
}

func (g *Game) Update() error {
	if g.Score > 0 {
		if g.CurrentSpeed < MaximumSpeed {
			g.CurrentSpeed *= 1 + Acceleration
		} else {
			g.CurrentSpeed = MaximumSpeed
		}
	}
	minY := TrunkBasePosY
	for _, trunk := range g.Trunks {
		trunk.CalcPhysics()
		if trunk.Position.Y > minY {
			trunk.Position.Y = minY
			trunk.Velocity = Vector{}.Zero()
		}
		minY -= TrunkHeight
	}
	if g.Failed {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.New()
		}
	} else {
		if g.HealthBar.Loss(g) {
			g.GameOver()
			return nil
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			if g.Timber.Direction == 1 {
				g.Timber.Flip = true
				g.Timber.Direction = -1
				g.Timber.Position.X = WorldWidth/2 - TimberDeltaX
			}
			g.Timber.CutDown(g)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			if g.Timber.Direction == -1 {
				g.Timber.Flip = false
				g.Timber.Direction = 1
				g.Timber.Position.X = WorldWidth/2 + TimberDeltaX
			}
			g.Timber.CutDown(g)
		}
		g.Timber.GameObject.Image = g.Timber.Next()
	}
	for _, obj := range g.ThrowingTrunks {
		obj.CalcPhysics()
	}
	g.DestroyIfOver()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.DrawOnScreen(screen)
	for _, trunk := range g.Trunks {
		trunk.DrawOnScreen(screen)
	}
	for _, obj := range g.ThrowingTrunks {
		obj.DrawOnScreen(screen)
	}
	g.HealthBar.DrawOnScreen(screen)
	if !g.Failed {
		g.Timber.DrawOnScreen(screen)
	} else {

	}
	g.DrawScore(screen)
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	s := fmt.Sprintf("Scores %d", g.Score)
	rec := text.BoundString(font, s).Size()
	w, h := screen.Size()
	text.Draw(screen, s, font, (w-rec.X)/2, int(g.HealthBar.Border.Position.Y*2/WorldHeight*float64(h)), color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	NewWidth, NewHeight := ScreenWidth/ScreenHeight*float64(outsideHeight), ScreenHeight/ScreenWidth*float64(outsideWidth)
	return int(math.Min(float64(outsideWidth), NewWidth)), int(math.Min(float64(outsideHeight), NewHeight))
}

func Init() *Game {
	g := &Game{
		Background: &GameObject{
			Position: Point{
				X: WorldWidth / 2,
				Y: WorldHeight / 2,
			},
			Flip:     false,
			Rotation: 0,
			RStep:    0,
			Velocity: Vector{}.Zero(),
			IWWidth:  WorldWidth,
			IWHeight: WorldHeight,
			Image:    assets.LoadImage(`background.png`),
		},
	}
	return g
}

func (g *Game) New() *Game {
	g.Failed = false
	g.CurrentSpeed = 1
	g.HealthBar = &HealthBar{
		Border: &GameObject{
			Velocity: Vector{}.Zero(),
			Position: Point{
				X: WorldWidth / 2,
				Y: HealthBarY,
			},
			Flip:     false,
			IWWidth:  HealthBarWidth,
			IWHeight: HealthBarHeight,
			Image:    assets.LoadImage(`healthBar.png`),
			Rotation: 0,
			RStep:    0,
		},
		Color:        color.RGBA{G: 255, B: 64, A: 255},
		Value:        100,
		MaximumValue: 100,
	}
	g.Timber = &Timber{
		GameObject: &GameObject{
			Velocity: Vector{}.Zero(),
			Position: Point{
				X: WorldWidth/2 + TimberDeltaX,
				Y: TimberY,
			},
			Flip:     false,
			IWWidth:  TimberWidth,
			IWHeight: TimberHeight,
			Image:    assets.LoadImage(`man1.png`),
			Rotation: 0,
			RStep:    0,
		},
		Animation: &Animation{
			Images: []*ebiten.Image{
				assets.LoadImage(`man1.png`),
				assets.LoadImage(`man2.png`),
				assets.LoadImage(`man3.png`),
				assets.LoadImage(`man4.png`)},
			Duration: TimberAniDuration,
			State:    0,
			Stop:     true,
			LastTime: time.Now(),
		},
		Direction: 1,
	}
	g.ThrowingTrunks = []*GameObject{}
	g.Trunks = TrunkGenerator{}.Init()
	g.Score = 0
	return g
}
