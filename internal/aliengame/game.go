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
	Collision(left *resolv.Object, collision *resolv.Collision)
	Draw(obj *resolv.Object, screen *ebiten.Image)
}

func (g *Game) Update() error {

	dx, dy := shellImg.Size()
	for _, obj := range g.Objects() {
		if collision := obj.Check(float64(dx), float64(dy), shellTag); collision != nil {
			gameObj, ok := obj.Data.(GameObject)
			if ok {
				gameObj.Collision(obj, collision)
			}
		}
	}

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	controls := Control(g.keys)
	for _, obj := range g.Objects() {
		gameObj, ok := obj.Data.(GameObject)
		if ok {
			gameObj.Update(obj, controls)
		}
	}

	AddAlien(g)
	//if rand.Intn(10) == 0 {
	//	AddAlien(g)
	//}

	return nil
}

func GetClosestCollision(collision *resolv.Collision) (vector.Vector, *resolv.Object) {
	smallestDist := collision.ContactWithObject(collision.Objects[0])
	closesObject := collision.Objects[0]
	for i := range collision.Objects[1:] {
		dist := collision.ContactWithObject(collision.Objects[i])
		if smallestDist.X()+smallestDist.Y() > dist.Y()+dist.X() {
			smallestDist = dist
			closesObject = collision.Objects[i]
		}
	}
	return smallestDist, closesObject
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff})

	objects := g.Objects()
	sort.SliceStable(objects, func(i, j int) bool {
		return objects[i].X < objects[j].X
	})

	for _, obj := range objects {
		gameObj, ok := obj.Data.(GameObject)
		if ok {
			gameObj.Draw(obj, screen)
		}

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
