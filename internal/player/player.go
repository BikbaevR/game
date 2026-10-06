package player

type ID uint32

type Player struct {
	ID    ID
	Name  string
	X, Y  float64
	Speed float64
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
