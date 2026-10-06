package world

type TileType uint8

// TileSize — размер стороны одного тайла в пикселях.
const TileSize = 16

const (
	TileWater TileType = iota //замена enum, первый элемент 0 дальше 1, 2, 3 и так далее
	TileGrass
	TileSand
)

type TileInfo struct {
	Name     string
	Walkable bool
}

var tileInfos = [...]TileInfo{
	TileWater: {Name: "Вода", Walkable: false},
	TileGrass: {Name: "Трава", Walkable: true},
	TileSand:  {Name: "Песок", Walkable: true},
}

func (t TileType) Info() TileInfo {
	return tileInfos[t]
}

type Tile struct {
	Type TileType
}
