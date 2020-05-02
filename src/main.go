package main

import (
	"SomeTest/pixel/pixelgl"
	/*"math"*/
	"github.com/faiface/pixel"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	ScreenW, ScreenH := pixelgl.Monitors()[0].Size()
	/*Unit := math.Sqrt(ScreenH * ScreenW) / 200*/

	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, ScreenW*0.7, ScreenH*0.6),
		Title:  "Gotanks"}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		win.Update()
	}
}

type Tank struct {
	X                     float64
	Y                     float64
	ReloadTime            float64
	CurrentReloadCoundown float64
	Speed                 float64
	Angle                 float64
	Buttons               [5]pixelgl.Button
	color                 [3]int8
	id                    int
}

func (t *Tank) move(vec pixel.Vec) {
	t.X = vec.X
	t.Y = vec.Y
}
func (t *Tank) rotate(angle float64) {
	t.Angle = angle
}
func (t *Tank) draw() {

}
func (t *Tank) update() {

}
func (t *Tank) shoot() {

}

type Bullet struct {
	X         float64
	Y         float64
	angle     float64
	direction pixel.Vec
	speed     float64
}

func (b *Bullet) update() {
	b.X += b.direction.X * b.speed
	b.Y += b.direction.Y * b.speed
}
func (b *Bullet) draw() {

}

type Controllable interface {
	move(pixel.Vec)
	rotate(float64)
}
type Drawable interface {
	draw()
}
type Updatable interface {
	update()
}
type Shootable interface {
	shoot()
}
