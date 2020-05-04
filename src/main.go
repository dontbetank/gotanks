package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"os"
	"time"
)

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
	pixelgl.Run(run)
}

func run() {
	ScreenW, ScreenH := pixelgl.Monitors()[0].Size()
	Unit := math.Sqrt(ScreenH*ScreenW) / 200
	pic, err := loadPicture("assets/images/tank.png")
	if err != nil {
		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, ScreenW*0.7, ScreenH*0.6),
		Title:  "Gotanks",
		VSync:  true}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tankFrames := []pixel.Rect{}
	tankFrames = append(tankFrames, pixel.R(0, 0, 64, 32))
	tankFrames = append(tankFrames, pixel.R(0, 32, 64, 64))

	t := Tank{
		X:            500,
		Y:            500,
		color:        color.RGBA{A: 0xff, B: 0x00, R: 0xff, G: 0x00},
		id:           0,
		Angle:        0,
		Radius:       2.5 * Unit,
		Buttons:      [5]pixelgl.Button{pixelgl.KeyW, pixelgl.KeyS, pixelgl.KeyA, pixelgl.KeyD, pixelgl.KeySpace},
		Speed:        0.5 * Unit,
		AngularSpeed: 2.0,
		ReloadTime:   600,
		ReloadSpeed:  10,
	}
	tick := time.Tick(time.Second)
	frame := 0
	for !win.Closed() {
		win.Clear(color.RGBA{A: 0xff, R: 0xc8, G: 0xc8, B: 0xc8})
		t.update(win)
		t.draw(win, tankFrames, pic)
		win.Update()
		select {
		case <-tick:
			win.SetTitle(fmt.Sprintf("Gotanks |%d FPS", frame))
			frame = 0
		default:
		}
		frame++
	}
}

type Tank struct {
	X                     float64
	Y                     float64
	ReloadTime            float64
	CurrentReloadCoundown float64
	ReloadSpeed           float64
	Speed                 float64
	AngularSpeed          float64
	Buttons               [5]pixelgl.Button
	color                 color.RGBA
	Angle                 float64
	Radius                float64
	id                    int
}

func (t *Tank) move(vec pixel.Vec) {
	t.X += vec.X
	t.Y += vec.Y
}

func (t *Tank) draw(target pixel.Target, frames []pixel.Rect, spritesheet pixel.Picture) {
	head := pixel.NewSprite(spritesheet, frames[1])
	body := pixel.NewSprite(spritesheet, frames[0])
	bat := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	head.Draw(bat, pixel.IM.Scaled(pixel.ZV, 4).Rotated(pixel.ZV, Radians(t.Angle)).Moved(pixel.V(t.X, t.Y)))
	body.Draw(bat, pixel.IM.Scaled(pixel.ZV, 4).Rotated(pixel.ZV, Radians(t.Angle)).Moved(pixel.V(t.X, t.Y)))
	bat.Draw(target)
}

func (t *Tank) update(window *pixelgl.Window) {

	if window.Pressed(t.Buttons[2]) {
		t.Angle += t.AngularSpeed
	}
	if window.Pressed(t.Buttons[3]) {
		t.Angle -= t.AngularSpeed
	}
	for t.Angle > 360 {
		t.Angle -= 360
	}
	for t.Angle < 0 {
		t.Angle += 360
	}
	if window.Pressed(t.Buttons[0]) {
		t.move(pixel.V(1, 0).Rotated(Radians(t.Angle)).Scaled(t.Speed))
	}
	if window.Pressed(t.Buttons[1]) {
		t.move(pixel.V(1, 0).Rotated(Radians(t.Angle)).Scaled(-t.Speed))
	}
	if window.Pressed(t.Buttons[4]) {
		if t.CurrentReloadCoundown <= 0 {
			fmt.Println("Shooted")
			t.shoot()
			t.CurrentReloadCoundown = t.ReloadTime
		}
	}
	if t.CurrentReloadCoundown > 0 {
		t.CurrentReloadCoundown -= t.ReloadSpeed
	}
}

func (t *Tank) shoot() {

}

type Bullet struct {
	X         float64
	Y         float64
	Direction pixel.Vec
	Speed     float64
	Radius    float64
}

func (b *Bullet) update() {
	b.X += b.Direction.X * b.Speed
	b.Y += b.Direction.Y * b.Speed
}

func (b *Bullet) draw(target pixel.Target) {
	im := imdraw.New(nil)
	im.Push(pixel.V(b.X, b.Y))
	im.Circle(b.Radius, 0)
	im.Draw(target)

}

type Controllable interface {
	move(pixel.Vec)
}
type Drawable interface {
	draw(pixel.Target, float64)
}
type Updatable interface {
	update()
}
type Shootable interface {
	shoot()
}
