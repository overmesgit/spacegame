package aliengame

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"sort"
)

const (
	screenWidth  = 1280
	screenHeight = 960
	shellSpeed   = 5
)

type Game struct {
	keys []ebiten.Key
	*resolv.Space
}

type GameObject interface {
	Update(obj *resolv.Object, actions []Action)
	Collision(left *resolv.Object, right *resolv.Object)
}

var shipImg *ebiten.Image
var shellImg *ebiten.Image

func init() {
	var err error
	shipImg, _, err = ebitenutil.NewImageFromFile("files/ship_0000.png")
	if err != nil {
		log.Fatal(err)
	}

	shellImg, _, err = ebitenutil.NewImageFromFile("files/tile_0000.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	controls := Control(g.keys)
	for _, obj := range g.Objects() {
		gameObj, ok := obj.Data.(GameObject)
		if collision := obj.Check(0, shellSpeed, "shell"); collision != nil {
			if ok {
				dist := collision.ContactWithObject(collision.Objects[0])
				//fmt.Println(gameObj, dist)
				if dist.Y() < 0 {
					gameObj.Collision(obj, collision.Objects[0])
				}
			}
		}
		gameObj.Update(obj, controls)
	}

	AddAlien(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff})

	img := shipImg
	objects := g.Objects()

	sort.SliceStable(objects, func(i, j int) bool {
		return objects[i].X < objects[j].X
	})

	for _, obj := range objects {
		op := &ebiten.DrawImageOptions{}

		switch obj.Data.(type) {
		case *Ship:
			img = shipImg
		case *Alien:
			img = shipImg
			op.GeoM.Rotate(math.Pi)
		case *Shell:
			img = shellImg
		}

		op.GeoM.Translate(obj.X, obj.Y)
		screen.DrawImage(img, op)

	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %v", len(objects)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func AddAlien(g *Game) {
	alien := RandomAlien(screenWidth, 100, 100)
	g.Add(alien)
}

func RunGame() {
	//f, err := os.Create("prof")
	//if err != nil {
	//	log.Fatal("could not create CPU profile: ", err)
	//}
	//defer f.Close() // error handling omitted for example
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal("could not start CPU profile: ", err)
	//}
	//defer pprof.StopCPUProfile()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Spacecraft!")
	space := resolv.NewSpace(screenWidth, screenHeight, 32, 32)
	game := Game{Space: space}
	game.Add(NewShip(screenWidth/2, screenHeight-50))
	for i := 0; i < 50; i++ {
		AddAlien(&game)
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
