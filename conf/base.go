package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Cfg *goconfig.ConfigFile
var HttpWeb string

//加载配置文件
func init() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	Check("无法加载配置文件:%s", err)
	Cfg = cfg
	httpport, _ := cfg.GetValue("", "httpport")
	HttpWeb, _ = cfg.GetValue("", "httpweb")
	HttpWeb = HttpWeb + httpport
}

func Check(info string, err error) {
	if err != nil {
		log.Fatalf(info, err)
	}
}
