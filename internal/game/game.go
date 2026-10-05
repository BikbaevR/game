package game

import (
	"github.com/BikbaevR/game/internal/render"
	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *world.Grid
}

func New(w *world.Grid) *Game {
	return &Game{w}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DrawGrid(screen, g.world)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.world.Width() * render.TileSize, g.world.Height() * render.TileSize
}
