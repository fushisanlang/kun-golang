package until

import (
	"github.com/garyburd/redigo/redis"
)

func GetPrintStatus() string {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()
	PrintSwitch, err := redis.String(RedisClient.Do("GET", "PrintSwitch"))
	CheckErr(err)
	return PrintSwitch
}
func SetPrintStatus(PrintSwitch string) {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()

	_, err = RedisClient.Do("SET", "PrintSwitch", PrintSwitch)
	CheckErr(err)

}

