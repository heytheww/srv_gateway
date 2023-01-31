package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"srv_gateway/auth"
)

// 创建反向代理处理方法
func NewReverseProxy() *httputil.ReverseProxy {

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

		a := auth.Auth{HMACSecret: []byte("abcdefg123456")}
		loginStatus, token := a.DoAuth(req)
		fmt.Println("loginStatus:", loginStatus)

		//设置代理服务地址信息
		type Config struct {
			ProxyTarget   string `json:"ProxyTarget"`
			Scheme        string `json:"Scheme"`
			Path          string `json:"Path"`
			ForbiddenPath string `json:"ForbiddenPath"`
		}

		c := Config{}

		f, err2 := os.Open("conf/config.json")
		if err2 != nil {
			panic(err2)
		}

		data, err3 := ioutil.ReadAll(f)
		if err3 != nil {
			panic(err3)
		}

		err4 := json.Unmarshal(data, &c)
		if err4 != nil {
			panic(err4)
		}
		fmt.Println(c)

		// token验证失败 或 无法 签发token
		// 拦截请求：转到网关监听的/403路由中处理
		if loginStatus == http.StatusForbidden {
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:8080"
			req.URL.Path = "/403"
		} else {
			req.URL.Scheme = c.Scheme
			req.URL.Host = c.ProxyTarget
			req.URL.Path = c.Path
			req.Header.Set("auth", token)
		}
	}

	modifyResponse := func(*http.Response) error {
	https: //pandaychen.github.io/2021/07/01/GOLANG-REVERSEPROXY-LIB-ANALYSIS/
		return nil
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}
