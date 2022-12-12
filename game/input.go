package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (i *Input) Update(ship *Ship) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ship.x -= 1

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ship.x += 1
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ship.y -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ship.y += 1
	}
}
