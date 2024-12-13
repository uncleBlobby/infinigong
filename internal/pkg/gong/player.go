package gong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
)

var PLAYER_ONE_COLOUR = color.RGBA{0xff, 0xff, 0x00, 0xff}
var PLAYER_TWO_COLOUR = color.RGBA{0xff, 0x00, 0xff, 0xff}

var PLAYER_ONE_PADDLE_START_POS = fig.Vector2{X: 20, Y: 600/2 - PADDLE_LENGTH/2}
var PLAYER_TWO_PADDLE_START_POS = fig.Vector2{X: 800 - 40, Y: 600/2 - PADDLE_LENGTH/2}

var PLAYER_ONE_BALL_START_POS = fig.Vector2{X: 400 / 2, Y: 600 / 2}
var PLAYER_TWO_BALL_START_POS = fig.Vector2{X: 400/2 + 400, Y: 600 / 2}

var PLAYER_ONE_BALL_START_VEL = fig.Vector2{X: 100, Y: 100}
var PLAYER_TWO_BALL_START_VEL = fig.Vector2{X: -100, Y: -100}

type Player struct {
	Ball   *Ball
	Paddle *Paddle
	Score  int
}

func (p *Player) Update(dt float32) {
	p.Ball.Update(dt)
	p.Paddle.Update(dt)
}

func NewPlayer(color color.RGBA, ballPos, ballVel, paddlePos fig.Vector2) Player {
	return Player{
		Score:  0,
		Ball:   NewBall(color, ballPos, ballVel),
		Paddle: NewPaddle(color, paddlePos),
	}
}

func (p *Player) IncrementScore(X int) {
	p.Score += X
	// Janky Feature:
	// Increase player ball radius with score
	p.Ball.Circle.Radius = BALL_SIZE + float32(p.Score)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Ball.Draw(screen)
	p.Paddle.Draw(screen)
}
