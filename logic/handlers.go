/*
 * Author: zheng-ji.info
 */

package logic

import "net/http"
import (
	"github.com/cihub/seelog"
	"fmt"
)

// Resp Type
type Resp struct {
	Code     int    `json:"c"`
	TargetIP string `json:"targetip,omitempty"`
	Host     string `json:"host, omitempty"`
	Msg      string `json:"msg, omitempty"`
}

const (
	SUCC   = 0
	FAILED = -1
	HTTP   = "http://"
)

// PingHandler Func
func PingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("ok"))
}

// ResolveHandler Func
func ResolveHandler(writer http.ResponseWriter, request *http.Request) {

	url := request.FormValue("url")

	targetIPstr, hostStr, err := getResultFromCache(url)
	if err == nil {
		resp := Resp{
			Code:     SUCC,
			TargetIP: targetIPstr,
			Host:     hostStr,
		}
		writer.Write([]byte(resp.jsonString()))
		return
	}

	targetIP, host, err := DnsDecoder(url)
	if err != nil {
		resp := Resp{
			Code: FAILED,
			Msg:  fmt.Sprintf("%s", err),
		}
		seelog.Errorf("[ResolveHandler] error: %v", err)
		writer.Write([]byte(resp.jsonString()))
		return
	}
	resp := Resp{
		Code:     SUCC,
		TargetIP: *targetIP,
		Host:     *host,
	}
	cacheResp(url, *host, *targetIP)
	seelog.Infof("[ResolveHandler] host:%s targetIp:%s", *host, *targetIP)
	writer.Write([]byte(resp.jsonString()))
}
