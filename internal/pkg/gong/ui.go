package gong

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/uncleBlobby/infinigong/internal/pkg/fig"
	"github.com/uncleBlobby/infinigong/internal/pkg/gfx"
)

type UserInterface struct {
	Buttons []Button
}

type Button struct {
	Uuid       string
	Rect       gfx.Rectangle
	Text       string
	Color      color.RGBA
	HoverColor color.RGBA
	IsHovered  bool
	onClickFn  func()
}

func (b *Button) Update(dt float32) {
	mouseX, mouseY := ebiten.CursorPosition()

	if mouseX > int(b.Rect.Position.X) && mouseX < int(b.Rect.Position.X)+int(b.Rect.Length) && mouseY > int(b.Rect.Position.Y) && mouseY < int(b.Rect.Position.Y)+int(b.Rect.Width) {
		b.IsHovered = true
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			b.OnClick()
		}
	} else {
		b.IsHovered = false
	}

}

func (b *Button) OnClick() {
	//fmt.Println("clicked the button")
	fmt.Println("calling the on click function!")
	b.onClickFn()
}

func (b *Button) Draw(screen *ebiten.Image) {
	if !b.IsHovered {
		vector.DrawFilledRect(screen, b.Rect.Position.X, b.Rect.Position.Y, b.Rect.Length, b.Rect.Width, b.Color, true)
	}
	if b.IsHovered {
		vector.DrawFilledRect(screen, b.Rect.Position.X, b.Rect.Position.Y, b.Rect.Length, b.Rect.Width, b.HoverColor, true)
	}

	ebitenutil.DebugPrintAt(screen, b.Text, int(b.Rect.Position.X)+int(b.Rect.Length)/2, int(b.Rect.Position.Y)+int(b.Rect.Width)/2)
}

func NewButton(length, width float32, startPos fig.Vector2, text string, ocf func()) *Button {
	return &Button{
		Rect:       gfx.NewRectangle(length, width, startPos),
		Text:       text,
		Color:      color.RGBA{0xff, 0x00, 0x00, 0xff},
		HoverColor: color.RGBA{0x00, 0xff, 0x00, 0xff},
		IsHovered:  false,
		onClickFn:  ocf,
	}
}
