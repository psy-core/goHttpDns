/*
 * Author: zheng-ji.info
 */

package logic

import (
	"encoding/json"
	"fmt"

	"github.com/cihub/seelog"
	"github.com/psy-core/goHttpDns/common"

	"github.com/psy-core/goHttpDns/redis"
)

func (resp *Resp) jsonString() string {
	b, _ := json.Marshal(resp)
	return string(b)
}

func cacheResp(url, host, targetIP string) {

	key := fmt.Sprintf("%s_host", url)
	if err := redis.SetEx(key, host, common.AppConf.TTL); err != nil {
		seelog.Errorf("cache resp host to redis error: %v", err)
	}
	key = fmt.Sprintf("%s_ip", url)
	if err := redis.SetEx(key, targetIP, common.AppConf.TTL); err != nil {
		seelog.Errorf("cache resp ip to redis error: %v", err)
	}
}

func getResultFromCache(url string) (string, string, error) {

	key := fmt.Sprintf("%s_host", url)
	host, err := redis.Get(key)
	if err != nil {
		seelog.Errorf("get host from cache error: %v", err)
		return "", "", err
	}
	key = fmt.Sprintf("%s_ip", url)
	targetIP, err := redis.Get(key)
	if err != nil {
		seelog.Errorf("get target ip from cache error: %v", err)
		return "", "", err
	}
	return targetIP, host, nil
}
