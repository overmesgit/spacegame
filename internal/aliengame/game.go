package aliengame

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kvartborg/vector"
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
	Collision(left *resolv.Object, right *resolv.Object, dist vector.Vector)
}

func (g *Game) Update() error {

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	controls := Control(g.keys)
	dx, dy := shellImg.Size()
	for _, obj := range g.Objects() {
		if collision := obj.Check(float64(dx), float64(dy), "shell"); collision != nil {
			smallestDist := collision.ContactWithObject(collision.Objects[0])
			closesObject := collision.Objects[0]
			for i := range collision.Objects[1:] {
				dist := collision.ContactWithObject(collision.Objects[i])
				if smallestDist.X()+smallestDist.Y() > dist.Y()+dist.X() {
					smallestDist = dist
					closesObject = collision.Objects[i]
				}
			}
			gameObj, ok := obj.Data.(GameObject)
			if ok {
				gameObj.Collision(obj, closesObject, smallestDist)
			}
		}
	}

	for _, obj := range g.Objects() {
		gameObj, ok := obj.Data.(GameObject)
		if ok {
			gameObj.Update(obj, controls)
		}
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
			img = alienImg
			w, h := img.Size()
			op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
			op.GeoM.Rotate(math.Pi)
		case *Shell:
			img = shellImg
		case *Explosion:
			img = explosionImg
			explosion := obj.Data.(*Explosion)
			w, h := img.Size()
			op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
			op.GeoM.Rotate(math.Pi * float64(explosion.Exists) / 180)
			//var scale float64
			//val := math.Pow(float64(explosion.Exists), 1.5)
			//if val/(val/2) < 2 {
			//	scale = val / (val / 2) / 120
			//} else {
			//	scale = val / (val / 2)
			//}
			//op.GeoM.Scale(scale, scale)
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
