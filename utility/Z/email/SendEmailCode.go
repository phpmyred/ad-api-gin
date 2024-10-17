package email

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendEmailCode(to, code string) {
	var apiKey = viper.GetString("email.apiKey")
	var tempId = viper.GetString("email.tempId")
	var apiUrl = viper.GetString("email.apiUrl")

	urlStr := fmt.Sprintf(`%s?app_key=%s&template_id=%s&to=%s&data={"code":"%s"}`, apiUrl, apiKey, tempId, to, code)

	method := "POST"

	client := &http.Client{}
	if viper.GetString("proxy.url") != "" && viper.GetString("proxy.port") != "" {
		proxy, err := url.Parse(viper.GetString("proxy.url") + viper.GetString("proxy.port"))
		if err != nil {
			panic("解析代理 URL 失败!")
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}

	req, err := http.NewRequest(method, urlStr, nil)

	if err != nil {
		panic("验证码发送失败!")
	}

	res, err := client.Do(req)
	if err != nil {
		panic("验证码发送失败!")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("验证码发送失败!")
	}
	var resData resData
	err = json.Unmarshal(body, &resData)
	if err != nil {
		panic("验证码发送失败!")
	}
	if resData.Code != 200 {
		panic(resData.Message)
	}
}

type resData struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
