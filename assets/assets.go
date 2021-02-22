package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
)

//go:embed *
var files embed.FS

var assetCache = make(map[string]*ebiten.Image)

func LoadFont() font.Face {
	f, err := files.Open(`ArcadeClassic.TTF`)
	buf, _ := ioutil.ReadAll(f)
	tt, err := opentype.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}
	res, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     72,
	})
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func LoadImage(name string) *ebiten.Image {
	if v, ok := assetCache[name]; ok {
		return v
	}
	f, err := files.Open(name)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	res := ebiten.NewImageFromImage(img)
	assetCache[name] = res
	return res
}
