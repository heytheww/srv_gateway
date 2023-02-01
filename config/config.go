package config

// 配置文件
// todo 解决配置文件相对路径问题

type Config struct {
	ReqHost       string `json:"reqHost"`       // 目标请求地址
	Scheme        string `json:"scheme"`        // 请求的协议
	ReqPath       string `json:"reqPath"`       // 目标请求路径
	ForbiddenPath string `json:"forbiddenPath"` // 拦截请求后的转发地址
	GListenPort   string `json:"GListenPort"`   // 网关监听端口
	LoginPath     string `json:"LoginPath"`     // 验证用户登录接口
}

var config Config = Config{
	ReqHost:       "localhost:1234",
	Scheme:        "http",
	ReqPath:       "/buy",
	ForbiddenPath: "/403",
	GListenPort:   "8080",
	LoginPath:     "http://localhost:1234/login",
}

func (c *Config) GetConfig() *Config {

	return &config
}
