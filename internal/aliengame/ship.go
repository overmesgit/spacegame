package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const shellTag = "shell"

type Ship struct {
	vx int
	vy int

	reloadTime int
	reload     int

	bigReloadTime int
	bigReload     int

	speed int
}

var _ GameObject = (*Ship)(nil)

func NewShip(x int, y int) *resolv.Object {
	w, h := shipImg.Size()
	obj := resolv.NewObject(float64(x), float64(y), float64(w), float64(h))
	obj.Data = &Ship{speed: 3, reloadTime: 10, bigReloadTime: 30}
	return obj
}

func (s *Ship) Update(obj *resolv.Object, actions []Action) {
	s.reload -= 1
	s.bigReload -= 1

	s.vx = 0
	s.vy = 0
	for _, action := range actions {
		switch action {
		case MoveLeft:
			s.vx = -s.speed
		case MoveRight:
			s.vx = s.speed
		case MoveUp:
			s.vy = -s.speed
		case MoveDown:
			s.vy = s.speed
		case Shoot:
			if s.reload <= 0 {
				obj.Space.Add(NewShell(obj.X, obj.Y))
				s.reload = s.reloadTime
			}
		case BigShoot:
			if s.bigReload <= 0 {
				obj.Space.Add(NewBigExplosion(obj.X, obj.Y))
				s.bigReload = s.bigReloadTime
			}
		}
	}

	newX := obj.X + float64(s.vx)
	newY := obj.Y + float64(s.vy)
	if (s.vx != 0 || s.vy != 0) && newX > 20 && newX < screenWidth-20 && newY > 20 && newY < screenHeight-20 {
		obj.X = newX
		obj.Y = newY
		obj.Update()
	}

}

func (s *Ship) Collision(left *resolv.Object, collision *resolv.Collision) {
	return
}

func (s *Ship) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(shipImg, op)
}

type Shell struct {
	vx    int
	vy    int
	speed int
}

func NewShell(x float64, y float64) *resolv.Object {
	w, h := shellImg.Size()
	obj := resolv.NewObject(x, y, float64(w), float64(h), shellTag)
	obj.Data = &Shell{0, -shellSpeed, 0}
	return obj
}

func (s Shell) Update(obj *resolv.Object, actions []Action) {
	if obj.Y < 0 {
		obj.Space.Remove(obj)
	} else {
		obj.X += float64(s.vx)
		obj.Y += float64(s.vy)
		obj.Update()
	}
}

func (a *Shell) Collision(left *resolv.Object, collision *resolv.Collision) {
	return
}

func (s Shell) Draw(obj *resolv.Object, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(obj.X, obj.Y)
	screen.DrawImage(shellImg, op)
}
