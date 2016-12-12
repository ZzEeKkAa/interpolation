package main

import (
	"fmt"
	"math"

	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {
	x0, h, n := 0., 0.2, 8
	var points []Point

	for x, i := x0, 0; i < n; i++ {
		points = append(points, Point{x, f(x), df(x)})
		x += h
	}

	polLag := Lagranzh(points...)
	polNew1 := NewtonFirst(h, points...)
	polNew2 := NewtonSecond(h, points...)
	pols := BuildSplines(h, points...)

	fmt.Println(points)
	fmt.Println(polLag)
	fmt.Println(polNew1)
	fmt.Println(polNew2)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Interpolation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	fPlotter := plotter.NewFunction(f)
	fPlotter.Color = color.RGBA{B: 255, A: 255}

	lagPlotter := plotter.NewFunction(makePolunomialFunctinon(polLag))
	lagPlotter.Color = color.RGBA{R: 255, A: 100}

	new1Plotter := plotter.NewFunction(makePolunomialFunctinon(polNew1))
	new1Plotter.Color = color.RGBA{G: 255, A: 100}

	new2Plotter := plotter.NewFunction(makePolunomialFunctinon(polNew2))
	new2Plotter.Color = color.RGBA{G: 255, R: 255, A: 100}

	splinePlotter := plotter.NewFunction(makeSplineFunction(pols, points))
	splinePlotter.Color = color.RGBA{G: 255, B: 255, A: 100}

	p.Add(fPlotter)
	p.Add(lagPlotter)
	p.Add(new1Plotter)
	p.Add(new2Plotter)
	p.Add(splinePlotter)
	p.X.Min = x0
	p.X.Max = x0 + float64(n-1)*h
	p.Y.Min = f(x0)
	p.Y.Max = f(x0 + float64(n-1)*h)

	if err := p.Save(14*vg.Inch, 14*vg.Inch, "interpolation8.png"); err != nil {
		panic(err.Error())
	}

}

func f(x float64) float64 {
	return math.Pow(x, 7) + 5*x - 6
}

func df(x float64) float64 {
	return 7*math.Pow(x, 6) + 5
}

func makePolunomialFunctinon(p Polynomial) func(float64) float64 {
	return func(x float64) float64 {
		var res float64
		for i, v := range p.a {
			res += math.Pow(x, float64(i)) * v
		}
		return res
	}
}

func makeSplineFunction(polynomials []Polynomial, points []Point) func(float64) float64 {
	return func(x float64) float64 {
		var res float64
		if x < points[0].x || x > points[len(points)-1].x {
			res = 0
		} else {
			j := 0
			for i, v := range points {
				if v.x > x {
					j = i - 1
					break
				}
			}
			for i, v := range polynomials[j].a {
				res += math.Pow(x, float64(i)) * v
			}
		}
		return res
	}
}
