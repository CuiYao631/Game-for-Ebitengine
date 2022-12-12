package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64
	y      float64
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./game/image/ship.png")
	if err != nil {
		log.Fatal(err)
	}
	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}
	return ship
}
func (s *Ship) Draw(screen *ebiten.Image, cfg *Config) {

	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(float64(cfg.ScreenWidth-s.width)/2, float64(cfg.ScreenHeight-s.height))
	if s.x < 0 {
		s.x = 0
	}
	if s.y < 0 {
		s.y = 0
	}
	if s.y > 192 {
		s.y = 192
	}
	if s.x > 259 {
		s.x = 259
	}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.image, op)
}
