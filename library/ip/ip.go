package ip

import (
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// GetIpAddr 获取本机公网ip地址
func GetIpAddr() string {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		log.GetLogger().Error("get ip addr fail", zap.Error(err))
		return ""
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.GetLogger().Error("parse resp data fail", zap.Error(err))
		return ""
	}
	return string(ipBytes)
}
