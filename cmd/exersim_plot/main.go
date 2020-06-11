package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/pirosuke/exercise-simulator/internal/models"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func createPoints(exerciseResults []models.ExerciseResult) plotter.XYs {
	pts := make(plotter.XYs, len(exerciseResults))
	for i, result := range exerciseResults {
		pts[i].X = float64(result.Week)
		pts[i].Y = result.StartWeight
	}

	return pts
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: exersim_plot [flags]\n")
		flag.PrintDefaults()
	}

	simulateResultFilePath := flag.String("f", "", "exersim result file path")
	flag.Parse()

	if !fileExists(*simulateResultFilePath) {
		fmt.Println("result file does not exist")
		return
	}

	jsonContent, err := ioutil.ReadFile(*simulateResultFilePath)
	if err != nil {
		fmt.Println("failed to read result file")
		return
	}

	simulatedResultData := new(models.PlansSimulatedResult)
	if err := json.Unmarshal(jsonContent, simulatedResultData); err != nil {
		fmt.Println("failed to read result file")
		return
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Exercise Result Simulation"
	p.X.Label.Text = "Week"
	p.Y.Label.Text = "Weight"

	p.Add(plotter.NewGrid())
	p.Legend.Left = true

	for _, weeklyResult := range simulatedResultData.WeeklyResults {
		points := createPoints(weeklyResult.Results)

		err = plotutil.AddLinePoints(p,
			weeklyResult.Name, points)

		if err != nil {
			panic(err)
		}
	}

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "exercise_simulation.png"); err != nil {
		panic(err)
	}

}
