package Exercise

type Exercise struct {
	Name      string  `json:"Name"`
	Reps      int     `json:"Reps"`
	Sets      int     `json:"Sets"`
	IsEXM     float32 `json:"isEXM"`
	ShouldEXM float32 `json:"shouldEXM"`
	Load      float32 `json:"load"`
}

func NewExercise(Name string, Reps int, Sets int, IsEXM float32, ShouldEXM float32, Load float32) *Exercise {
	return &Exercise{Name, Reps, Sets, IsEXM, ShouldEXM, Load}
}
