package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"srv_gateway/jwt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v4"
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
	auth := req.Header["Auth"]
	var token string
	var flag bool

	if len(auth) > 0 {
		token = auth[0]
	}

	// token为空，尝试获取openid，签发token
	if len(token) == 0 {
		user := req.FormValue("user")
		pass := req.FormValue("password")

		urlValues := url.Values{}
		urlValues.Add("user", user)
		urlValues.Add("password", pass)

		// 验证用户名密码，返回登录状态
		resp, err := http.PostForm("http://localhost:1234/login", urlValues)
		if err != nil {
			panic(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)

		lr := LoginResp{}
		err = json.Unmarshal(body, &lr)
		if err != nil {
			panic(err)
		}

		// 登录失败，应拦截请求
		if lr.Status == http.StatusForbidden {
			// 验证失败，返回403
			return http.StatusForbidden, ""
		} else {
			j := jwt.JWT{
				HMACSecret: a.HMACSecret, // 秘钥
			}

			// 偷懒格式化
			claims := &jwt.MyClaims{
				user,
				gojwt.RegisteredClaims{
					ExpiresAt: gojwt.NewNumericDate(time.Now().Add(3 * time.Minute)), // 过期时间
					IssuedAt:  gojwt.NewNumericDate(time.Now()),                      // 签发时间
					NotBefore: gojwt.NewNumericDate(time.Now()),                      // 生效时间
					Issuer:    "admin",                                               // 签发人
					Subject:   "loginByPwd",                                          // 主题
					ID:        "-",
					Audience:  []string{"buyer"}, // 受众
				},
			}

			newToken, err2 := j.Sign(claims)
			if err2 != nil {
				panic(err2)
			}
			token = newToken

			return http.StatusOK, token
		}

	}

	// token不为空，验证token有效性
	j := jwt.JWT{}
	flag = j.Validate(token)

	if flag {
		// 验证成功，返回200
		return http.StatusOK, token
	} else {
		// 验证失败，返回403
		return http.StatusForbidden, token
	}
}
