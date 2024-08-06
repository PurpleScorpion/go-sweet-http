package sweetHttp

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/PurpleScorpion/go-sweet-http/logger"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
使用默认配置的Tls配置
cacert.pem
*/
func DefaultTlsConfig(pemPath string) *tls.Config {
	// 加载根证书
	caCert, err := os.ReadFile(pemPath)
	if err != nil {
		logger.Error("Failed to read root certificate: %v\n", err)
		return nil
	}

	// 创建根证书池
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		logger.Error("Failed to append root certificate to pool")
		return nil
	}

	// 创建自定义的 TLS 配置
	tlsConfig := &tls.Config{
		RootCAs:    caCertPool,
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		},
		InsecureSkipVerify: true,
	}
	return tlsConfig
}

func HttpsPost4FormData(tlsConfig *tls.Config, url string, parma url.Values, headers map[string]string) HttpResponse {
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
	// 创建 HTTP 客户端并配置使用自定义的 TLS 配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := client.Do(req)
	// 预关闭资源
	defer resp.Body.Close()
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_SEND_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	if resp.StatusCode != 200 {
		return HttpResponse{
			HttpCode: resp.StatusCode,
			Body:     nil,
			Error:    "HTTP request failed",
		}
	}
	bytes, err := io.ReadAll(resp.Body)

	if logFlag {
		logger.Info("-------------------------Post Resp----------------------------")
		logger.Info("respCode: %s", reqCode)
		logger.Info("status: %d", resp.StatusCode)
		logger.Info("body: %s", string(bytes))
		logger.Info("--------------------------------------------------------- ----")
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

func HttpsPost(tlsConfig *tls.Config, url string, parma interface{}, headers map[string]string) HttpResponse {
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
	// 创建 HTTP 客户端并配置使用自定义的 TLS 配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := client.Do(req)
	// 预关闭资源
	defer resp.Body.Close()
	if err != nil {
		return HttpResponse{
			HttpCode: HTTP_REQUEST_SEND_FAILED,
			Body:     nil,
			Error:    err.Error(),
		}
	}
	if resp.StatusCode != 200 {
		return HttpResponse{
			HttpCode: resp.StatusCode,
			Body:     nil,
			Error:    "HTTP request failed",
		}
	}
	bytes, err := io.ReadAll(resp.Body)

	if logFlag {
		logger.Info("-------------------------Post Resp----------------------------")
		logger.Info("respCode: %s", reqCode)
		logger.Info("status: %d", resp.StatusCode)
		logger.Info("body: %s", string(bytes))
		logger.Info("--------------------------------------------------------- ----")
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

func GetAndEncodeUrlValues(data map[string]string) string {
	return GetUrlValues(data).Encode()
}

func GetUrlValues(data map[string]string) url.Values {
	vals := url.Values{}
	for key, value := range data {
		vals.Add(key, value)
	}
	return vals
}

func EncodeUrlValues(data url.Values) string {
	return data.Encode()
}
