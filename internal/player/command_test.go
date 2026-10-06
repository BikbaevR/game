package player

import (
	"math"
	"testing"
)

const epsilon = 1e-9

// terrainFunc превращает обычную функцию в Terrain (как http.HandlerFunc).
type terrainFunc func(x, y float64) bool

func (f terrainFunc) WalkableAt(x, y float64) bool { return f(x, y) }

// openField — местность без препятствий.
var openField = terrainFunc(func(x, y float64) bool { return true })

// assertPos проверяет позицию игрока с допуском на ошибки округления.
func assertPos(t *testing.T, p *Player, wantX, wantY float64) {
	t.Helper()
	// Условие записано как «не (разница <= допуск)», а не «разница > допуск»:
	// любое сравнение с NaN даёт false, и вариант «>» пропустил бы NaN в позиции.
	if !(math.Abs(p.X-wantX) <= epsilon) {
		t.Errorf("X: получили %v, ожидали %v", p.X, wantX)
	}
	if !(math.Abs(p.Y-wantY) <= epsilon) {
		t.Errorf("Y: получили %v, ожидали %v", p.Y, wantY)
	}
}

func TestPlayer_Apply_OpenField(t *testing.T) {
	const speed = 2
	diag := 10 + math.Sqrt2 // 10 + 2/√2

	cases := []struct {
		name         string
		dx, dy       int
		wantX, wantY float64
	}{
		{name: "вправо", dx: 1, dy: 0, wantX: 12, wantY: 10},
		{name: "влево", dx: -1, dy: 0, wantX: 8, wantY: 10},
		{name: "вверх", dx: 0, dy: -1, wantX: 10, wantY: 8},
		{name: "вниз", dx: 0, dy: 1, wantX: 10, wantY: 12},
		{name: "нет команды", dx: 0, dy: 0, wantX: 10, wantY: 10},
		{name: "диагональ", dx: 1, dy: 1, wantX: diag, wantY: diag},
		{name: "большие значения", dx: 50, dy: 0, wantX: 12, wantY: 10},
		{name: "большая диагональ", dx: 50, dy: 50, wantX: diag, wantY: diag},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := New(1, "test", 10, 10, speed)
			p.Apply(MoveCommand{PlayerID: 1, DX: c.dx, DY: c.dy}, openField)
			assertPos(t, p, c.wantX, c.wantY)
		})
	}
}

func TestPlayer_Apply_DiagonalStepLength(t *testing.T) {
	const speed = 2

	for _, dir := range [][2]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}, {3, 4}} {
		p := New(1, "test", 10, 10, speed)
		p.Apply(MoveCommand{PlayerID: 1, DX: dir[0], DY: dir[1]}, openField)

		dist := math.Hypot(p.X-10, p.Y-10)
		if math.Abs(dist-speed) > epsilon {
			t.Errorf("направление %v: шаг длиной %v, ожидали %v", dir, dist, float64(speed))
		}
	}
}

func TestPlayer_Apply_Collisions(t *testing.T) {
	const speed = 2

	// Стена справа: непроходимо всё, где x >= 100.
	wallRight := terrainFunc(func(x, y float64) bool { return x < 100 })

	// Одиночный блок 16x16 с левым верхним углом в (100, 100).
	block := terrainFunc(func(x, y float64) bool {
		return !(x >= 100 && x < 116 && y >= 100 && y < 116)
	})

	cases := []struct {
		name         string
		terrain      Terrain
		startX       float64
		startY       float64
		dx, dy       int
		wantX, wantY float64
	}{
		{
			name: "свободное место у стены: идём вправо", terrain: wallRight,
			startX: 90, startY: 50, dx: 1, dy: 0, wantX: 92, wantY: 50,
		},
		{
			// правый край хитбокса 93+6=99; шаг до 95 дал бы 101 — внутри стены
			name: "упёрся в стену: остаётся на месте", terrain: wallRight,
			startX: 93, startY: 50, dx: 1, dy: 0, wantX: 93, wantY: 50,
		},
		{
			name: "уход от стены разрешён", terrain: wallRight,
			startX: 93, startY: 50, dx: -1, dy: 0, wantX: 91, wantY: 50,
		},
		{
			name: "по диагонали вдоль стены: скользит по Y", terrain: wallRight,
			startX: 93, startY: 50, dx: 1, dy: 1, wantX: 93, wantY: 50 + math.Sqrt2,
		},
		{
			// центр (95, 95) вне блока, но угол хитбокса (101, 101) внутри
			name: "задевает блок углом, а не центром", terrain: block,
			startX: 93, startY: 95, dx: 1, dy: 0, wantX: 93, wantY: 95,
		},
		{
			name: "вертикально мимо блока", terrain: block,
			startX: 93, startY: 90, dx: 0, dy: 1, wantX: 93, wantY: 92,
		},
		{
			// внутри блока по y оказывается только верхний правый угол (93+6, 113-6)
			name: "правый верхний угол упирается в блок", terrain: block,
			startX: 93, startY: 113, dx: 1, dy: 0, wantX: 93, wantY: 113,
		},
		{
			// идём влево; внутри блока оказывается только верхний левый угол
			name: "левый верхний угол упирается в блок", terrain: block,
			startX: 123, startY: 113, dx: -1, dy: 0, wantX: 123, wantY: 113,
		},
		{
			// идём влево; внутри блока оказывается только нижний левый угол
			name: "левый нижний угол упирается в блок", terrain: block,
			startX: 123, startY: 103, dx: -1, dy: 0, wantX: 123, wantY: 103,
		},
		{
			name: "падаем сверху на блок: ось Y проверяется", terrain: block,
			startX: 108, startY: 93, dx: 0, dy: 1, wantX: 108, wantY: 93,
		},
		{
			name: "поднимаемся снизу в блок: ось Y проверяется", terrain: block,
			startX: 108, startY: 123, dx: 0, dy: -1, wantX: 108, wantY: 123,
		},
		{
			// X сдвигается и перекрывает колонку блока, и только после этого
			// шаг вниз упирается в блок: Y обязан учитывать уже сдвинутый X
			name: "диагональ: проверка Y использует новый X", terrain: block,
			startX: 93, startY: 93, dx: 1, dy: 1, wantX: 93 + math.Sqrt2, wantY: 93,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := New(1, "test", c.startX, c.startY, speed)
			p.Apply(MoveCommand{PlayerID: 1, DX: c.dx, DY: c.dy}, c.terrain)
			assertPos(t, p, c.wantX, c.wantY)
		})
	}
}
