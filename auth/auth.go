package auth

import (
	"net/http"
	"srv_gateway/jwt"
)

// 在这里进行业务授权和鉴权

type Auth struct {
	HMACSecret []byte
}

type LoginResp struct {
	User   string `json:"user"`
	Pass   string `json:"password"`
	Status int    `json:"status"`
}

func (a *Auth) DoAuth(req *http.Request) (int, string) {

	// 获取header
	auth := req.Header.Get("Auth")
	var token string
	var flag bool

	if len(auth) > 0 {
		token = auth
	}

	// 登录、注册放行
	if req.URL.Path == "/login" || req.URL.Path == "/register" {
		return http.StatusOK, ""
	}

	// token为空
	if len(token) == 0 {
		// 验证失败，返回403
		return http.StatusForbidden, ""
	}

	// token不为空，验证token有效性
	j := jwt.JWT{HMACSecret: a.HMACSecret}
	flag = j.Validate(token)

	if flag {
		// 验证成功，返回200
		return http.StatusOK, token
	} else {
		// 验证失败，返回403
		return http.StatusForbidden, token
	}
}
