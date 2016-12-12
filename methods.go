package main

func Lagranzh(p ...Point) Polynomial {
	res := Polynomial{[]float64{0.}}
	for _, p1 := range p { // j
		tp := Polynomial{[]float64{1.}}
		for _, p2 := range p { // i
			if p1 != p2 {
				pp := NewSimplePolynomial(-p2.x/(p1.x-p2.x), 1/(p1.x-p2.x))
				tp.Mul(pp)
			}
		}
		tp.MulK(p1.y)
		res.Add(tp)
	}
	return res
}

func NewtonFirst(h float64, p ...Point) Polynomial {
	var f func(int, int) float64
	f = func(pos, pow int) float64 {
		if pow == 0 {
			return p[pos].y
		}
		return f(pos+1, pow-1) - f(pos, pow-1)
	}

	res := Polynomial{[]float64{0.}}
	q := NewSimplePolynomial(1.-p[0].x/h, 1/h)
	pol := Polynomial{[]float64{1.}}
	for i := range p {
		fy := f(0, i)
		pol.MulK(fy)
		res.Add(pol)
		pol.MulK(1. / fy)

		//////////////////

		q.Add(Polynomial{[]float64{-1.}})
		pol.Mul(q)
		pol.MulK(1. / float64(i+1))
	}

	return res
}

func NewtonSecond(h float64, p ...Point) Polynomial {
	var f func(int, int) float64
	f = func(pos, pow int) float64 {
		if pow == 0 {
			return p[pos].y
		}
		return f(pos+1, pow-1) - f(pos, pow-1)
	}

	res := Polynomial{[]float64{0.}}
	q := NewSimplePolynomial(-1.-p[len(p)-1].x/h, 1/h)
	pol := Polynomial{[]float64{1.}}
	for i := range p {
		fy := f(len(p)-1-i, i)
		pol.MulK(fy)
		res.Add(pol)
		pol.MulK(1. / fy)

		//////////////////

		q.Add(Polynomial{[]float64{1.}})
		pol.Mul(q)
		pol.MulK(1. / float64(i+1))
	}

	return res
}

func BuildSplines(h float64, p ...Point) []Polynomial {
	n := len(p)
	var res []Polynomial
	for i := 0; i < n-1; i++ {
		pol := Polynomial{[]float64{1.}}
		pol.Mul(Polynomial{[]float64{p[i+1].x / h, -1. / h}})
		pol.Mul(Polynomial{[]float64{p[i+1].x / h, -1. / h}})
		pol.Mul(Polynomial{[]float64{(h - 2*p[i].x) / h, 2. / h}})
		pol.MulK(p[i].y)

		pol2 := Polynomial{[]float64{1.}}
		pol2.Mul(Polynomial{[]float64{p[i].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{p[i].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{(h + 2*p[i+1].x) / h, -2. / h}})
		pol2.MulK(p[i+1].y)

		pol.Add(pol2)

		pol2 = Polynomial{[]float64{1.}}
		pol2.Mul(Polynomial{[]float64{p[i+1].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{p[i+1].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{-p[i].x, 1.}})
		pol2.MulK(p[i].m)

		pol.Add(pol2)

		pol2 = Polynomial{[]float64{1.}}
		pol2.Mul(Polynomial{[]float64{p[i].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{p[i].x / h, -1. / h}})
		pol2.Mul(Polynomial{[]float64{-p[i+1].x, 1.}})
		pol2.MulK(p[i+1].m)

		pol.Add(pol2)
		res = append(res, pol)
	}

	return res
}
