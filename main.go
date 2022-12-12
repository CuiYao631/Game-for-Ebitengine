package main

import (
	"Game-for-Ebitengine/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
