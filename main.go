package main

import (
	"math"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

var (
	frameCount = 0
	gravity    = 1.0

	r1  = 200.0
	r2  = 200.0
	m1  = 20.0
	m2  = 20.0
	a1  = math.Pi / 2
	a2  = math.Pi / 2
	a1V = 0.0
	a2V = 0.0
)

func main() {
	w, h := float64(900), float64(700)

	wnd, cv, err := sdlcanvas.CreateWindow(int(w), int(h), "Double Pendulum")
	if err != nil {
		panic(err)
	}

	defer wnd.Destroy()

	wnd.MainLoop(func() {
		n1 := -gravity * (2*m1 + m2) * math.Sin(a1)
		n2 := -m2 * gravity * math.Sin(a1-2*a2)
		n3 := -2 * math.Sin(a1-a2) * m2
		n4 := a2V*a2V*r2 + a1V*a1V*r1*math.Cos(a1-a2)
		den := r1 * (2*m1 + m2 - m2*math.Cos(2*a1-2*a2))
		a1A := (n1 + n2 + n3*n4) / den

		n1 = 2 * math.Sin(a1-a2)
		n2 = a1V * a1V * r1 * (m1 + m2)
		n3 = gravity * (m1 + m2) * math.Cos(a1)
		n4 = a2V * a2V * r2 * m2 * math.Cos(a1-a2)
		den = r2 * (2*m1 + m2 - m2*math.Cos(2*a1-2*a2))
		a2A := (n1 * (n2 + n3 + n4)) / den

		// clean
		cv.SetTransform(1, 0, 0, 1, 0, 0)
		cv.ClearRect(0, 0, w, h)

		// white background
		cv.SetFillStyle("#fff")
		cv.FillRect(0, 0, w, h)

		// translate center
		cv.Translate(w/2, 200)

		x1 := r1 * math.Sin(a1)
		y1 := r1 * math.Cos(a1)
		x2 := x1 + r2*math.Sin(a2)
		y2 := y1 + r2*math.Cos(a2)
		drawPendulum(cv, 0, 0, x1, y1)
		drawPendulum(cv, x1, y1, x2, y2)

		a1V += a1A
		a2V += a2A
		a1 += a1V
		a2 += a2V

		// damping
		//a1V *= 0.999
		//a2V *= 0.999
	})
}

func drawPendulum(cv *canvas.Canvas, sx, sy, x, y float64) {
	// rod
	cv.SetLineWidth(2)
	cv.SetStrokeStyle("#6ad7e5")
	cv.BeginPath()
	cv.MoveTo(sx, sy)
	cv.LineTo(x, y)
	cv.ClosePath()
	cv.Stroke()

	// bob
	cv.BeginPath()
	cv.Ellipse(x, y, m1, m1, 0, 0, 2*math.Pi, false)
	cv.SetFillStyle("#6ad7e5")
	cv.ClosePath()
	cv.Fill()
}
