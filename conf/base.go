package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Cfg *goconfig.ConfigFile

//加载配置文件
func init() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	Check("无法加载配置文件:%s", err)
	Cfg = cfg
}

func Check(info string, err error) {
	if err != nil {
		log.Fatalf(info, err)
	}
}
