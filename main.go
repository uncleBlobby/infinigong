package main

import (
	"image/color"
	"log"

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

func (g *Game) Update() error {
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

type Ball struct {
	Color  color.RGBA
	Circle Circle
}

type Player struct {
	Ball Ball
}

func (b *Brick) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.Rect.Position.X), float32(b.Rect.Position.Y), b.Rect.Width, b.Rect.Length, b.Color, true)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.Circle.Position.X, b.Circle.Position.Y, b.Circle.Radius, b.Color, true)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Ball.Draw(screen)
}

func (g *Game) Initialize() {

	p1 := Player{
		Ball: Ball{
			Color: PLAYER_ONE_COLOUR,
			Circle: Circle{
				Position: Vector2{
					X: 400 / 2,
					Y: 600 / 2,
				},
				Radius: 10,
			},
		},
	}

	g.Players = append(g.Players, p1)

	p2 := Player{
		Ball: Ball{
			Color: PLAYER_TWO_COLOUR,
			Circle: Circle{
				Position: Vector2{
					X: 400/2 + 400,
					Y: 600 / 2,
				},
				Radius: 10,
			},
		},
	}

	g.Players = append(g.Players, p2)

	for x := 0; x < 800/2; x += 20 {
		for y := 0; y < 600; y += 40 {
			b := Brick{
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
			g.Bricks = append(g.Bricks, b)
			//vector.DrawFilledRect(screen, float32(x), float32(y), 20, 40, color.RGBA{0xff, 0xff, 0x00, 0xff}, true)
		}
	}

	for x := 400; x < 800; x += 20 {
		for y := 0; y < 600; y += 40 {
			b := Brick{
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
			g.Bricks = append(g.Bricks, b)
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
	}

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
