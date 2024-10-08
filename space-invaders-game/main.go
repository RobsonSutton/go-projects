package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//
//go:embed images
var images embed.FS

// Load Ship
var ship = MustLoadImage("images/ship.png")
var title = "Space Invaders"

// Game implements ebiten.Game interface.
type Game struct{}

// Open image, decode and store as object
func MustLoadImage(name string) *ebiten.Image {
	file, err := images.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(150, 200)
	//op.GeoM.Scale(0.25, 0.25)
	screen.DrawImage(ship, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a tripled size is specified.
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle(title)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
