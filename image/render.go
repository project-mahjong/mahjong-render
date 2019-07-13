package image

import (
	model "github.com/project-mahjong/mahjong-backend/core"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"strconv"
)

var resourse []image.Image
var fileList []string

func init() {
	initFileList()
	readResourse()
}

func initFileList() {
	for k := 1; k <= 3; k++ {
		for i := 0; i <= 9; i++ {
			t := i
			fileList = append(fileList, strconv.Itoa(t)+"m")
			fileList = append(fileList, strconv.Itoa(t)+"p")
			fileList = append(fileList, strconv.Itoa(t)+"s")
		}
	}
	for i := 1; i <= 7; i++ {
		fileList = append(fileList, strconv.Itoa(i)+"z")
	}
}

func readResourse() {
	for _, i := range fileList {
		file, err := os.Open("resourse/" + i + ".png")
		if err != nil {
			log.Panicln("can't open resoure file")
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Panicln("can't read resoure file")
		}
		resourse = append(resourse, img)
	}
}

func RenderMahjong(title *model.TitleModel, renderType int) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 800, 1000))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
	for i := 0; i < 13; i++ {
		draw.Draw(img, image.Rectangle{image.Point{100 + 42*(i-1), 200}, image.Point{100 + 42*i, 200 + 59}}, resourse[i], image.ZP, draw.Src)
	}
	for i := 0; i < 13; i++ {
		draw.Draw(img, image.Rectangle{image.Point{100 + 42*(i-1), 1000 - 59}, image.Point{100 + 42*i, 1000}}, resourse[i], image.ZP, draw.Src)
	}
	return image.Image(img), nil
}
