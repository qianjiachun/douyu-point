package apis_common

import (
	"douyu-point/common"
	"net"
	"net/http"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

func VerifyDyToken(token string) bool {
	// 用于判断斗鱼的token是否有效
	var ret bool
	content := common.HttpPost("https://pcapi.douyucdn.cn/japi/tasksys/ytxb/box", "token="+token)
	if common.GetStrMiddle(content, `"error":`, `,`) == "0" {
		ret = true
	} else {
		ret = false
	}
	return ret
}

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
