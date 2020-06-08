package until

import (
	"fmt"
	"math/rand"
	"strconv"
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
	if LevelRange > 100 {
		LevelRange = 100
	}
	if LevelRange < 10 {
		LevelRange = 10
	}
	if GradeRange > 2 {
		GradeRange = 2
	}
	rand.Seed(time.Now().Unix())
	AttributesDirection := rand.Intn(3)
	AttributesTreasure := rand.Intn(6)
	AttributesLevel := rand.Intn(LevelRange)
	AttributesGrade := rand.Intn(GradeRange)
	return EnemyStruct{AttributesDirection, AttributesTreasure, AttributesLevel, AttributesGrade}
}

//自动生成敌人，可与EnemyCreate函数整合
func AutoEnemyCreate() {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)
		//生成敌人的时间间隔
		if i%60 == 0 {
			LevelRange := i/100 + 1
			GradeRange := i/10000 + 1
			//关闭打印开关，用来显示敌人
			SetPrintStatus("off")
			SetEnemyStatus(EnemyCreate(LevelRange, GradeRange))
		}
	}
}

//遭遇敌人
func FindEnemy(GameEndChan chan bool) {
	for {
		//获取打印开关
		PrintSwitch := GetPrintStatus()
		//开关打开则无敌人，默认打印个人信息
		if PrintSwitch == "on" {

			time.Sleep(1 * time.Second)
			Clear()
			OwnStruct := GetOwnStatus()
			if OwnStruct.Grade == 9 {
				panic("战败，游戏结束")
			}
			OwnName := PrintEgg(OwnStruct.Grade, OwnStruct.Direction, OwnStruct.Treasure)
			OwnStatus := "你的" + OwnName + "现在状态为：\n" + "\t位阶：" + strconv.Itoa(OwnStruct.Grade) + "阶\n\t等级：" + strconv.Itoa(OwnStruct.Level) + "级\n\t经验：" + strconv.Itoa(OwnStruct.Experience) + "点 \n"
			fmt.Println(OwnStatus)
			//开关关闭 ，显示敌人信息
		} else {
			Clear()
			SetPrintStatus("on")
			EnemyStruct := GetEnemyStatus()
			EnemyName := PrintEgg(EnemyStruct.Grade, EnemyStruct.Direction, EnemyStruct.Treasure)
			EnemyStatus := "发现敌人" + EnemyName + "：\n" + "\t位阶：" + strconv.Itoa(EnemyStruct.Grade) + "阶\n\t等级：" + strconv.Itoa(EnemyStruct.Level) + "级\n"
			OwnStruct := GetOwnStatus()
			OwnName := PrintEgg(OwnStruct.Grade, OwnStruct.Direction, OwnStruct.Treasure)
			OwnStatus := "你的" + OwnName + "现在状态为：\n" + "\t位阶：" + strconv.Itoa(OwnStruct.Grade) + "阶\n\t等级：" + strconv.Itoa(OwnStruct.Level) + "级\n\t经验：" + strconv.Itoa(OwnStruct.Experience) + "点 \n"
			fmt.Println(EnemyStatus)
			fmt.Println()
			fmt.Println()
			fmt.Println(OwnStatus)
			BasePart := AddPartGet(OwnStruct.Treasure, EnemyStruct.Treasure)
			WinPart := 50 + BasePart
			fmt.Println("获胜概率为", WinPart)
			fmt.Println("逃跑扣50点经验，经验不足50点则掉级。")
			//???
			fmt.Println("战败掉", EnemyStruct.Grade, "级别。")
			fmt.Println("战胜增加", EnemyStruct.Grade, "级别和", EnemyStruct.Level, "经验")
			fmt.Println("输入y迎战，其他按键逃跑。")
			time.Sleep(1 * time.Second)
			var fightselect string
			fmt.Scanf("%d", &fightselect)
			if fightselect == "y" {
				if Fight(WinPart) == 1 {
					fmt.Println("胜利")
					OwnExAdd(OwnStruct, EnemyStruct.Level, EnemyStruct.Grade, 0, GameEndChan)
				} else {
					fmt.Println("战败")
					//??
					OwnExAdd(OwnStruct, 0, -1, 0, GameEndChan)
				}
			} else {
				OwnExAdd(OwnStruct, -50, 0, 0, GameEndChan)
			}
		}
	}

}
