package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	ship  *Ship
	cfg   *Config
}

func NewGame() *Game {
	cfg := LoadConfig()
	ship := NewShip(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		input: &Input{},
		ship:  ship,
		cfg:   cfg,
	}
}

func (g *Game) Update() error {
	g.input.Update(g.ship)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}
