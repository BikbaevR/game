package game

import (
	"github.com/BikbaevR/game/internal/player"
	"github.com/BikbaevR/game/internal/render"
	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world  *world.Grid
	player *player.Player
}

func New(w *world.Grid, p *player.Player) *Game {
	return &Game{w, p}
}

func (g *Game) Update() error {
	cmd := readMoveCommand(g.player.ID)
	g.player.Apply(cmd, g.world)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DrawGrid(screen, g.world)
	render.DrawPlayer(screen, g.player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.world.Width() * world.TileSize, g.world.Height() * world.TileSize
}
