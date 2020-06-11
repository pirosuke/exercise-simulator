# exercise-simulator
Exercise Diet Plan Simulator.

## Usage

Create result JSON from plans config.

```
exersim [flags]
  -p string
        Plans file path
```

Create plot image from result JSON.

```
exersim_plot [flags]
  -f string
        exersim result file path
```

## Plans File Example

```json
{
    "weight": 80.0,
    "weeks_to_output": 10,
    "weekly_plans": [
        {
            "name": "1jog",
            "plans": [
                {
                    "title": "jogging",
                    "mets": 7,
                    "hours": 1,
                    "times_per_week": 1
                }
            ]
        },
        {
            "name": "1jog_2dance",
            "plans": [
                {
                    "title": "jogging",
                    "mets": 7,
                    "hours": 1,
                    "times_per_week": 1
                },
                {
                    "title": "dancing",
                    "mets": 4.8,
                    "hours": 1,
                    "times_per_week": 2
                }
            ]
        },
        {
            "name": "1jog_3dance",
            "plans": [
                {
                    "title": "jogging",
                    "mets": 7,
                    "hours": 1,
                    "times_per_week": 1
                },
                {
                    "title": "dancing",
                    "mets": 4.8,
                    "hours": 1,
                    "times_per_week": 3
                }
            ]
        }
    ]
}
```

## Result File Example

```json
{
    "title": "plans_sample.json",
    "weeks": 10,
    "start_weight_kg": 80,
    "weekly_results": [
        {
            "name": "1jog",
            "end_weight_kg": 79.18707470146855,
            "removed_fat_kg": 0.8129252985314537,
            "results": [
                {
                    "week": 1,
                    "start_weight_kg": 80,
                    "exercise_kcal": 588,
                    "removed_fat_kg": 0.08166666666666665
                },
                {
                    "week": 2,
                    "start_weight_kg": 79.91833333333334,
                    "exercise_kcal": 587.39975,
                    "removed_fat_kg": 0.08158329861111112
                },
                ...
                {
                    "week": 10,
                    "start_weight_kg": 79.26799411212467,
                    "exercise_kcal": 582.6197567241163,
                    "removed_fat_kg": 0.08091941065612727
                }
            ]
        },
        {
            "name": "1jog_2dance",
            "end_weight_kg": 78.08429527466448,
            "removed_fat_kg": 1.915704725335516,
            "results": [
                {
                    "week": 1,
                    "start_weight_kg": 80,
                    "exercise_kcal": 1394.4,
                    "removed_fat_kg": 0.19366666666666665
                },
                {
                    "week": 2,
                    "start_weight_kg": 79.80633333333333,
                    "exercise_kcal": 1391.02439,
                    "removed_fat_kg": 0.1931978319444444
                },
                ...
                {
                    "week": 10,
                    "start_weight_kg": 78.27378305781694,
                    "exercise_kcal": 1364.3120386977494,
                    "removed_fat_kg": 0.1894877831524652
                }
            ]
        },
        {
            "name": "1jog_3dance",
            "end_weight_kg": 77.53810568508884,
            "removed_fat_kg": 2.461894314911163,
            "results": [
                {
                    "week": 1,
                    "start_weight_kg": 80,
                    "exercise_kcal": 1797.6,
                    "removed_fat_kg": 0.24966666666666665
                },
                {
                    "week": 2,
                    "start_weight_kg": 79.75033333333333,
                    "exercise_kcal": 1791.98999,
                    "removed_fat_kg": 0.2488874986111111
                },
                ...
                {
                    "week": 10,
                    "start_weight_kg": 77.78084674430335,
                    "exercise_kcal": 1747.7356263444963,
                    "removed_fat_kg": 0.2427410592145134
                }
            ]
        }
    ]
}
```
