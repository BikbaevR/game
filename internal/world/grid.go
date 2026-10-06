package world

import "math"

type Grid struct {
	width, height int
	tiles         []Tile
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		width:  width,
		height: height,
		tiles:  make([]Tile, width*height),
	}
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Get(x, y int) (Tile, bool) {
	if !g.InBounds(x, y) {
		return Tile{}, false
	}
	return g.tiles[g.index(x, y)], true
}

func (g *Grid) Set(x, y int, t Tile) bool {
	if !g.InBounds(x, y) {
		return false
	}
	g.tiles[g.index(x, y)] = t
	return true
}

func (g *Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

func (g *Grid) index(x, y int) int {
	return y*g.width + x
}

func (g *Grid) WalkableAt(x, y float64) bool {
	tile, ok := g.Get(int(math.Floor(x/TileSize)), int(math.Floor(y/TileSize)))
	return ok && tile.Type.Info().Walkable
}
