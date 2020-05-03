package main

import (
<<<<<<< HEAD
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math"
)

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
=======
	"SomeTest/pixel/pixelgl"
	/*"math"*/
	"github.com/faiface/pixel"
)

>>>>>>> origin/pixel
func main() {
	pixelgl.Run(run)
}

func run() {
	ScreenW, ScreenH := pixelgl.Monitors()[0].Size()
<<<<<<< HEAD
	Unit := math.Sqrt(ScreenH*ScreenW) / 200

	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, ScreenW*0.7, ScreenH*0.6),
		Title:  "Gotanks",
		VSync:  true}
=======
	/*Unit := math.Sqrt(ScreenH * ScreenW) / 200*/

	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, ScreenW*0.7, ScreenH*0.6),
		Title:  "Gotanks"}
>>>>>>> origin/pixel
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}
<<<<<<< HEAD

	t := Tank{
		X:            500,
		Y:            500,
		color:        color.RGBA{A: 0xff, B: 0x00, R: 0xff, G: 0x00},
		id:           0,
		Angle:        90.0,
		Radius:       2.5 * Unit,
		Buttons:      [5]pixelgl.Button{pixelgl.KeyW, pixelgl.KeyS, pixelgl.KeyA, pixelgl.KeyD, pixelgl.KeySpace},
		Speed:        0.5 * Unit,
		AngularSpeed: 2.0,
	}
	for !win.Closed() {
		win.Clear(color.RGBA{A: 0xff, R: 0xc8, G: 0xc8, B: 0xc8})
		t.update(win)
		t.draw(win, Unit)
=======
	for !win.Closed() {
>>>>>>> origin/pixel
		win.Update()
	}
}

type Tank struct {
	X                     float64
	Y                     float64
	ReloadTime            float64
	CurrentReloadCoundown float64
<<<<<<< HEAD
	ReloadSpeed           float64
	Speed                 float64
	AngularSpeed          float64
	Buttons               [5]pixelgl.Button
	color                 color.RGBA
	Angle                 float64
	Radius                float64
=======
	Speed                 float64
	Angle                 float64
	Buttons               [5]pixelgl.Button
	color                 [3]int8
>>>>>>> origin/pixel
	id                    int
}

func (t *Tank) move(vec pixel.Vec) {
<<<<<<< HEAD
	t.X += vec.X
	t.Y += vec.Y
}

func (t *Tank) draw(target pixel.Target, Unit float64) {
	im := imdraw.New(nil)
	im.Color = t.color
	im.Push(pixel.V(t.X, t.Y))
	im.Circle(t.Radius, 0)
	im.Color = pixel.RGBA{A: 0xff}
	im.Push(pixel.V(t.X, t.Y))
	im.CircleArc(t.Radius, Radians(t.Angle)-math.Pi/6, Radians(t.Angle)+math.Pi/6, 0)

	im.Draw(target)

}

func (t *Tank) update(window *pixelgl.Window) {

	if window.Pressed(t.Buttons[2]) {
		t.Angle += t.AngularSpeed
	}
	if window.Pressed(t.Buttons[3]) {
		t.Angle -= t.AngularSpeed
	}
	if window.Pressed(t.Buttons[0]) {
		t.move(pixel.V(1, 0).Rotated(Radians(t.Angle)).Scaled(t.Speed))
	}
	if window.Pressed(t.Buttons[1]) {
		t.move(pixel.V(1, 0).Rotated(Radians(t.Angle)).Scaled(-t.Speed))
	}
	if window.Pressed(t.Buttons[4]) {
		if t.CurrentReloadCoundown <= 0 {
			t.shoot()
			t.CurrentReloadCoundown = t.ReloadTime
		}
	}
	if t.CurrentReloadCoundown > 0 {
		t.CurrentReloadCoundown -= t.ReloadSpeed
	}
}

=======
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
>>>>>>> origin/pixel
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
<<<<<<< HEAD

=======
>>>>>>> origin/pixel
func (b *Bullet) draw() {

}

type Controllable interface {
	move(pixel.Vec)
<<<<<<< HEAD
}
type Drawable interface {
	draw(pixel.Target, float64)
=======
	rotate(float64)
}
type Drawable interface {
	draw()
>>>>>>> origin/pixel
}
type Updatable interface {
	update()
}
type Shootable interface {
	shoot()
}
