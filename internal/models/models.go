package models

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
