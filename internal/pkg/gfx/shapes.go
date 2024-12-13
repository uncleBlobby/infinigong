package gfx

import "github.com/uncleBlobby/infinigong/internal/pkg/fig"

type Rectangle struct {
	Length   float32
	Width    float32
	Position fig.Vector2
}

func NewRectangle(l, w float32, position fig.Vector2) Rectangle {
	return Rectangle{
		Length:   l,
		Width:    w,
		Position: position,
	}
}

func (r *Rectangle) GetMiddlePosition() fig.Vector2 {
	return fig.Vector2{
		X: r.Position.X + r.Width/2,
		Y: r.Position.Y + r.Length/2,
	}
}

type Circle struct {
	Radius   float32
	Position fig.Vector2
}

func NewCircle(radius float32, position fig.Vector2) Circle {
	return Circle{
		Radius:   radius,
		Position: position,
	}
}
