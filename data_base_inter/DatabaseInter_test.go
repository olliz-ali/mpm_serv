package databaseinter_test

import (
	databaseinter "mpm_server/data_base_inter"
	"testing"
)

func TestInsertJson(t *testing.T) {
	db := databaseinter.NewDBinter("root", "d8TM7i#%", "ai_coach_db")
	jsonstring := `{"username":"oli","gender":"male","userweight":120,"userhight":1.98,"userage":24,"num_of_days":4,"qual_of_sleep":10,"lvl_of_stress":0,"lvl_of_exp":6,"specifikation":"muskelaufbau","time_per_workout":2,"remarks":"bad knee","equipment":{"dubbells":6,"barbell":10,"cabels":4,"machines":3,"weight_belt":0},"muscle_to_train":{"chest":1,"back":1,"arms":1,"shoulders":1,"legs":1,"glutes":1,"calves":1}}`
	db.InsertJSON(jsonstring, "user")
}
