package gong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
	"github.com/uncleBlobby/infinigong/internal/pkg/gfx"
)

const PADDLE_LENGTH = 80
const PADDLE_WIDTH = 10
const PADDLE_SPEED = 400

type Paddle struct {
	Color    color.RGBA
	Rect     gfx.Rectangle
	Velocity fig.Vector2
	Speed    float32
}

func NewPaddle(color color.RGBA, startPos fig.Vector2) *Paddle {
	return &Paddle{
		Speed: PADDLE_SPEED,
		Color: color,
		Rect:  gfx.NewRectangle(PADDLE_LENGTH, PADDLE_WIDTH, startPos),
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.Rect.Position.X), float32(p.Rect.Position.Y), p.Rect.Width, p.Rect.Length, p.Color, true)
}

func (p *Paddle) IncreasePaddleSize(X float32) {
	p.Rect.Length += X
	p.Rect.Width += X
}

func (p *Paddle) Update(dt float32) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if p.Color == PLAYER_ONE_COLOUR {
			p.Velocity.Y = -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.Color == PLAYER_TWO_COLOUR {
			p.Velocity.Y = -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if p.Color == PLAYER_ONE_COLOUR {
			p.Velocity.Y = 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.Color == PLAYER_TWO_COLOUR {
			p.Velocity.Y = 1
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		if p.Color == PLAYER_ONE_COLOUR {
			p.Velocity.Y = 0
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.Color == PLAYER_TWO_COLOUR {
			p.Velocity.Y = 0
		}
	}

	p.Rect.Position.X += p.Velocity.X * p.Speed * dt
	p.Rect.Position.Y += p.Velocity.Y * p.Speed * dt

}
