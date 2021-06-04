package shmup

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Ship struct {
	rect    *shape.Rect
	circle  *shape.Circle
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedX  float64
}

func newShip(g *Game) *Ship {
	s := &Ship{}
	s.image = g.res.images["ship"]
	s.options = &ebiten.DrawImageOptions{}
	s.speedX = 4

	s.rect = shape.NewRectFromImage(s.image)
	s.rect.SetCenterX(g.rect.CenterX())
	s.rect.SetBottom(g.rect.Bottom() - 10)

	s.circle = shape.NewCircle(s.rect.CenterX(), s.rect.CenterY(), s.rect.Height()/2-2)
	return s
}

func (s *Ship) draw(screen *ebiten.Image) {
	s.options.GeoM.Reset()
	s.options.GeoM.Translate(s.rect.Left(), s.rect.Top())
	screen.DrawImage(s.image, s.options)

	// x := s.rect.centerX() - s.radius
	// y := s.rect.centerY() - s.radius
	// ebitenutil.DrawRect(screen, x, y, s.radius*2, s.radius*2, color.White)
}

func (s *Ship) update(g *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.move(s.speedX, g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.move(-s.speedX, g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fire(g)
	}
	if g.meteors.collideCircle(s.circle) {
		return fmt.Errorf("collide meteor")
	}
	return nil
}

func (s *Ship) move(dx float64, g *Game) {
	s.rect.MoveX(dx)
	if s.rect.Right() > g.rect.Right() {
		s.rect.SetRight(g.rect.Right())
	}
	if s.rect.Left() < g.rect.Left() {
		s.rect.SetLeft(g.rect.Left())
	}
	s.circle.SetCenter(s.rect.CenterX(), s.rect.CenterY())
}

func (s *Ship) fire(g *Game) {
	b := g.bullets.newBullet()
	b.rect.SetCenterX(s.rect.CenterX())
	b.rect.SetBottom(s.rect.Top())
}
