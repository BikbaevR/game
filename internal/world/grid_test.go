package world

import "testing"

func TestGrid_SetGet(t *testing.T) {
	// Все координаты должны быть разными: иначе один случай законно затрёт другой.
	cases := []struct {
		name string
		x, y int
		typ  TileType
	}{
		{name: "левый верхний угол", x: 0, y: 0, typ: TileGrass},
		{name: "левая колонка, средняя строка", x: 0, y: 1, typ: TileSand},
		{name: "левый нижний угол", x: 0, y: 2, typ: TileGrass},
		{name: "центр", x: 1, y: 1, typ: TileSand},
		{name: "правый верхний угол", x: 2, y: 0, typ: TileSand},
		{name: "правый нижний угол", x: 2, y: 2, typ: TileGrass},
	}

	// Сетка одна на все случаи: баг в формуле индекса проявляется,
	// только когда одна клетка затирает другую.
	grid := NewGrid(3, 3)

	for _, c := range cases {
		if !grid.Set(c.x, c.y, Tile{Type: c.typ}) {
			t.Fatalf("%s: Set(%d, %d) вернул false для клетки внутри карты", c.name, c.x, c.y)
		}
	}

	// Читаем только после того, как записаны все клетки.
	for _, c := range cases {
		got, ok := grid.Get(c.x, c.y)
		if !ok {
			t.Errorf("%s: Get(%d, %d) вернул ok=false для клетки внутри карты", c.name, c.x, c.y)
			continue
		}
		if got.Type != c.typ {
			t.Errorf("%s: Get(%d, %d) вернул вид %v, ожидали %v", c.name, c.x, c.y, got.Type, c.typ)
		}
	}
}

func TestGrid_OutOfBounds(t *testing.T) {
	cases := []struct {
		name string
		x, y int
	}{
		{name: "левее карты", x: -1, y: 1},
		{name: "правее карты", x: 3, y: 1},
		{name: "выше карты", x: 1, y: -1},
		{name: "ниже карты", x: 1, y: 3},
	}

	grid := NewGrid(3, 3)

	for _, c := range cases {
		ok := grid.Set(c.x, c.y, Tile{Type: TileSand})
		if ok {
			t.Errorf("%s: Set(%d, %d) вернул %v, ожидали false для клетки во вне карты", c.name, c.x, c.y, ok)
		}
	}

	for _, c := range cases {
		_, ok := grid.Get(c.x, c.y)
		if ok {
			t.Errorf("%s: Get(%d, %d) вернул %v, ожидали false для клетки во вне карты", c.name, c.x, c.y, ok)
		}
	}
}

func TestGrid_WalkableAt(t *testing.T) {
	// Сетка 3x3: везде трава, кроме воды в клетке (1, 1).
	grid := NewGrid(3, 3)
	for x := range 3 {
		for y := range 3 {
			grid.Set(x, y, Tile{Type: TileGrass})
		}
	}
	grid.Set(1, 1, Tile{Type: TileWater})

	cases := []struct {
		name string
		x, y float64
		want bool
	}{
		{name: "трава в центре клетки", x: 0.5 * TileSize, y: 0.5 * TileSize, want: true},
		{name: "вода в центре клетки", x: 1.5 * TileSize, y: 1.5 * TileSize, want: false},
		{name: "последний пиксель травы перед водой", x: 1*TileSize - 0.1, y: 1.5 * TileSize, want: true},
		{name: "первый пиксель воды (граница клеток)", x: 1 * TileSize, y: 1.5 * TileSize, want: false},
		{name: "чуть левее карты", x: -0.5, y: 1.5 * TileSize, want: false},
		{name: "чуть выше карты", x: 1.5 * TileSize, y: -0.5, want: false},
		{name: "правее карты (граница)", x: 3 * TileSize, y: 0.5 * TileSize, want: false},
		{name: "ниже карты (граница)", x: 0.5 * TileSize, y: 3 * TileSize, want: false},
		{name: "последний пиксель карты", x: 3*TileSize - 0.1, y: 3*TileSize - 0.1, want: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := grid.WalkableAt(c.x, c.y); got != c.want {
				t.Errorf("WalkableAt(%v, %v) = %v, ожидали %v", c.x, c.y, got, c.want)
			}
		})
	}
}
