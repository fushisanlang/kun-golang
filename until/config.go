package until

import (
	"flag"
	"log"
	"runtime"

	"github.com/larspensjo/config"
)

//配置文件位置
var (
	configFile = flag.String("configfile", "./config.ini", "General configuration file")
)

var TOPIC = make(map[string]string)
var parameter string

func readconf(parameter string) string {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	if cfg.HasSection("kun") {
		section, err := cfg.SectionOptions("kun")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("kun", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	return TOPIC[parameter]
}
