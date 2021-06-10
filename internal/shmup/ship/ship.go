package ship

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/serg-volodeev/shmup/internal/shape"
	"github.com/serg-volodeev/shmup/internal/shmup/bullet"
	"github.com/serg-volodeev/shmup/internal/shmup/meteor"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
)

type Opts struct {
	Bounds  *shape.Rect
	Meteors *meteor.Meteors
	Bullets *bullet.Bullets
	Res     *res.Res
}

type Ship struct {
	bounds  *shape.Rect
	meteors *meteor.Meteors
	bullets *bullet.Bullets
	rect    *shape.Rect
	circle  *shape.Circle
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedX  float64
}

func NewShip(o *Opts) *Ship {
	s := &Ship{}
	s.bounds = o.Bounds
	s.meteors = o.Meteors
	s.bullets = o.Bullets
	s.image = o.Res.GetImage("ship")
	s.options = &ebiten.DrawImageOptions{}
	s.speedX = 4

	s.rect = shape.NewRectFromImage(s.image)
	s.rect.SetCenterX(s.bounds.CenterX())
	s.rect.SetBottom(s.bounds.Bottom() - 10)

	s.circle = shape.NewCircle(s.rect.CenterX(), s.rect.CenterY(), s.rect.Height()/2-2)
	return s
}

func (s *Ship) Draw(screen *ebiten.Image) {
	s.options.GeoM.Reset()
	s.options.GeoM.Translate(s.rect.Left(), s.rect.Top())
	screen.DrawImage(s.image, s.options)

	// x := s.rect.centerX() - s.radius
	// y := s.rect.centerY() - s.radius
	// ebitenutil.DrawRect(screen, x, y, s.radius*2, s.radius*2, color.White)
}

func (s *Ship) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.move(s.speedX)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.move(-s.speedX)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fire()
	}
	if s.meteors.CollideCircle(s.circle) {
		return fmt.Errorf("collide meteor")
	}
	return nil
}

func (s *Ship) move(dx float64) {
	s.rect.MoveX(dx)
	if s.rect.Right() > s.bounds.Right() {
		s.rect.SetRight(s.bounds.Right())
	}
	if s.rect.Left() < s.bounds.Left() {
		s.rect.SetLeft(s.bounds.Left())
	}
	s.circle.SetCenter(s.rect.CenterX(), s.rect.CenterY())
}

func (s *Ship) fire() {
	s.bullets.NewBullet(s.rect)
	// b.SetCenterX(s.rect.CenterX())
	// b.SetBottom(s.rect.Top())
}
