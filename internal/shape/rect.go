package shape

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Rect struct {
	x, y float64
	w, h float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{x: x, y: y, w: w, h: h}
}

func NewRectFromImage(image *ebiten.Image) *Rect {
	return NewRect(0, 0, float64(image.Bounds().Dx()), float64(image.Bounds().Dy()))
}

func (r *Rect) Height() float64 {
	return r.h
}

func (r *Rect) Width() float64 {
	return r.w
}

func (r *Rect) CenterX() float64 {
	return r.x + r.w/2
}

func (r *Rect) SetCenterX(x float64) {
	r.x = x - r.w/2
}

func (r *Rect) CenterY() float64 {
	return r.y + r.h/2
}

func (r *Rect) SetCenterY(y float64) {
	r.y = y - r.h/2
}

func (r *Rect) Left() float64 {
	return r.x
}

func (r *Rect) SetLeft(x float64) {
	r.x = x
}

func (r *Rect) Right() float64 {
	return r.x + r.w
}

func (r *Rect) SetRight(x float64) {
	r.x = x - r.w
}

func (r *Rect) Top() float64 {
	return r.y
}

func (r *Rect) SetTop(y float64) {
	r.y = y
}

func (r *Rect) Bottom() float64 {
	return r.y + r.h
}

func (r *Rect) SetBottom(y float64) {
	r.y = y - r.h
}

func (r *Rect) MoveX(dx float64) {
	r.x += dx
}

func (r *Rect) MoveY(dy float64) {
	r.y += dy
}
