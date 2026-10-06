package player

type Terrain interface {
	WalkableAt(x, y float64) bool
}

func canStandAt(t Terrain, x, y float64) bool {
	const half = Size / 2

	return t.WalkableAt(x-half, y-half) &&
		t.WalkableAt(x+half, y-half) &&
		t.WalkableAt(x-half, y+half) &&
		t.WalkableAt(x+half, y+half)
}
