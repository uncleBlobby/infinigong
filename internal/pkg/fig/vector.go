package fig

type Vector2 struct {
	X, Y float32
}

func NewVector2(x, y float32) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}
