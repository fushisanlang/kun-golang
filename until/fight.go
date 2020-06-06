package until

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//输入概率摆动参数addpart，随机生成概率。如果概率大约50+addpart，则获胜，返回true，反之失败，返回false
func Fight(addpart int) int {
	rand.Seed(time.Now().Unix())
	probability := rand.Intn(100)
	rule := 50 + addpart
	if probability < rule {
		return 0
	}
	return 1
}

//属性克制函数。通过输入五行属性，返回相应概率摆动参数 传参加1为了与属性表对应，1是金，属性表中为0
func AddPartGet(myattributes int, enemyattributes int) int {
	ma := myattributes + 1
	ea := enemyattributes + 1
	//根据平衡性调整
	addpart := 30
	basepart := 0
	switch {
	case ma == 1 && ea == 7 || ea == 3 || ea == 6 || ea == 4:
		basepart = basepart - addpart
	case ma == 1 && ea == 2 || ea == 5:
		basepart = basepart + addpart
	case ma == 2 || ea == 4 || ea == 6 || ea == 1:
		basepart = basepart - addpart
	case ma == 2 && ea == 5 || ea == 3:
		basepart = basepart + addpart
	case ma == 3 && ea == 2 || ea == 5:
		basepart = basepart - addpart
	case ma == 3 && ea == 1 || ea == 4:
		basepart = basepart + addpart
	case ma == 4 && ea == 7 || ea == 3 || ea == 5:
		basepart = basepart - addpart
	case ma == 4 && ea == 2 || ea == 1:
		basepart = basepart + addpart
	case ma == 5 && ea == 1 || ea == 2:
		basepart = basepart - addpart
	case ma == 5 && ea == 4 || ea == 3:
		basepart = basepart + addpart
	case ma == 6 && ea == 2 || ea == 1:
		basepart = basepart + addpart
	case ma == 7 && ea == 1 || ea == 2:
		basepart = basepart + addpart
	}
	return basepart
}

func FindEnemy() {
	for {
		PrintSwitch := GetPrintStatus()
		if PrintSwitch == "on" {

			time.Sleep(1 * time.Second)
			Clear()
			OwnStruct := GetOwnStatus()
			if OwnStruct.Grade == 9 {
				panic("game over")
			}
			OwnStatus := "你的" /*+ OwnStruct.Name */ + "现在状态为：\n" + "\t级别：" + strconv.Itoa(OwnStruct.Grade) + "\n\t等级：" + strconv.Itoa(OwnStruct.Level) + "\n\t经验：" + strconv.Itoa(OwnStruct.Experience) + "% \n"
			fmt.Println(OwnStatus)
		} else {
			Clear()
			SetPrintStatus("on")
			EnemyStruct := GetEnemyStatus()
			EnemyStatus := "发现敌人" /*+ EnemyStruct.Name */ + "：\n" + "\t级别：" + strconv.Itoa(EnemyStruct.Grade) + "\n\t等级：" + strconv.Itoa(EnemyStruct.Level) + "\n"
			OwnStruct := GetOwnStatus()
			OwnStatus := "你的" /*+ OwnStruct.Name */ + "现在状态为：\n" + "\t级别：" + strconv.Itoa(OwnStruct.Grade) + "\n\t等级：" + strconv.Itoa(OwnStruct.Level) + "\n\t经验：" + strconv.Itoa(OwnStruct.Experience) + "% \n"
			fmt.Println(EnemyStatus)
			fmt.Println()
			fmt.Println()
			fmt.Println(OwnStatus)
			BasePart := AddPartGet(OwnStruct.Treasure, EnemyStruct.Treasure)
			WinPart := 50 + BasePart
			fmt.Println("获胜概率为", WinPart)
			fmt.Println("逃跑扣50%经验，经验不足50%则掉级。")
			fmt.Println("战败掉", EnemyStruct.Grade, "级别。")
			fmt.Println("战胜增加", EnemyStruct.Grade, "级别和", EnemyStruct.Level, "经验")
			fmt.Println("输入y迎战，其他按键逃跑。")
			time.Sleep(1 * time.Second)
			//var fightselect string
			//fmt.Scanf("%d", &fightselect)
			//		if fightselect == "y" {
			if Fight(WinPart) == 1 {
				fmt.Println("win")
				OwnExAdd(OwnStruct, EnemyStruct.Level, EnemyStruct.Grade, 0)
				fmt.Println("winok")
			} else {
				fmt.Println("shibai")
				OwnExAdd(OwnStruct, 0, -1, 0)
			}
			//		} else {
			//			OwnExAdd(OwnStruct, -50, 0, 0)
			//		}
		}
	}

}
