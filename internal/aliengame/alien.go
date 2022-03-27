package aliengame

import (
	"github.com/kvartborg/vector"
	"github.com/solarlune/resolv"
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

func (a *Alien) Collision(left *resolv.Object, right *resolv.Object, dist vector.Vector) {
	if dist.Y() > -8 || dist.X() > -3 {
		return
	}
	_, shellCollision := right.Data.(*Shell)
	if shellCollision {
		right.Space.Add(NewExplosion(left.X, left.Y))
		right.Space.Remove(right)
		left.Space.Remove(left)
	}
}
