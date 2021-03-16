package stbz

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// APIResult APIResult
type APIResult struct {
	ID   string      `json:"id"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var currentConfig *Config

// SetConfig SetConfig
func SetConfig(config *Config) {
	currentConfig = config
}

// API API
func API(method method, api string, queryMap map[string]string, bodyMap map[string]interface{}) (result *APIResult, err error) {

	header := new(signV2)
	header.AppKey = currentConfig.AccessKey

	header.TimeStamp = uint64(time.Now().UnixNano() / 1e6)
	header.Nonce = strconv.Itoa(int(time.Now().UnixNano()))

	sign := Sign(*header, queryMap, bodyMap, currentConfig.SecretKey)

	header.Sign = sign
	var response *ghttp.ClientResponse
	switch method {
	case Method.POST:
		response, err = g.Client().SetHeaderMap(header.toMap()).ContentJson().Post(currentConfig.Host+api, bodyMap)
		if nil != err {
			return
		}

	case Method.GET:
		queryString := "?"
		i := 0
		for k, v := range queryMap {
			if 0 != i {
				queryString += "&"
			}
			queryString += k + "=" + url.QueryEscape(v)
			i++
		}

		if len(queryMap) > 0 {
			response, err = g.Client().SetHeaderMap(header.toMap()).Get(currentConfig.Host+api+queryString)
		}else {
			response, err = g.Client().ContentJson().SetHeaderMap(header.toMap()).Get(currentConfig.Host+api+queryString, bodyMap)
		}

		if nil != err {
			return
		}

	case Method.PUT:
		response, err = g.Client().SetHeaderMap(header.toMap()).ContentJson().Put(currentConfig.Host+api, bodyMap)
		if nil != err {
			return
		}
	case Method.PATCH:
		response, err = g.Client().SetHeaderMap(header.toMap()).ContentJson().Patch(currentConfig.Host+api, bodyMap)
		if nil != err {
			return
		}
	case Method.DELETE:
		response, err = g.Client().SetHeaderMap(header.toMap()).ContentJson().Delete(currentConfig.Host+api, bodyMap)
		if nil != err {
			return
		}

	}

	result = new(APIResult)

	err = json.Unmarshal(response.ReadAll(), result)
	response.Close()
	return
}
