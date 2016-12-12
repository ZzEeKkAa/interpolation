package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y, m float64
}

type Polynomial struct {
	a []float64
}

func (p *Polynomial) Add(p2 Polynomial) *Polynomial {
	n := len(p.a)
	if len(p2.a) > n {
		n = len(p2.a)
		tmp := p.a
		p.a = make([]float64, n)
		copy(p.a, tmp)
	}
	for i, v := range p2.a {
		p.a[i] += v
	}
	return p
}

func (p *Polynomial) MulK(k float64) *Polynomial {
	for i := range p.a {
		p.a[i] *= k
	}
	return p
}

func (p *Polynomial) Mul(p2 Polynomial) *Polynomial {
	res := make([]float64, len(p.a)+len(p2.a)-1)
	for i, a := range p.a {
		for j, b := range p2.a {
			res[i+j] += a * b
		}
	}
	p.a = res
	return p
}

func (p *Polynomial) LoadFromFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	var n int
	fmt.Fscanf(f, "%d", &n)

	p.a = make([]float64, n+1)
	for i := range p.a {
		fmt.Fscanf(f, "%f", &p.a[i])
	}
	return err
}

func NewSimplePolynomial(a, b float64) Polynomial {
	return Polynomial{[]float64{a, b}}
}
