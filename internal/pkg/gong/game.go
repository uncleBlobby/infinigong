package gong

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
)

const DELTA_TIME = 0.016
const GAME_WIDTH = 800
const GAME_HEIGHT = 600

type Game struct {
	Bricks        []Brick
	Players       []Player
	UserInterface UserInterface
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, brick := range g.Bricks {
		brick.Draw(screen)
	}

	for _, player := range g.Players {
		player.Draw(screen)
	}

	for _, button := range g.UserInterface.Buttons {
		button.Draw(screen)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()), 250, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player 1 Score: %d", g.Players[0].Score), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player 2 Score: %d", g.Players[1].Score), 600, 10)
}

func NewGame() *Game {

	g := &Game{}
	g.InitializePlayers()
	g.InitializeBricks()

	// debug button test
	g.InitializeButton(50, 50, g.onClickWrapper)

	return g
}

func (g *Game) onClickWrapper() {
	g.Players[0].Paddle.IncreasePaddleSize(1)
	fmt.Println("hello, on click function!")
}

func (g *Game) InitializeButton(length, width float32, ocf func()) {
	b := NewButton(50, 50, fig.NewVector2(10, 100), "test button", ocf)
	g.UserInterface.Buttons = append(g.UserInterface.Buttons, *b)
}

func (g *Game) InitializePlayers() {
	p1 := NewPlayer(PLAYER_ONE_COLOUR, PLAYER_ONE_BALL_START_POS, PLAYER_ONE_BALL_START_VEL, PLAYER_ONE_PADDLE_START_POS)
	g.Players = append(g.Players, p1)
	p2 := NewPlayer(PLAYER_TWO_COLOUR, PLAYER_TWO_BALL_START_POS, PLAYER_TWO_BALL_START_VEL, PLAYER_TWO_PADDLE_START_POS)
	g.Players = append(g.Players, p2)
}

func (g *Game) InitializeBricks() {
	for x := 0; x < GAME_WIDTH/2; x += BRICK_WIDTH {
		for y := 0; y < GAME_HEIGHT; y += BRICK_HEIGHT {
			g.Bricks = append(g.Bricks, NewBrick(PLAYER_TWO_COLOUR, fig.NewVector2(float32(x), float32(y))))
		}
	}

	for x := GAME_WIDTH / 2; x < GAME_WIDTH; x += BRICK_WIDTH {
		for y := 0; y < GAME_HEIGHT; y += BRICK_HEIGHT {
			g.Bricks = append(g.Bricks, NewBrick(PLAYER_ONE_COLOUR, fig.NewVector2(float32(x), float32(y))))
		}
	}
}

func (g *Game) Update() error {

	g.CheckCollisions()

	for i := 0; i < len(g.Players); i++ {
		g.Players[i].Update(DELTA_TIME)
	}

	for i := 0; i < len(g.UserInterface.Buttons); i++ {
		g.UserInterface.Buttons[i].Update(DELTA_TIME)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) CheckCollisions() {
	for j := 0; j < len(g.Players); j++ {
		if g.Players[j].Ball.Circle.Position.X-g.Players[j].Ball.Circle.Radius < 0 || g.Players[j].Ball.Circle.Position.X+g.Players[j].Ball.Circle.Radius > 800 {
			g.Players[j].Ball.Velocity.X *= -1
		}
		if g.Players[j].Ball.Circle.Position.Y-g.Players[j].Ball.Circle.Radius < 0 || g.Players[j].Ball.Circle.Position.Y+g.Players[j].Ball.Circle.Radius > 600 {
			g.Players[j].Ball.Velocity.Y *= -1
		}

		if g.Players[j].Ball.CollidesWithPaddle(*g.Players[j].Paddle) {
			g.Players[j].Ball.Velocity.X *= -1
		}

		for i := 0; i < len(g.Bricks); i++ {
			if g.Players[j].Ball.CollidesWithBrick(g.Bricks[i]) {
				if g.Bricks[i].Color == g.Players[j].Ball.Color && g.Players[j].Ball.Color == PLAYER_ONE_COLOUR {
					g.Players[j].Ball.Velocity.X *= -1
					g.Bricks[i].Color = PLAYER_TWO_COLOUR
					g.Players[j].IncrementScore(1)
					// increment ball speed with score just for fun :P
					g.Players[j].Ball.Speed *= 1.01
				}
				if g.Bricks[i].Color == g.Players[j].Ball.Color && g.Players[j].Ball.Color == PLAYER_TWO_COLOUR {
					g.Players[j].Ball.Velocity.X *= -1
					g.Bricks[i].Color = PLAYER_ONE_COLOUR
					g.Players[j].IncrementScore(1)
					// increment ball speed with score just for fun :P
					g.Players[j].Ball.Speed *= 1.01
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GAME_WIDTH, GAME_HEIGHT
}
