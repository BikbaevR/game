package player

// Size — сторона квадратного хитбокса игрока в пикселях.
// Меньше тайла, чтобы игрок проходил в проход шириной в один тайл.
const Size = 12

type ID uint32

type Player struct {
	ID    ID
	Name  string
	X, Y  float64 // центр игрока в пикселях
	Speed float64 // пикселей за тик
}

func New(id ID, name string, x, y float64, speed float64) *Player {
	return &Player{
		ID:    id,
		Name:  name,
		X:     x,
		Y:     y,
		Speed: speed,
	}
}
