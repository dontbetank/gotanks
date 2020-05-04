package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"google.golang.org/api/compute/v1"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"os"
	"time"
)

var (
	screenWidth, screenHeight, unit float64
)

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
	screenWidth, screenHeight = pixelgl.Monitors()[0].Size()
	unit = math.Sqrt(math.Pow(screenWidth, 2) + math.Pow(screenHeight, 2)) / 200
	pic, err := loadPicture("assets/images/tank.png")
	if err != nil {
		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, 120, 100),
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
		Radius:       2.5 * unit,
		Buttons:      ButtonConfig{
			Forward:  pixelgl.KeyW,
			Backward: pixelgl.KeyS,
			Right:    0,
			Left:     0,
			Shoot:    0,
		},
		Speed:        0.5 * unit,
		AngularSpeed: 2.0,
		ReloadTime:   600,
		ReloadSpeed:  10,
	}
	last := time.Now()
	for !win.Closed() {
		var dt float64 = time.Since(last).Seconds()
		win.Clear(color.RGBA{A: 0xff, R: 0xc8, G: 0xc8, B: 0xc8})
		t.update(win)
		t.draw(win, tankFrames, pic)
		win.Update()
	}
}

type ButtonConfig struct {
	Forward		pixelgl.Button
	Backward 	pixelgl.Button
	Right		pixelgl.Button
	Left		pixelgl.Button
	Shoot 		pixelgl.Button
}

type Tank struct {
	X                     float64
	Y                     float64
	ReloadTime            float64
	CurrentReloadCoundown float64
	ReloadSpeed           float64
	Speed                 float64
	AngularSpeed          float64
	Buttons				  ButtonConfig
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
	head.Draw(bat, pixel.IM.Scaled(pixel.ZV, 4).Rotated(pixel.ZV, t.Angle).Moved(pixel.V(t.X, t.Y)))
	body.Draw(bat, pixel.IM.Scaled(pixel.ZV, 4).Rotated(pixel.ZV, t.Angle).Moved(pixel.V(t.X, t.Y)))
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
		t.move(pixel.V(1, 0).Rotated(t.Angle).Scaled(t.Speed))
	}
	if window.Pressed(t.Buttons[1]) {
		t.move(pixel.V(1, 0).Rotated(t.Angle).Scaled(-t.Speed))
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
