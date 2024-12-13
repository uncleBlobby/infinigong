package gong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
	"github.com/uncleBlobby/infinigong/internal/pkg/gfx"
)

const BRICK_WIDTH = 20
const BRICK_HEIGHT = 40

type Brick struct {
	Color color.RGBA
	Rect  gfx.Rectangle
}

func NewBrick(color color.RGBA, position fig.Vector2) Brick {
	return Brick{
		Color: color,
		Rect:  gfx.NewRectangle(BRICK_HEIGHT, BRICK_WIDTH, position),
	}
}

func (b *Brick) GetMiddlePosition() fig.Vector2 {
	return b.Rect.GetMiddlePosition()
}

func (b *Brick) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.Rect.Position.X), float32(b.Rect.Position.Y), b.Rect.Width, b.Rect.Length, b.Color, true)
}
