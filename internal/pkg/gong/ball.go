package gong

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
	"github.com/uncleBlobby/infinigong/internal/pkg/gfx"
)

const BALL_SIZE = 5
const BALL_SPEED = 2

type Ball struct {
	Color    color.RGBA
	Circle   gfx.Circle
	Velocity fig.Vector2
	Speed    float32
}

func NewBall(color color.RGBA, startPos fig.Vector2, startVel fig.Vector2) *Ball {
	return &Ball{
		Speed:    BALL_SPEED,
		Color:    color,
		Circle:   gfx.NewCircle(BALL_SIZE, startPos),
		Velocity: startVel,
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.Circle.Position.X, b.Circle.Position.Y, b.Circle.Radius, b.Color, true)
}

func (b *Ball) Update(dt float32) {
	b.Circle.Position.X += b.Velocity.X * dt * b.Speed
	b.Circle.Position.Y += b.Velocity.Y * dt * b.Speed

	if b.Circle.Radius > BALL_SIZE {
		b.Circle.Radius = b.Circle.Radius * 0.9995
	}

	if b.Circle.Radius > 50 {
		b.Circle.Radius = 50
	}

	if b.Speed > 7 {
		b.Speed = 7
	}
}

func (b *Ball) CollidesWithPaddle(paddle Paddle) bool {
	testX := b.Circle.Position.X
	testY := b.Circle.Position.Y

	//if ball is to the left of the brick
	if b.Circle.Position.X < paddle.Rect.Position.X {
		// test against the left edge of the paddle
		testX = paddle.Rect.Position.X
		// if ball is to the right of the paddle
	} else if b.Circle.Position.X > paddle.Rect.Position.X+paddle.Rect.Width {
		// test against the right edge of the paddle
		testX = paddle.Rect.Position.X + paddle.Rect.Width
	}

	//if ball is above the paddle
	if b.Circle.Position.Y < paddle.Rect.Position.Y {
		// test agains the top edge of the paddle
		testY = paddle.Rect.Position.Y
		// if ball is underneath the paddle
	} else if b.Circle.Position.Y > paddle.Rect.Position.Y+paddle.Rect.Length {
		// test against bottom edge of the paddle
		testY = paddle.Rect.Position.Y + paddle.Rect.Length
	}

	distX := math.Abs(float64(b.Circle.Position.X - testX))
	distY := math.Abs(float64(b.Circle.Position.Y - testY))
	distance := math.Sqrt((distX * distX) + (distY * distY))

	return distance <= float64(b.Circle.Radius)
}

func (b *Ball) CollidesWithBrick(brick Brick) bool {
	testX := b.Circle.Position.X
	testY := b.Circle.Position.Y

	//if ball is to the left of the brick
	if b.Circle.Position.X < brick.Rect.Position.X {
		// test against the left edge of the brick
		testX = brick.Rect.Position.X
		// if ball is to the right of the brick
	} else if b.Circle.Position.X > brick.Rect.Position.X+brick.Rect.Width {
		// test against the right edge of the brick
		testX = brick.Rect.Position.X + brick.Rect.Width
	}

	//if ball is above the brick
	if b.Circle.Position.Y < brick.Rect.Position.Y {
		// test agains the top edge of the brick
		testY = brick.Rect.Position.Y
		// if ball is underneath the brick
	} else if b.Circle.Position.Y > brick.Rect.Position.Y+brick.Rect.Length {
		// test against bottom edge of the brick
		testY = brick.Rect.Position.Y + brick.Rect.Length
	}

	distX := math.Abs(float64(b.Circle.Position.X - testX))
	distY := math.Abs(float64(b.Circle.Position.Y - testY))
	distance := math.Sqrt((distX * distX) + (distY * distY))

	return distance <= float64(b.Circle.Radius)
}
