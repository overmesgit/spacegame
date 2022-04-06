package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"math"
)

type Explosion struct {
	Exists   int
	Lifetime int

	vx, vy float64
}

const ExplosionTime = 30

func (e *Explosion) Collision(left *resolv.Object, collision *resolv.Collision) {
	return
}

var _ GameObject = (*Explosion)(nil)

func NewExplosion(x float64, y float64) *resolv.Object {
	w, h := explosionImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h), shellTag)
	obj.Data = &Explosion{Lifetime: ExplosionTime}
	return obj
}

func NewChildBigExplosion(x float64, y float64, vx, vy float64) *resolv.Object {
	w, h := explosionImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h), shellTag)
	obj.Data = &Explosion{Lifetime: ExplosionTime * 3, vx: vx, vy: vy}
	return obj
}

func (e *Explosion) Update(obj *resolv.Object, actions []Action) {
	e.Exists += 1
	if e.Exists > e.Lifetime {
		obj.Space.Remove(obj)
	}

	scale := math.Pow(float64(e.Exists), 2)/math.Pow(60, 2) + 1
	w, h := explosionImg.Size()
	obj.W = float64(w) * scale
	obj.H = float64(h) * scale
	if e.vx != 0 || e.vy != 0 {
		obj.X += e.vx
		obj.Y += e.vy
	}
	obj.Update()
}

func (e *Explosion) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := explosionImg.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(math.Pi * float64(e.Exists) / 90)
	scale := math.Pow(float64(e.Exists), 2)/math.Pow(60, 2) + 1
	scale = math.Min(scale, 3)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(explosionImg, op)
}
