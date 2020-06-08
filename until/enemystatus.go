package until

import (
	"github.com/garyburd/redigo/redis"
)

//用于设置和查看敌人参数
func GetEnemyStatus() EnemyStruct {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()
	enemydirection, err := redis.String(RedisClient.Do("GET", "enemydirection"))
	CheckErr(err)
	enemydirectionint := StrToInt(enemydirection)
	enemytreasure, err := redis.String(RedisClient.Do("GET", "enemytreasure"))
	CheckErr(err)
	enemytreasureint := StrToInt(enemytreasure)
	enemylv, err := redis.String(RedisClient.Do("GET", "enemylv"))
	CheckErr(err)
	enemylvint := StrToInt(enemylv)
	enemygrade, err := redis.String(RedisClient.Do("GET", "enemygrade"))
	CheckErr(err)
	enemygradeint := StrToInt(enemygrade)
	return EnemyStruct{enemydirectionint, enemytreasureint, enemylvint, enemygradeint}
}
func SetEnemyStatus(EnemyStatus EnemyStruct) {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()

	_, err = RedisClient.Do("MSET", "enemydirection", EnemyStatus.Direction, "enemytreasure", EnemyStatus.Treasure, "enemylv", EnemyStatus.Level, "enemygrade", EnemyStatus.Grade)
	CheckErr(err)

}
