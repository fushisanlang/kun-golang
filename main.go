/**
* @Author: fushisanlang
* @Date: 2020/5/30 1:49 下午
 */

package main

import (
	"kun-golang/until"
	"time"
)

func main() {
	//	StartOwnStruct := egg.GetAnEgg()
	go until.AutoOwnExAdd()
	go until.FindEnemy()
	go until.AutoEnemyCreate()

	time.Sleep(1 * time.Hour)

}
