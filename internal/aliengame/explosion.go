package aliengame

import (
	"github.com/kvartborg/vector"
	"github.com/solarlune/resolv"
)

type Explosion struct {
	Exists int
}

func NewExplosion(x float64, y float64) *resolv.Object {
	w, h := explosionImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h))
	obj.Data = &Explosion{}
	return obj
}

func (s *Explosion) Update(obj *resolv.Object, actions []Action) {
	s.Exists += 1
	if s.Exists > 120 {
		obj.Space.Remove(obj)
	}
}

func (s *Explosion) Collision(left *resolv.Object, right *resolv.Object, dist vector.Vector) {
	return
}
