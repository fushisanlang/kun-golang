package until

import (
	"math/rand"
	"time"
)

type EnemyStruct struct {
	Direction int
	Treasure  int
	Level     int
	Grade     int
}

//生成敌人
func EnemyCreate(LevelRange, GradeRange int) EnemyStruct {
	rand.Seed(time.Now().Unix())
	AttributesDirection := rand.Intn(3)
	AttributesTreasure := rand.Intn(6)
	AttributesLevel := rand.Intn(LevelRange)
	AttributesGrade := rand.Intn(GradeRange)
	return EnemyStruct{AttributesDirection, AttributesTreasure, AttributesLevel, AttributesGrade}
}
func AutoEnemyCreate() {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)
		if i%6 == 0 {
			SetPrintStatus("off")
			SetEnemyStatus(EnemyCreate(i/100, i/10000))
		}
	}
}

