package render

import (
	"image/color"

	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TileSize = 16

var tileColors = [...]color.RGBA{
	world.TileSand:  {R: 207, G: 212, B: 71, A: 255},
	world.TileGrass: {R: 71, G: 212, B: 118, A: 255},
	world.TileWater: {R: 71, G: 174, B: 212, A: 255},
}

func DrawGrid(screen *ebiten.Image, grid *world.Grid) {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			tile, ok := grid.Get(x, y)
			if !ok {
				continue
			}
			vector.FillRect(
				screen,
				float32(x*TileSize), float32(y*TileSize),
				TileSize, TileSize,
				tileColors[tile.Type],
				false,
			)
		}
	}
}
