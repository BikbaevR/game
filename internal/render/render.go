package render

import (
	"image/color"

	"github.com/BikbaevR/game/internal/player"
	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var tileColors = [...]color.RGBA{
	world.TileSand:  {R: 207, G: 212, B: 71, A: 255},
	world.TileGrass: {R: 71, G: 212, B: 118, A: 255},
	world.TileWater: {R: 71, G: 174, B: 212, A: 255},
}

// DrawGrid рисует все клетки сетки цветными квадратами.
func DrawGrid(screen *ebiten.Image, grid *world.Grid) {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			tile, ok := grid.Get(x, y)
			if !ok {
				continue
			}
			vector.FillRect(
				screen,
				float32(x*world.TileSize), float32(y*world.TileSize),
				world.TileSize, world.TileSize,
				tileColors[tile.Type],
				false,
			)
		}
	}
}

// DrawPlayer рисует игрока квадратом player.Size; X, Y игрока — его центр.
func DrawPlayer(screen *ebiten.Image, p *player.Player) {
	vector.FillRect(
		screen,
		float32(p.X-player.Size/2), float32(p.Y-player.Size/2),
		player.Size, player.Size,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		false,
	)
}
