// Author: zheng-ji.info

package main

import (
	"flag"

	"net/http"

	"github.com/psy-core/goHttpDns/common"
	"github.com/psy-core/goHttpDns/logic"
)

var (
	configFile = flag.String("c", "../etc/conf.yml", "配置文件路径，默认etc/conf.yml")
)

func main() {

	flag.Parse()
	if !common.Init(*configFile) {
		return
	}

	http.HandleFunc("/ping", logic.PingHandler)
	http.HandleFunc("/d", logic.ResolveHandler)
	http.ListenAndServe(common.AppConf.Listen+":"+common.AppConf.Port, nil)
}
