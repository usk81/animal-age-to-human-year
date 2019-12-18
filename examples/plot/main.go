package main

import (
	"log"
	"os"

	"github.com/usk81/animal-age-to-human-year/dog"
	"github.com/wcharczuk/go-chart"
)

const max = 20

func main() {
	ages := make([]float64, max+1)
	pt1 := make([]float64, max+1)
	pt2 := make([]float64, max+1)

	for i := 0; i <= max; i++ {
		fi := float64(i)
		ages = append(ages, fi)
		pt1 = append(pt1, float64(dog.Calc(fi)))
		pt2 = append(pt2, float64(dog.WrongCalc(fi)))
	}

	ts1 := chart.ContinuousSeries{
		Name:    "Calc",
		XValues: ages,
		YValues: pt1,
	}
	ts2 := chart.ContinuousSeries{
		Name:    "WrongCalc",
		XValues: ages,
		YValues: pt2,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Dog age",
			// ValueFormatter: chart.TimeMinuteValueFormatter, //TimeHourValueFormatter,
		},
		YAxis: chart.YAxis{
			Name: "Human year",
		},
		Series: []chart.Series{
			ts1,
			ts2,
		},
	}

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = graph.Render(chart.PNG, f)
	if err != nil {
		log.Fatal(err)
	}
}
