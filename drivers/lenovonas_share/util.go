package LenovoNasShare

import (
	"errors"

	"github.com/alist-org/alist/v3/drivers/base"
	"github.com/alist-org/alist/v3/pkg/utils"
	jsoniter "github.com/json-iterator/go"
)

func (d *LenovoNasShare) request(url string, method string, callback base.ReqCallback, resp interface{}) ([]byte, error) {
	req := base.RestyClient.R()
	req.SetHeaders(map[string]string{
		"origin":      "https://siot-share.lenovo.com.cn",
		"referer":     "https://siot-share.lenovo.com.cn/",
		"user-agent":  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.6045.199",
		"platform":    "web",
		"app-version": "3",
	})
	if callback != nil {
		callback(req)
	}
	if resp != nil {
		req.SetResult(resp)
	}
	res, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	body := res.Body()
	result := utils.Json.Get(body, "result").ToBool()
	if !result {
		return nil, errors.New(jsoniter.Get(body, "error", "msg").ToString())
	}
	return body, nil
}
