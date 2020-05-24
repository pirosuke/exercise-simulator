package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

/*
WeeklyExercisePlans describes weekley exercising plans.
*/
type WeeklyExercisePlans struct {
	Title        string  `json:"title"`
	METs         float64 `json:"mets"`
	Hours        float64 `json:"hours"`
	TimesPerWeek float64 `json:"times_per_week"`
}

/*
PlansConfig describes setting data format.
*/
type PlansConfig struct {
	Weight        float64               `json:"weight"`
	WeeksToOutput int64                 `json:"weeks_to_output"`
	WeeklyPlans   []WeeklyExercisePlans `json:"weekly_plans"`
}

/*
WeeklyResult is result data of week.
*/
type WeeklyResult struct {
	Week         int64   `json:"week"`
	StartWeight  float64 `json:"start_weight_kg"`
	ExerciseKCal float64 `json:"exercise_kcal"`
	RemovedFat   float64 `json:"removed_fat"`
}

/*
SimulatedResult is result data to output.
*/
type SimulatedResult struct {
	Title         string         `json:"title"`
	Weeks         int64          `json:"weeks"`
	StartWeight   float64        `json:"start_weight_kg"`
	EndWeight     float64        `json:"end_weight_kg"`
	WeeklyResults []WeeklyResult `json:"weekly_results"`
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func main() {
	planFilePath := flag.String("p", "", "Plans file path")
	flag.Parse()

	if !fileExists(*planFilePath) {
		fmt.Println("plans file does not exist")
		return
	}

	jsonContent, err := ioutil.ReadFile(*planFilePath)
	if err != nil {
		fmt.Println("failed to read plans file")
		return
	}

	plansConfig := new(PlansConfig)
	if err := json.Unmarshal(jsonContent, plansConfig); err != nil {
		fmt.Println("failed to read plans file")
		return
	}

	result := SimulatedResult{}
	result.Title = *planFilePath
	result.StartWeight = plansConfig.Weight
	result.Weeks = plansConfig.WeeksToOutput

	remainingWeight := plansConfig.Weight
	var week int64
	var weeklyResults []WeeklyResult
	for week = 1; week <= plansConfig.WeeksToOutput; week++ {
		weeklyResult := WeeklyResult{
			Week:        week,
			StartWeight: remainingWeight,
		}

		weeklyExerciseKCal := 0.0
		weeklyRemovedFat := 0.0
		for _, plans := range plansConfig.WeeklyPlans {
			usedKCal := 1.05 * plans.METs * plans.Hours * remainingWeight * plans.TimesPerWeek
			removedFat := usedKCal / 9 * 1.25 / 1000

			weeklyExerciseKCal = weeklyExerciseKCal + usedKCal
			weeklyRemovedFat = weeklyRemovedFat + removedFat

		}

		weeklyResult.ExerciseKCal = weeklyExerciseKCal
		weeklyResult.RemovedFat = weeklyRemovedFat

		weeklyResults = append(weeklyResults, weeklyResult)
		remainingWeight = remainingWeight - weeklyRemovedFat
	}

	result.EndWeight = remainingWeight
	result.WeeklyResults = weeklyResults

	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("failed to create result data")
		return
	}

	fmt.Println(string(resultJSON))
}
