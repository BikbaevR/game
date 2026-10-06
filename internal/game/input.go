package game

import (
	"github.com/BikbaevR/game/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
)

func readMoveCommand(id player.ID) player.MoveCommand {
	cmd := player.MoveCommand{PlayerID: id}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		cmd.DY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		cmd.DY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cmd.DX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		cmd.DX++
	}
	return cmd
}
