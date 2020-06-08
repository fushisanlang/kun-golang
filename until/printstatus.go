package until

import (
	"github.com/garyburd/redigo/redis"
)

//判断redis中，打印函数
func GetPrintStatus() string {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()
	PrintSwitch, err := redis.String(RedisClient.Do("GET", "PrintSwitch"))
	CheckErr(err)
	return PrintSwitch
}

//设置打印参数,用来控制屏显个人信息
func SetPrintStatus(PrintSwitch string) {
	RedisClient, err := redis.Dial("tcp", RedisInfo())
	CheckErr(err)
	defer RedisClient.Close()

	_, err = RedisClient.Do("SET", "PrintSwitch", PrintSwitch)
	CheckErr(err)

}
