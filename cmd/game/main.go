package main

import (
	"log"
	"math/rand"

	"github.com/BikbaevR/game/internal/game"
	"github.com/BikbaevR/game/internal/player"
	"github.com/BikbaevR/game/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerName  = "test"
	playerSpeed = 1.2 // пикселей за тик
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")

	grid := world.NewGrid(20, 15)

	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			tile := world.Tile{Type: getRandomTile()}
			grid.Set(x, y, tile)
		}
	}

	// Игрок появляется в центре клетки посередине карты; клетку делаем травой,
	// чтобы он не оказался внутри воды.
	spawnTileX, spawnTileY := grid.Width()/2, grid.Height()/2
	grid.Set(spawnTileX, spawnTileY, world.Tile{Type: world.TileGrass})

	spawnX := float64(spawnTileX*world.TileSize + world.TileSize/2)
	spawnY := float64(spawnTileY*world.TileSize + world.TileSize/2)
	p := player.New(1, playerName, spawnX, spawnY, playerSpeed)

	g := game.New(grid, p)
	if err := ebiten.RunGame(g); err != nil {
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
