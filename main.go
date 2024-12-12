package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Bricks  []Brick
	Players []Player
}

var PLAYER_ONE_COLOUR = color.RGBA{0xff, 0xff, 0x00, 0xff}
var PLAYER_TWO_COLOUR = color.RGBA{0xff, 0x00, 0xff, 0xff}

const BRICK_WIDTH = 20
const BRICK_HEIGHT = 40

const PADDLE_LENGTH = 80
const PADDLE_WIDTH = 10
const PADDLE_SPEED = 100

const BALL_SIZE = 5

const DELTA_TIME = 0.016

func (g *Game) Update() error {

	g.CheckCollisions()

	//log.Printf("calling update function")
	for _, player := range g.Players {
		player.Ball.Update(DELTA_TIME)
		//log.Printf("ball Position: %v", player.Ball.Circle.Position)
		player.Paddle.Update(DELTA_TIME)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

type Vector2 struct {
	X, Y float32
}

type Rectangle struct {
	Length   float32
	Width    float32
	Position Vector2
}

type Circle struct {
	Radius   float32
	Position Vector2
}

type Brick struct {
	Color color.RGBA
	Rect  Rectangle
}

type Paddle struct {
	Color    color.RGBA
	Rect     Rectangle
	Velocity Vector2
	Speed    float32
}

type Ball struct {
	Color    color.RGBA
	Circle   Circle
	Velocity Vector2
}

type Player struct {
	Ball   *Ball
	Paddle *Paddle
	Score  int
}

func (p *Player) IncrementScore(X int) {
	p.Score += X
}

func (b *Ball) CheckBrickCollisions(g *Game) {

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

func (g *Game) CheckCollisions() {
	// COLLISIONS WITH BRICKS
	// find middle position of rectangle

	// find middle position of circle

	// if the distance between those two points is less than the radius of the circle, they are colliding

	// COLLISIONS WITH WALLS

	for _, player := range g.Players {
		if player.Ball.Circle.Position.X-player.Ball.Circle.Radius < 0 || player.Ball.Circle.Position.X+player.Ball.Circle.Radius > 800 {
			player.Ball.Velocity.X *= -1
		}
		if player.Ball.Circle.Position.Y-player.Ball.Circle.Radius < 0 || player.Ball.Circle.Position.Y+player.Ball.Circle.Radius > 600 {
			player.Ball.Velocity.Y *= -1
		}

		if player.Ball.CollidesWithPaddle(*player.Paddle) {
			player.Ball.Velocity.X *= -1
		}

		// for i := 0; i < len(g.Bricks); i++ {
		// 	if math.Abs(float64(g.Bricks[i].GetMiddlePosition().X-player.Ball.Circle.Position.X)) < float64(player.Ball.Circle.Radius) && math.Abs(float64(g.Bricks[i].GetMiddlePosition().Y-player.Ball.Circle.Position.Y)) < float64(player.Ball.Circle.Radius) {
		// 		if g.Bricks[i].Color == player.Ball.Color && player.Ball.Color == PLAYER_ONE_COLOUR {
		// 			player.Ball.Velocity.X *= -1
		// 			g.Bricks[i].Color = PLAYER_TWO_COLOUR
		// 		}
		// 		if g.Bricks[i].Color == player.Ball.Color && player.Ball.Color == PLAYER_TWO_COLOUR {
		// 			player.Ball.Velocity.X *= -1
		// 			g.Bricks[i].Color = PLAYER_ONE_COLOUR
		// 		}
		// 	}
		// }

		for i := 0; i < len(g.Bricks); i++ {
			if player.Ball.CollidesWithBrick(g.Bricks[i]) {
				if g.Bricks[i].Color == player.Ball.Color && player.Ball.Color == PLAYER_ONE_COLOUR {
					player.Ball.Velocity.X *= -1
					g.Bricks[i].Color = PLAYER_TWO_COLOUR
					player.IncrementScore(1)
				}
				if g.Bricks[i].Color == player.Ball.Color && player.Ball.Color == PLAYER_TWO_COLOUR {
					player.Ball.Velocity.X *= -1
					g.Bricks[i].Color = PLAYER_ONE_COLOUR
					player.IncrementScore(1)
				}
			}
		}
	}

}

func (b *Brick) GetMiddlePosition() Vector2 {
	return Vector2{
		X: b.Rect.Position.X + b.Rect.Width/2,
		Y: b.Rect.Position.Y + b.Rect.Length/2,
	}
}

func (b *Brick) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.Rect.Position.X), float32(b.Rect.Position.Y), b.Rect.Width, b.Rect.Length, b.Color, true)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.Circle.Position.X, b.Circle.Position.Y, b.Circle.Radius, b.Color, true)

	// draw collider

	//vector.StrokeCircle(screen, b.Circle.Position.X, b.Circle.Position.Y, b.Circle.Radius, 2, color.White, true)
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.Rect.Position.X), float32(p.Rect.Position.Y), p.Rect.Width, p.Rect.Length, p.Color, true)
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

func (b *Ball) Update(dt float32) {
	b.Circle.Position.X += b.Velocity.X * dt
	b.Circle.Position.Y += b.Velocity.Y * dt
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Ball.Draw(screen)
	p.Paddle.Draw(screen)
}

func (g *Game) Initialize() {

	p1 := Player{
		Score: 0,
		Ball: &Ball{
			Color: PLAYER_ONE_COLOUR,
			Circle: Circle{
				Position: Vector2{
					X: 400 / 2,
					Y: 600 / 2,
				},
				Radius: BALL_SIZE,
			},
			Velocity: Vector2{
				X: 100,
				Y: 100,
			},
		},
		Paddle: &Paddle{
			Speed: PADDLE_SPEED,
			Color: PLAYER_ONE_COLOUR,
			Rect: Rectangle{
				Length: PADDLE_LENGTH,
				Width:  PADDLE_WIDTH,
				Position: Vector2{
					X: 20,
					Y: 600/2 - PADDLE_LENGTH/2,
				},
			},
		},
	}

	g.Players = append(g.Players, p1)

	p2 := Player{
		Score: 0,
		Ball: &Ball{
			Color: PLAYER_TWO_COLOUR,
			Circle: Circle{
				Position: Vector2{
					X: 400/2 + 400,
					Y: 600 / 2,
				},
				Radius: BALL_SIZE,
			},
			Velocity: Vector2{
				X: -100,
				Y: -100,
			},
		},
		Paddle: &Paddle{
			Speed: PADDLE_SPEED,
			Color: PLAYER_TWO_COLOUR,
			Rect: Rectangle{
				Length: PADDLE_LENGTH,
				Width:  PADDLE_WIDTH,
				Position: Vector2{
					X: 800 - 40,
					Y: 600/2 - PADDLE_LENGTH/2,
				},
			},
		},
	}

	g.Players = append(g.Players, p2)

	for x := 0; x < 800/2; x += 20 {
		for y := 0; y < 600; y += 40 {
			b := &Brick{
				Color: PLAYER_TWO_COLOUR,
				Rect: Rectangle{
					Length: BRICK_HEIGHT,
					Width:  BRICK_WIDTH,
					Position: Vector2{
						X: float32(x),
						Y: float32(y),
					},
				},
			}
			g.Bricks = append(g.Bricks, *b)
			//vector.DrawFilledRect(screen, float32(x), float32(y), 20, 40, color.RGBA{0xff, 0xff, 0x00, 0xff}, true)
		}
	}

	for x := 400; x < 800; x += 20 {
		for y := 0; y < 600; y += 40 {
			b := &Brick{
				Color: PLAYER_ONE_COLOUR,
				Rect: Rectangle{
					Length: BRICK_HEIGHT,
					Width:  BRICK_WIDTH,
					Position: Vector2{
						X: float32(x),
						Y: float32(y),
					},
				},
			}
			g.Bricks = append(g.Bricks, *b)
			//vector.DrawFilledRect(screen, float32(x), float32(y), 20, 40, color.RGBA{0xff, 0x00, 0xff, 0xff}, true)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for _, brick := range g.Bricks {
		brick.Draw(screen)
	}

	for _, player := range g.Players {
		player.Draw(screen)

		// draw collider

		// vector.StrokeCircle(screen, player.Ball.Circle.Position.X, player.Ball.Circle.Position.Y, player.Ball.Circle.Radius, 2, color.White, true)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player 1 Score: %d", g.Players[0].Score), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player 2 Score: %d", g.Players[1].Score), 600, 10)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}
	g.Initialize()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
