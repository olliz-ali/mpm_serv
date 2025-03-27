package Workout

import "mpm_server/data_structure/Exercise"

type Workout struct {
	Exercise_list []Exercise.Exercise
}

func (w *Workout) AddExercise(Name string, Reps int, Sets int, IsEXM float32, ShouldEXM float32, Load float32) {
	newExe := Exercise.NewExercise(Name, Reps, Sets, IsEXM, ShouldEXM, Load)
	w.Exercise_list = append(w.Exercise_list, *newExe)
}

func NewWorkout() *Workout {
	return &Workout{}
}
