package main

import (
	"log"
	"math/rand"

	"github.com/BikbaevR/game/internal/game"
	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	//ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	grid := world.NewGrid(20, 15)

	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			tile := world.Tile{Type: getRandomTile()}
			grid.Set(x, y, tile)
		}
	}

	worldGrid := game.New(grid)
	if err := ebiten.RunGame(worldGrid); err != nil {
		log.Fatal(err)
	}
}

func getRandomTile() world.TileType {
	randomIndex := rand.Intn(3)
	if randomIndex == 0 {
		return world.TileGrass
	} else if randomIndex == 1 {
		return world.TileWater
	}

	return world.TileSand
}
