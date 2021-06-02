package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Rect struct {
	x, y float64
	w, h float64
}

func newRect(x, y, w, h float64) *Rect {
	return &Rect{x: x, y: y, w: w, h: h}
}

func newRectFromImage(image *ebiten.Image) *Rect {
	return newRect(0, 0, float64(image.Bounds().Dx()), float64(image.Bounds().Dy()))
}

func (r *Rect) centerX() float64 {
	return r.x + r.w/2
}

func (r *Rect) setCenterX(x float64) {
	r.x = x - r.w/2
}

func (r *Rect) centerY() float64 {
	return r.y + r.h/2
}

func (r *Rect) setCenterY(y float64) {
	r.y = y - r.h/2
}

func (r *Rect) left() float64 {
	return r.x
}

func (r *Rect) setLeft(x float64) {
	r.x = x
}

func (r *Rect) right() float64 {
	return r.x + r.w
}

func (r *Rect) setRight(x float64) {
	r.x = x - r.w
}

func (r *Rect) top() float64 {
	return r.y
}

func (r *Rect) setTop(y float64) {
	r.y = y
}

func (r *Rect) bottom() float64 {
	return r.y + r.h
}

func (r *Rect) setBottom(y float64) {
	r.y = y - r.h
}

func (r *Rect) moveX(dx float64) {
	r.x += dx
}

func (r *Rect) moveY(dy float64) {
	r.y += dy
}
