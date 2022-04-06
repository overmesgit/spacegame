package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type BigExplosion struct {
	Exists   int
	Exploded bool
}

var _ GameObject = (*BigExplosion)(nil)

func NewBigExplosion(x float64, y float64) *resolv.Object {
	w, h := explosionImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h), shellTag)
	obj.Data = &BigExplosion{}
	return obj
}

func (e *BigExplosion) Update(obj *resolv.Object, actions []Action) {
	e.Exists += 1
	obj.Y -= 4
	obj.Update()
}

func (e *BigExplosion) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(bigShellImg, op)
}

func (e *BigExplosion) Collision(left *resolv.Object, collision *resolv.Collision) {

	//for _, right := range collision.Objects {
	//
	//}
}
