package User

import "mpm_server/data_structure/Trainingsplan"

type User struct {
	Name              string
	TrainingsplanList []Trainingsplan.Trainingsplan
}

func NewUser() *User {
	return &User{}
}
