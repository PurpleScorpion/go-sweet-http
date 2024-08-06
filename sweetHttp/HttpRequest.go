package sweetHttp

import (
	"bytes"
	"encoding/json"
	"github.com/PurpleScorpion/go-sweet-http/logger"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var logFlag = false

func OpenLog() {
	logFlag = true
}

func Post4FormData(url string, parma url.Values, headers map[string]string) HttpResponse {
	reqCode, _ := generateRandomString(16)

	//转义参数
	data := parma.Encode()
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_CREATION_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // 设置自定义的 Header
	if logFlag {
		logger.Info("-------------------------Post Req----------------------------")
		logger.Info("reqCode: %s", reqCode)
		logger.Info("url: %s", url)
		logger.Info("data: %s", data)
		logger.Info("-------------------------------------------------------------")
	}
	return do(req, reqCode)
}

func Post(url string, parma interface{}, headers map[string]string) HttpResponse {
	reqCode, _ := generateRandomString(16)
	var payload []byte
	if s, ok := parma.(string); ok {
		payload = []byte(s)
	} else {
		payload1, err1 := json.Marshal(parma)
		if err1 != nil {
			return HttpResponse{
				HttpCode: HTTP_BODY_READ_FAILED,
				Body:     nil,
				Error:    err1.Error(),
			}
		}
		payload = payload1
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_CREATION_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	req.Header.Set("Content-Type", "application/json") // 设置自定义的 Header

	if logFlag {
		logger.Info("-------------------------Post Req----------------------------")
		logger.Info("reqCode: %s", reqCode)
		logger.Info("url: %s", url)
		logger.Info("data: %s", string(payload))
		logger.Info("-------------------------------------------------------------")
	}
	return do(req, reqCode)
}

func Get(url string, headers map[string]string) HttpResponse {
	reqCode, _ := generateRandomString(16)

	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_CREATION_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	if logFlag {
		logger.Info("-------------------------Get Req----------------------------")
		logger.Info("reqCode: %s", reqCode)
		logger.Info("url: %s", url)
		logger.Info("-------------------------------------------------------------")
	}
	return do(req, reqCode)
}

func do(req *http.Request, reqCode string) HttpResponse {
	client := &http.Client{}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_SEND_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	// 预关闭资源
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return HttpResponse{
			HttpCode: resp.StatusCode,
			Body:     nil,
			Error:    "HTTP request failed",
		}
	}
	bytes, err := io.ReadAll(resp.Body)

	if logFlag {
		logger.Info("-------------------------" + req.Method + " Resp----------------------------")
		logger.Info("respCode: %s", reqCode)
		logger.Info("status: %d", resp.StatusCode)
		logger.Info("body: %s", string(bytes))
		logger.Info("--------------------------------------------------------------")
	}

	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_BODY_READ_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	return HttpResponse{
		HttpCode: resp.StatusCode,
		Body:     bytes,
		Error:    "",
	}
}
