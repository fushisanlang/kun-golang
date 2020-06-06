/**
* @Author: fushisanlang
* @Date: 2020/5/30 2:23 下午
 */

package until

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var str string
var num, num2 int

var GradeOne = [4][7]string{
	{"亢金龙", "角木蛟", "箕水豹", "尾火虎", "氐土貉", "房日兔", "心月狐"},
	{"鬼金羊", "井木犴", "轸水蚓", "翼火蛇", "柳土獐", "星日马", "张月鹿"},
	{"娄金狗", "奎木狼", "参水猿", "嘴火猴", "胃土雉", "昴日鸡", "毕月乌"},
	{"牛金牛", "斗木獬", "壁水貐", "室火猪", "女土蝠", "虚日鼠", "危月燕"},
}
var GradeTwo = [4][7]string{
	{"困顿", "青龙", "赤奋若", "摄提格", "饕餮", "功曹", "大冲"},
	{"穷奇", "单阏", "执徐", "朱雀", "大荒落", "胜先", "太一"},
	{"白虎", "梼杌", "敦牂", "协洽", "涒滩", "传送", "徵魁"},
	{"作噩", "阉茂", "玄武", "混沌", "大渊献", "神后", "徵明"},
}

var GradeThree = [4][7]string{
	{"困顿", "青华", "赤奋若", "摄提格", "缙云", "金乌", "顾兔"},
	{"少皞", "单阏", "执徐", "长生", "大荒落", "羲轮", "桂魄"},
	{"天皇", "颛顼", "敦牂", "协洽", "涒滩", "规毁", "素娥"},
	{"作噩", "阉茂", "紫薇", "帝鸿", "大渊献", "羲御", "玄烛"},
}

type OwnStruct struct {
	Direction  int
	Treasure   int
	Experience int
	Level      int
	Grade      int
}

func init() {
	Clear()
	//	color.Set(color.FgBlack, color.BgWhite)
	//defer color.Unset()
	fmt.Println("                tunshi  KUN              ")
	fmt.Println("                                         ")
	fmt.Println("                version 0.1              ")
	fmt.Println("               power by fu13             ")
	fmt.Println("                                         ")
	fmt.Println("                                         ")
	fmt.Println("                                         ")
	fmt.Println("           push anykey to start          ")
	fmt.Scanf("%s", &str)

	var treasure, direction int

	//fmt.Println(GradeOne)
	stage1 := "你现在有一颗鲲蛋，你想在什么地方孵化？\n" +
		"1.东胜神洲\n" +
		"2.南赡部洲\n" +
		"3.西牛贺洲\n" +
		"4.北俱芦洲\n"
	Clear()
	fmt.Println(stage1)
	fmt.Scanf("%d", &num)
	switch num {
	case 1, 2, 3, 4:
		direction = num - 1

	default:
		fmt.Println("输入位置不存在，鲲蛋已进入异次元。请少侠从头来过。。。")
		time.Sleep(1 * time.Minute)
	}
	stage2 := "现在你面前有七个法宝，选择其中一个给幼鲲吞噬：\n" +
		"1. 金蛟剪\n" +
		"2. 桃木杖\n" +
		"3. 定海珠\n" +
		"4. 火龙标\n" +
		"5. 劈地珠\n" +
		"6. 九龙神火罩\n" +
		"7. 清净琉璃瓶\n"
	Clear()
	fmt.Println(stage2)
	fmt.Scanf("%d", &num2)

	switch num2 {
	case 1, 2, 3, 4, 5, 6, 7:
		treasure = num2 - 1
	default:
		fmt.Println("投喂了位置物品，幼鲲死亡。请少侠从头来过。。。")
		time.Sleep(1 * time.Minute)
	}
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()

	_, err = RedisClient.Do("MSET", "owndirection", direction, "owntreasure", treasure, "ownex", 0, "ownlv", 0, "owngrade", 0)
	CheckErr(err)
	SetPrintStatus("on")
}
func PrintEgg(grade, direction, treasure int) string {
	grade = grade + 1
	var name string
	switch {
	case grade == 1:
		name = GradeOne[direction][treasure]
	case grade == 2:
		name = GradeTwo[direction][treasure]
	case grade == 3:
		name = GradeThree[direction][treasure]
	}
	return name
}

