/**
* @Author: fushisanlang
* @Date: 2020/5/30 1:49 下午
 */

package main

import (
	"fmt"
	"kun-golang/until"

	"github.com/gookit/color"
)

func main() {
	//设置颜色显示
	color.Set(color.FgCyan)
	//关闭颜色
	defer color.Reset()
	//创建一个用来结束游戏的通道
	GameEndChan := make(chan bool)
	//自动加经验
	go until.AutoOwnExAdd(GameEndChan)
	//遭遇敌人
	go until.FindEnemy(GameEndChan)
	//自动生成敌人
	go until.AutoEnemyCreate()
	//从结束通道中读取信息，只有在有结束信号时才结束主进程
	EndSignal := <-GameEndChan
	if EndSignal == true {
		fmt.Println("怒而飞，其翼若垂天之云。")
		fmt.Println("阁下的鲲已成鹏，游戏结束。")
	}
}
