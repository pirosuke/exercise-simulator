package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

/*
ExercisePlans describes exercising plans.
*/
type ExercisePlans struct {
	Title        string  `json:"title"`
	METs         float64 `json:"mets"`
	Hours        float64 `json:"hours"`
	TimesPerWeek float64 `json:"times_per_week"`
}

/*
WeeklyPlans describes weekly exercising plans.
*/
type WeeklyPlans struct {
	Name  string          `json:"name"`
	Plans []ExercisePlans `json:"plans"`
}

/*
PlansConfig describes setting data format.
*/
type PlansConfig struct {
	Weight        float64       `json:"weight"`
	WeeksToOutput int64         `json:"weeks_to_output"`
	WeeklyPlans   []WeeklyPlans `json:"weekly_plans"`
}

/*
ExerciseResult is calc result data of weekly exercise.
*/
type ExerciseResult struct {
	Week         int64   `json:"week"`
	StartWeight  float64 `json:"start_weight_kg"`
	ExerciseKCal float64 `json:"exercise_kcal"`
	RemovedFat   float64 `json:"removed_fat_kg"`
}

/*
WeeklyResult is result data of week
*/
type WeeklyResult struct {
	Name       string           `json:"name"`
	EndWeight  float64          `json:"end_weight_kg"`
	RemovedFat float64          `json:"removed_fat_kg"`
	Results    []ExerciseResult `json:"results"`
}

/*
PlansSimulatedResult is result data to output.
*/
type PlansSimulatedResult struct {
	Title         string         `json:"title"`
	Weeks         int64          `json:"weeks"`
	StartWeight   float64        `json:"start_weight_kg"`
	WeeklyResults []WeeklyResult `json:"weekly_results"`
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func calcExerciseResult(startWeight float64, weeksToOutput int64, weeklyPlans WeeklyPlans) WeeklyResult {
	remainingWeight := startWeight

	var week int64
	var exerciseResults []ExerciseResult
	for week = 1; week <= weeksToOutput; week++ {
		weeklyResult := ExerciseResult{
			Week:        week,
			StartWeight: remainingWeight,
		}

		weeklyExerciseKCal := 0.0
		weeklyRemovedFat := 0.0
		for _, plans := range weeklyPlans.Plans {
			usedKCal := 1.05 * plans.METs * plans.Hours * remainingWeight * plans.TimesPerWeek
			removedFat := usedKCal / 9 * 1.25 / 1000

			weeklyExerciseKCal = weeklyExerciseKCal + usedKCal
			weeklyRemovedFat = weeklyRemovedFat + removedFat

		}

		weeklyResult.ExerciseKCal = weeklyExerciseKCal
		weeklyResult.RemovedFat = weeklyRemovedFat

		exerciseResults = append(exerciseResults, weeklyResult)
		remainingWeight = remainingWeight - weeklyRemovedFat
	}

	weeklyResult := WeeklyResult{
		Name:       weeklyPlans.Name,
		EndWeight:  remainingWeight,
		RemovedFat: startWeight - remainingWeight,
		Results:    exerciseResults,
	}

	return weeklyResult
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: exersim [flags]\n")
		flag.PrintDefaults()
	}

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

	result := PlansSimulatedResult{}
	result.Title = *planFilePath
	result.StartWeight = plansConfig.Weight
	result.Weeks = plansConfig.WeeksToOutput

	var weeklyResults []WeeklyResult
	for _, weeklyPlans := range plansConfig.WeeklyPlans {
		weeklyResult := calcExerciseResult(plansConfig.Weight, plansConfig.WeeksToOutput, weeklyPlans)
		weeklyResults = append(weeklyResults, weeklyResult)
	}

	result.WeeklyResults = weeklyResults

	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("failed to create result data")
		return
	}

	fmt.Println(string(resultJSON))
}
