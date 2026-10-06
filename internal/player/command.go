package player

import "math"

type MoveCommand struct {
	PlayerID ID
	DX, DY   int
}

func (p *Player) Apply(cmd MoveCommand) {
	if cmd.DX == 0 && cmd.DY == 0 {
		return
	}

	dx := float64(cmd.DX)
	dy := float64(cmd.DY)

	length := math.Sqrt(dx*dx + dy*dy)

	p.X += dx / length * p.Speed
	p.Y += dy / length * p.Speed
}
