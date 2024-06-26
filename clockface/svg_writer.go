package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteLineLength = 80
	hourLineLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, centerPin)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(MinuteHandPoint(t), minuteLineLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:4px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(HourHandPoint(t), hourLineLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#0f0;stroke-width:5px;"/>`, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}             // scale
	p = Point{p.X, -p.Y}                              // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} // translate
	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const centerPin = `
  <!-- center pin -->
  <circle cx="150" cy="150" r="4" style="fill:#fff;stroke:#000;stroke-width:2px;"/>
		
	`
const svgEnd = `</svg>`
