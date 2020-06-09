package simulator

import (
	"github.com/pirosuke/exercise-simulator/internal/models"
)

func calcExerciseResult(startWeight float64, weeksToOutput int64, weeklyPlans models.WeeklyPlans) models.WeeklyResult {
	remainingWeight := startWeight

	var week int64
	var exerciseResults []models.ExerciseResult
	for week = 1; week <= weeksToOutput; week++ {
		weeklyResult := models.ExerciseResult{
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

	weeklyResult := models.WeeklyResult{
		Name:       weeklyPlans.Name,
		EndWeight:  remainingWeight,
		RemovedFat: startWeight - remainingWeight,
		Results:    exerciseResults,
	}

	return weeklyResult
}

/*
SimulatePlans returns simulated result of specified plans.
*/
func SimulatePlans(plansConfig *models.PlansConfig, planTitle string) models.PlansSimulatedResult {
	result := models.PlansSimulatedResult{}
	result.Title = planTitle
	result.StartWeight = plansConfig.Weight
	result.Weeks = plansConfig.WeeksToOutput

	var weeklyResults []models.WeeklyResult
	for _, weeklyPlans := range plansConfig.WeeklyPlans {
		weeklyResult := calcExerciseResult(plansConfig.Weight, plansConfig.WeeksToOutput, weeklyPlans)
		weeklyResults = append(weeklyResults, weeklyResult)
	}

	result.WeeklyResults = weeklyResults

	return result
}
