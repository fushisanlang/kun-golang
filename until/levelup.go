package until

import (
	"fmt"
	"time"
)

//升级函数，经验100生级，100级升阶，三阶后成鹏，游戏结束
func ExAdd(Experience, Level, Grade, ExBase, LvBase, GradeBase int, GameEndChan chan bool) (int, int, int) {
	Experience = Experience + ExBase
	Level = Level + LvBase
	Grade = Grade + GradeBase
	if Experience > 99 {
		fmt.Println("升级")
		Level = Level + 1
		if Level > 99 {
			fmt.Println("进化")
			Grade = Grade + 1
			if Grade > 2 {
				GameEndChan <- true
			}
			Level = Level - 100
		}
		Experience = Experience - 100
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

//调用升级函数,可与ExAdd函数合并
func OwnExAdd(OwnStructIn OwnStruct, ExBase int, LvBase int, GradeBase int, GameEndChan chan bool) OwnStruct {
	OwnDirection := OwnStructIn.Direction
	OwnTreasure := OwnStructIn.Treasure
	OwnExperience, OwnLevel, OwnGrade := ExAdd(OwnStructIn.Experience, OwnStructIn.Level, OwnStructIn.Grade, ExBase, LvBase, GradeBase, GameEndChan)

	SetOwnStatus(OwnStruct{OwnDirection, OwnTreasure, OwnExperience, OwnLevel, OwnGrade})
	return OwnStruct{OwnDirection, OwnTreasure, OwnExperience, OwnLevel, OwnGrade}
}

//自动升级，每秒加1经验
func AutoOwnExAdd(GameEndChan chan bool) {
	for {
		time.Sleep(1 * time.Second)
		OwnStatus := GetOwnStatus()
		OwnStatus = OwnExAdd(OwnStatus, 1, 0, 0, GameEndChan)
		SetOwnStatus(OwnStatus)
	}
}
