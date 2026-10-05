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

	grid := world.NewGrid(20, 15)

	for i := 0; i <= 19; i++ {
		for j := 0; j <= 14; j++ {
			tile := world.Tile{Type: getRandomTile()}
			grid.Set(i, j, tile)
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
