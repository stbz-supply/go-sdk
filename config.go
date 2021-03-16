package stbz

// Config Config
type Config struct {
	AccessKey 	string
	SecretKey 	string
	Host		string
}

type stbzApi struct{
	ApiHost   	string
	ShopApiHost string
}

// Method Method
var StbzApi = new(stbzApi)

// NewConfig NewConfig
func NewConfig(host,accessKey ,secretKey string) *Config {
	var config = new(Config)
	config.AccessKey = accessKey
	config.SecretKey = secretKey
	config.Host = host
	return config
}

func init() {
	StbzApi.ApiHost = "https://api.jxhh.com"
	StbzApi.ShopApiHost = "http://shopapi.jxhh.com"
}


