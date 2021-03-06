// Copyright 2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build example

package main

import (
	"image/color"
	"log"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"go-hep.org/x/hep/hbook"
	"go-hep.org/x/hep/hplot"
)

func main() {
	const npoints = 10000

	// Create a normal distribution.
	dist := distuv.Normal{
		Mu:    0,
		Sigma: 1,
		Src:   rand.New(rand.NewSource(0)),
	}

	// Draw some random values from the standard
	// normal distribution.
	hist := hbook.NewH1D(20, -4, +4)
	for i := 0; i < npoints; i++ {
		v := dist.Rand()
		hist.Fill(v, 1)
	}

	// normalize histogram
	area := 0.0
	for _, bin := range hist.Binning.Bins {
		area += bin.SumW() * bin.XWidth()
	}
	hist.Scale(1 / area)

	// Make a plot and set its title.
	p := hplot.New()
	p.Title.Text = "Histogram"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Create a histogram of our values drawn
	// from the standard normal.
	h := hplot.NewH1D(hist)
	h.Infos.Style = hplot.HInfoSummary
	p.Add(h)

	// The normal distribution function
	norm := hplot.NewFunction(dist.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	// draw a grid
	p.Add(hplot.NewGrid())
	p.Add(plotter.NewGlyphBoxes())

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Centimeter, -1, "testdata/ex-00-hist.png"); err != nil {
		log.Fatal(err)
	}

	// Save the plot to a PDF file.
	if err := p.Save(15*vg.Centimeter, -1, "testdata/ex-00-hist.pdf"); err != nil {
		log.Fatal(err)
	}
}
