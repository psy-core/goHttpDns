/*
 * Author: zheng-ji.info
 */

package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/cihub/seelog"
	"github.com/psy-core/goHttpDns/redis"
	goyaml "gopkg.in/yaml.v2"
)

//AppConf 对象
var AppConf AppConfig

// AppConfig Type
type AppConfig struct {
	Redis      redis.RedisConfig `yaml:"redis"`
	Logconf    string            `yaml:"log_config"`
	Listen     string            `yaml:"listen"`
	Port       string            `yaml:"port"`
	TTL        int               `yaml:"ttl"`
	Dnsservers []string          `yaml:"dnsservers"`
}

func (ac *AppConfig) isValid() bool {
	return ac.Redis.IsValid() &&
		len(ac.Listen) > 0 &&
		len(ac.Port) > 0
}

//解析配置文件
func parseConfigFile(filepath string) error {
	config, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err = goyaml.Unmarshal(config, &AppConf); err != nil {
		return err
	}
	if AppConf.TTL == 0 {
		AppConf.TTL = 10 * 60
	}
	if !AppConf.isValid() {
		return errors.New("Invalid configuration")
	}
	return nil
}

// Init Func
func Init(conf string) bool {

	//基本配置文件
	err := parseConfigFile(conf)
	if nil != err {
		fmt.Printf("init config file error: %v\n", err)
		return false
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	//seelog 配置
	logger, err := seelog.LoggerFromConfigAsFile(AppConf.Logconf)
	if nil != err {
		fmt.Printf("init seelog file error: %v\n", err)
		return false
	}
	seelog.UseLogger(logger)

	//redis 链接
	redis.Init(&AppConf.Redis)

	return true
}
