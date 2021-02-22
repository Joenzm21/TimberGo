package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Animation struct {
	Images   []*ebiten.Image
	Duration []uint16
	State    int
	LastTime time.Time
	Stop     bool
}

func (a *Animation) Start() {
	a.LastTime = time.Now()
	a.Stop = false
	a.State = 0
}

func (a *Animation) Next() *ebiten.Image {
	if a.Stop {
		a.State = 0
	} else if time.Now().Sub(a.LastTime).Milliseconds() >= int64(a.Duration[a.State]) {
		a.State = (a.State + 1) % len(a.Images)
		a.LastTime = time.Now()
		if a.State == 0 {
			a.Stop = true
		}
	}
	return a.Images[a.State]
}
