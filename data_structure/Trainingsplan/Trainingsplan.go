package Trainingsplan

import "mpm_server/data_structure/Workout"

type Trainingsplan struct {
	Workout_list []Workout.Workout
}

func (t *Trainingsplan) AddWorkout() {
	newWork := Workout.NewWorkout()
	t.Workout_list = append(t.Workout_list, *newWork)
}

func NewTrainingsplan() *Trainingsplan {
	return &Trainingsplan{}
}
