package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"math"
	"math/rand"
)

const (
	SPEED = 1
)

type Alien struct {
	vx int
	vy int

	dead    bool
	timeout int
}

func NewAlien(x int, y int) *resolv.Object {
	w, h := alienImg.Size()
	obj := resolv.NewObject(float64(x), float64(y), float64(w), float64(h))
	obj.Data = &Alien{}
	return obj
}

func RandomAlien(limitX int, limitY int, offset int) *resolv.Object {
	x := rand.Intn(limitX-2*offset) + offset
	y := rand.Intn(limitY)
	return NewAlien(x, y)
}

func (a *Alien) Update(obj *resolv.Object, actions []Action) {
	a.timeout -= 1

	if a.timeout <= 0 {
		a.timeout = 60 + rand.Intn(60)

		a.vx = rand.Intn(SPEED*2+1) - SPEED
		a.vy = rand.Intn(SPEED) + 1
	}

	obj.X += float64(a.vx)
	obj.Y += float64(a.vy)
	obj.Update()
}

func (a *Alien) Collision(left *resolv.Object, collision *resolv.Collision) {

	for _, right := range collision.Objects {
		_, shellCollision := right.Data.(*Shell)
		if shellCollision {
			dist := collision.ContactWithObject(right)
			if dist.Y() > -4 || dist.X() > -3 || dist.X() < -29 {
				return
			}
			right.Space.Add(NewExplosion(left.X, left.Y))
			right.Space.Remove(right)
			left.Space.Remove(left)
			// this 2 breaks improve performance drastically
			break
		}
		_, explosionCollision := right.Data.(*Explosion)
		if explosionCollision {
			dist := collision.ContactWithObject(right)
			if dist.Y() > -16 || dist.X() > -16 {
				return
			}
			right.Space.Add(NewExplosion(left.X, left.Y))
			left.Space.Remove(left)
			break
		}
		_, bigShellCollision := right.Data.(*BigExplosion)
		if bigShellCollision {
			dist := collision.ContactWithObject(right)
			if dist.Y() > -4 || dist.X() > -3 || dist.X() < -29 {
				return
			}
			right.Space.Add(NewExplosion(left.X, left.Y))
			for i := 0; i < 32; i++ {
				angle := rand.Float64() * 2 * math.Pi
				vx := 2 * math.Cos(angle)
				vy := 2 * math.Sin(angle)
				right.Space.Add(NewChildBigExplosion(left.X, left.Y, vx, vy))
			}
			right.Space.Remove(right)
			left.Space.Remove(left)

			break
		}
	}
}

func (a *Alien) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := alienImg.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(math.Pi)
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(alienImg, op)
}
