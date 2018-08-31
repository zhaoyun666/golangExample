package test

import (
	"image"
	"os"
	"testing"
	//"bytes"
	//"bufio"
	"bytes"
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image/png"
	"log"
)

func TestImageByte(t *testing.T) {
	//f, _ := os.Open("./testData/8313b8c76d8fd73026c9901061b3ad9e.png")
	f, _ := os.Open("./testData/4c9d985c3f34ffc2c5bae7d93d6eccb5.jpg")
	defer f.Close()
	bt := readFileToBytes(f)
	fmt.Println(len(bt), cap(bt))
	bbb := bytes.NewReader(bt)
	img, imgType, _ := image.Decode(bbb)

	//fmt.Println(img)
	fmt.Println(imgType)
	//return
	dst := image.NewRGBA(image.Rect(0, 0, 80, 80))
	err := graphics.Scale(dst, img) //缩小图片
	if err != nil {
		log.Fatal(err)
	}
	SaveImage("./testData/thumbnailimg.jpg", dst)
}

func readFileToBytes(f *os.File) []byte {
	chunks := make([]byte, 0)
	for {
		buf := make([]byte, 1024)
		switch n, err := f.Read(buf[:]); true {
		case n < 0:
			panic(err)
		case n == 0: //EOF
			return chunks
		case n > 0:
			chunks = append(chunks, buf...)
		}
	}
}

func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)

	defer imgfile.Close()
	err = png.Encode(imgfile, img) //编码图片
	if err != nil {
		log.Fatal("Save fail:", err)
	}
	return
}

func TestMap(t *testing.T) {
	var a = map[string]interface{}{
		"Msg": "xxx",
	}
	fmt.Println(a)
}
