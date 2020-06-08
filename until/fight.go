package until

import (
	"math/rand"
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
