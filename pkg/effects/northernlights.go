package effects

import (
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

type NorthernLights struct {
	Threshold float64 // chance to generate new pixel
	Fade      float64 // how big percentage to fade out every iteration
	Cutoff    float64 // how small difference is needed to zero out a dying color
	Base      Color
	Color1    Color
	Color2    Color
}

func (e *NorthernLights) Draw(pixels []colorful.Color) {
	for i := range pixels {
		if e.Threshold < rand.Float64() {
			if pixels[i].DistanceLab(e.Base.Color) > e.Cutoff {
				pixels[i] = pixels[i].BlendHcl(e.Base.Color, e.Fade)
			} else {
				pixels[i] = e.Base.Color
			}
		} else {
			pixels[i] = e.Color1.BlendHcl(e.Color2.Color, rand.Float64())
		}
	}
}
