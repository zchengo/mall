package config

// 组合全部配置模型
type Config struct {
	Server       Server       `mapstructure:"server"`
	Mysql        Mysql        `mapstructure:"mysql"`
	Upload       Upload       `mapstructure:"upload"`
	Jwt          Jwt          `mapstructure:"jwt"`
	Code2Session Code2Session `mapstructure:"code2Session"`
	Feedback     Feedback     `mapstructure:"feedback"`
}

// 服务启动端口号配置
type Server struct {
	Post string `mapstructure:"post"`
}

// MySQL数据源配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

// 文件上传相关路径配置
type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

// 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}

// 微信小程序相关配置
type Code2Session struct {
	Code      string `mapstructure:"code"`
	AppId     string `mapstructure:"appId"`
	AppSecret string `mapstructure:"appSecret"`
}

// 问题反馈配置
type Feedback struct {
	QqSmtp        string `mapstructure:"qqSmtp"`
	QqEmail       string `mapstructure:"qqEmail"`
	QqEmailSecret string `mapstructure:"qqEmailSecret"`
}
