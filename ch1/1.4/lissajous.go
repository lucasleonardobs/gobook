// Lissajous gera animações GIF de figuras lissajous aleatórias
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color
	blackIndex = 1 // next color
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 2    // número de revoluções completas do oscilador x
		res     = 0.01 // resolução angular
		size    = 256  // canvas da imagem cobre de [-size...+size]
		nframes = 512  // número de quadros da animação
		delay   = 8    // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 // frequencia relativa do oscilador your
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for j := 0.0; j < cycles*2*math.Pi; j += res {
			x := math.Sin(j)
			y := math.Sin(j*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de ondificação
}
