package until

import (
	"github.com/garyburd/redigo/redis"
)

func GetOwnStatus() OwnStruct {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()
	owndirection, err := redis.String(RedisClient.Do("GET", "owndirection"))
	CheckErr(err)
	owndirectionint := StrToInt(owndirection)
	owntreasure, err := redis.String(RedisClient.Do("GET", "owntreasure"))
	CheckErr(err)
	owntreasureint := StrToInt(owntreasure)
	ownlv, err := redis.String(RedisClient.Do("GET", "ownlv"))
	CheckErr(err)
	ownlvint := StrToInt(ownlv)
	ownex, err := redis.String(RedisClient.Do("GET", "ownex"))
	CheckErr(err)
	ownexint := StrToInt(ownex)
	owngrade, err := redis.String(RedisClient.Do("GET", "owngrade"))
	CheckErr(err)
	owngradeint := StrToInt(owngrade)
	return OwnStruct{owndirectionint, owntreasureint, ownexint, ownlvint, owngradeint}
}
func SetOwnStatus(OwnStatus OwnStruct) {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()

	_, err = RedisClient.Do("MSET", "owndirection", OwnStatus.Direction, "owntreasure", OwnStatus.Treasure, "ownex", OwnStatus.Experience, "ownlv", OwnStatus.Level, "owngrade", OwnStatus.Grade)
	CheckErr(err)

}

