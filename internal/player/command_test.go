package player

import (
	"testing"
)

func TestCommand_Apply(t *testing.T) {

	cases := []struct {
		name                             string
		currentX, currentY               float64
		dx, dy                           int
		expectedResultX, expectedResultY float64
		speed                            float64
	}{
		{name: "вправо", currentX: 10, currentY: 10, dx: 1, dy: 0, expectedResultX: 12, expectedResultY: 10, speed: 2},
		{name: "влево", currentX: 10, currentY: 10, dx: -1, dy: 0, expectedResultX: 8, expectedResultY: 10, speed: 2},
		{name: "вверх", currentX: 10, currentY: 10, dx: 0, dy: -1, expectedResultX: 10, expectedResultY: 8, speed: 2},
		{name: "вниз", currentX: 10, currentY: 10, dx: 0, dy: 1, expectedResultX: 10, expectedResultY: 12, speed: 2},
		{name: "нет команды", currentX: 10, currentY: 10, dx: 0, dy: 0, expectedResultX: 10, expectedResultY: 10, speed: 2},
		{name: "диагональ", currentX: 10, currentY: 10, dx: 1, dy: 1, expectedResultX: 11.414213562373096, expectedResultY: 11.414213562373096, speed: 2},
		{name: "большие значения", currentX: 10, currentY: 10, dx: 50, dy: 0, expectedResultX: 12, expectedResultY: 10, speed: 2},
		{name: "большая диагональ", currentX: 10, currentY: 10, dx: 50, dy: 50, expectedResultX: 11.414213562373096, expectedResultY: 11.414213562373096, speed: 2},
	}

	for _, c := range cases {
		p := New(1, "test", c.currentX, c.currentY, c.speed)
		command := MoveCommand{PlayerID: 1, DX: c.dx, DY: c.dy}
		p.Apply(command)

		if p.X != c.expectedResultX {
			t.Errorf("TestCommand_Apply X %s: expected %v, got %v", c.name, c.expectedResultX, p.X)
		}

		if p.Y != c.expectedResultY {
			t.Errorf("TestCommand_Apply Y %s: expected %v, got %v", c.name, c.expectedResultY, p.Y)
		}
	}
}
