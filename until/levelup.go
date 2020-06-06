package until

import (
	"time"
)

func ExAdd(Experience, Level, Grade, ExBase, LvBase, GradeBase int) (int, int, int) {
	Experience = Experience + ExBase
	Level = Level + LvBase
	Grade = Grade + GradeBase
	if Experience > 99 {
		Level = Level + 1
		if Level > 99 {
			Grade = Grade + 1
			if Grade > 2 {
				return 100, 100, 2
			}
			Level = 0
		}
		Experience = 0
	}
	for Experience < 0 {
		Level = Level - 1
		Experience = Experience + 100
	}
	for Level < 0 {
		Level = Level + 100
		Grade = Grade - 1
	}
	if Grade < 0 {
		return 9, 9, 9
	}
	return Experience, Level, Grade
}
func OwnExAdd(OwnStructIn OwnStruct, ExBase int, LvBase int, GradeBase int) OwnStruct {
	OwnDirection := OwnStructIn.Direction
	OwnTreasure := OwnStructIn.Treasure
	OwnExperience, OwnLevel, OwnGrade := ExAdd(OwnStructIn.Experience, OwnStructIn.Level, OwnStructIn.Grade, ExBase, LvBase, GradeBase)
	//	OwnName := PrintEgg(OwnGrade, OwnDirection, OwnTreasure)

	SetOwnStatus(OwnStruct{OwnDirection, OwnTreasure, OwnExperience, OwnLevel, OwnGrade})
	return OwnStruct{OwnDirection, OwnTreasure, OwnExperience, OwnLevel, OwnGrade}
}
func AutoOwnExAdd() {
	for {
		time.Sleep(1 * time.Second)
		OwnStatus := GetOwnStatus()
		OwnStatus = OwnExAdd(OwnStatus, 1, 0, 0)
		SetOwnStatus(OwnStatus)
	}
}

