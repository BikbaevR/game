package player

import "math"

type MoveCommand struct {
	PlayerID ID
	DX, DY   int
}

func (p *Player) Apply(cmd MoveCommand, t Terrain) {
	if cmd.DX == 0 && cmd.DY == 0 {
		return
	}

	dx := float64(cmd.DX)
	dy := float64(cmd.DY)

	length := math.Sqrt(dx*dx + dy*dy)

	stepX := dx / length * p.Speed
	stepY := dy / length * p.Speed

	if canStandAt(t, p.X+stepX, p.Y) {
		p.X += stepX
	}
	if canStandAt(t, p.X, p.Y+stepY) {
		p.Y += stepY
	}
}
