package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Get some data.
	n, m := 5, 10
	pts := make([]plotter.XYer, n)
	for i := range pts {
		xys := make(plotter.XYs, m)
		pts[i] = xys
		center := float64(i)
		for j := range xys {
			xys[j].X = center + (rand.Float64() - 0.5)
			xys[j].Y = center + (rand.Float64() - 0.5)
		}
	}

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Create two lines connecting points and error bars. For
	// the first, each point is the mean x and y value and the
	// error bars give the 95% confidence intervals.  For the
	// second, each point is the median x and y value with the
	// error bars showing the minimum and maximum values.
	mean95, err := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pts...)
	if err != nil {
		panic(err)
	}
	medMinMax, err := plotutil.NewErrorPoints(plotutil.MedianAndMinMax, pts...)
	if err != nil {
		panic(err)
	}
	plotutil.AddLinePoints(plt,
		"mean and 95% confidence", mean95,
		"median and minimum and maximum", medMinMax)
	plotutil.AddErrorBars(plt, mean95, medMinMax)

	// Add the points that are summarized by the error points.
	plotutil.AddScatters(plt, pts[0], pts[1], pts[2], pts[3], pts[4])

	plt.Save(8*vg.Inch, 6*vg.Inch, "04-points-error_bars.png")
}