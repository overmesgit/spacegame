package aliengame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

var (
	shipsTiles       *ebiten.Image
	environmentTiles *ebiten.Image
	shipImg          *ebiten.Image
	shellImg         *ebiten.Image
	bigShellImg      *ebiten.Image
	alienImg         *ebiten.Image
	explosionImg     *ebiten.Image
)

const (
	shipTileSize        = 32
	environmentTileSize = 16
)

func init() {
	var err error

	shipsTiles, _, err = ebitenutil.NewImageFromFile("files/ships_packed.png")
	if err != nil {
		log.Fatal(err)
	}

	environmentTiles, _, err = ebitenutil.NewImageFromFile("files/tiles_packed.png")
	if err != nil {
		log.Fatal(err)
	}

	shipImg = GetShipImage(0, 1, 0, 0, 4, 4)
	alienImg = GetShipImage(1, 0, 0, 0, 4, 4)

	shellImg = GetEnvironmentImage(0, 0, 5, 5, 0, 0)
	explosionImg = GetEnvironmentImage(5, 0, 0, 0, 0, 0)
	bigShellImg = GetEnvironmentImage(0, 1, 4, 4, 0, 0)

}

func GetShipImage(x, y int, left, right, top, bottom int) *ebiten.Image {
	x = x * shipTileSize
	y = y * shipTileSize
	return shipsTiles.SubImage(image.Rect(x+left, y+top, x+shipTileSize-right, y+shipTileSize-bottom)).(*ebiten.Image)
}

func GetEnvironmentImage(x, y int, left, right, top, bottom int) *ebiten.Image {
	x = x * environmentTileSize
	y = y * environmentTileSize
	return environmentTiles.SubImage(image.Rect(x+left, y+top, x+environmentTileSize-right, y+environmentTileSize-bottom)).(*ebiten.Image)
}
