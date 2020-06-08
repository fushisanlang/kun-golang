package until

//返回redis连接信息
func RedisInfo() string {
	ServerIp := readconf("ServerIp")
	ServerPort := readconf("ServerPort")
	ReturnStr := ServerIp + ":" + ServerPort
	return ReturnStr
}
