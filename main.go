package main

import (
	"encoding/json"
	"log"
	"net/http"
	"srv_gateway/proxy"
	"srv_gateway/seelog"
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

func main() {

	sg := seelog.SeeLog{FileName: "./conf/seelog.xml"}
	slog, err := sg.NewLog()
	if err != nil {
		log.Fatal(err)
	}
	defer slog.Flush()

	proxy := proxy.NewReverseProxy()
	http.Handle("/", proxy)
	http.HandleFunc("/403", Handle403)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error(err.Error())
	}
}
