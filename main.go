package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"srv_gateway/auth"
	"srv_gateway/config"
	"srv_gateway/proxy"
)

type ErrMsg struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
}

func Handle403(resp http.ResponseWriter, req *http.Request) {
	v := ErrMsg{
		Code:    403,
		Massage: "请登录",
	}

	vj, _ := json.Marshal(v)

	resp.Write(vj)
}

// 代理的director内不能发送请求，因此在这里处理token验证，
// 然后需要转发到哪，再传给ReverseProxy执行
func handleProxy(resp http.ResponseWriter, req *http.Request) {

	a := auth.Auth{HMACSecret: []byte("abc123")}
	loginStatus, token := a.DoAuth(req)
	fmt.Println("loginStatus:", loginStatus)

	//设置代理服务配置信息
	c := config.Config{}
	conf := c.GetConfig()

	// 执行ReverseProxy
	p := proxy.NewReverseProxy(loginStatus, token, conf)
	p.ServeHTTP(resp, req)
}

func main() {

	http.HandleFunc("/", handleProxy)
	http.HandleFunc("/auth", handleProxy)
	http.HandleFunc("/403", Handle403)

	c := config.Config{}
	conf := c.GetConfig()

	err := http.ListenAndServe(":"+conf.GListenPort, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
