package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"srv_gateway/config"
)

// 创建反向代理处理方法
func NewReverseProxy(loginStatus int, token string, c *config.Config) *httputil.ReverseProxy {

	//创建代理处理器
	director := func(req *http.Request) {

		//查询原始请求路径
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}

		// 按照分隔符 / 对路径进行分解
		// pathArray := strings.Split(reqPath, "/")
		// serviceName := pathArray[1]

		// token验证失败 或 无法 签发token
		// 拦截请求：转到网关监听的/403路由中处理
		if loginStatus == http.StatusForbidden {
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:" + c.GListenPort
			req.URL.Path = "/403"
		} else {
			req.URL.Scheme = c.Scheme
			req.URL.Host = c.ReqHost
		}

	}

	modifyResponse := func(r *http.Response) error {
		fmt.Println(r)
		return nil
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}
