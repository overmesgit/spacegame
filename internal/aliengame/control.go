package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action = int8

const (
	MoveUp = iota + 1
	MoveDown
	MoveLeft
	MoveRight
	Shoot
	BigShoot
)

func Control(keys []ebiten.Key) []Action {
	actions := make([]Action, 0)

	vx := 0
	vy := 0
	for _, key := range keys {
		switch key {

		case ebiten.KeyLeft:
			vx += -1
		case ebiten.KeyRight:
			vx += 1
		case ebiten.KeyUp:
			vy += -1
		case ebiten.KeyDown:
			vy += 1
		case ebiten.KeyZ:
			actions = append(actions, Shoot)
		case ebiten.KeySpace:
			actions = append(actions, BigShoot)
		}

	}

	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) || inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		vx = 0
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		vy = 0
	}
	if vx > 0 {
		actions = append(actions, MoveRight)
	} else if vx < 0 {
		actions = append(actions, MoveLeft)
	}
	if vy > 0 {
		actions = append(actions, MoveDown)
	} else if vy < 0 {
		actions = append(actions, MoveUp)
	}
	return actions
}
