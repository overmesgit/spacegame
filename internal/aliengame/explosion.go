package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"math"
)

type Explosion struct {
	Exists int
}

func (e *Explosion) Collision(left *resolv.Object, collision *resolv.Collision) {
	return
}

var _ GameObject = (*Explosion)(nil)

func NewExplosion(x float64, y float64) *resolv.Object {
	w, h := explosionImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h))
	obj.Data = &Explosion{}
	return obj
}

func (e *Explosion) Update(obj *resolv.Object, actions []Action) {
	e.Exists += 1
	if e.Exists > 60 {
		obj.Space.Remove(obj)
	}
}

func (e *Explosion) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := explosionImg.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(math.Pi * float64(e.Exists) / 180)
	scale := math.Pow(float64(e.Exists), 2)/math.Pow(60, 2) + 1
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(explosionImg, op)
}
