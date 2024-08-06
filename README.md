# go-sweet-http
基于go语言的http/https请求服务工具

## 适用版本 go 1.20
## 使用方法
1. 引入包
```text
go get github.com/PurpleScorpion/go-sweet-http
```
2. 方法介绍
```text
    1. Http请求
        1.1 Get(url string, headers map[string]string)
            释义: Http的Get请求
                第一个参数是Url地址
                第二个参数是自定义Header
        1.2 Post(url string, parma interface{}, headers map[string]string)
            释义: Http的Post请求-Json格式
                第一个参数是Url地址
                第二个参数是请求参数,请求参数可以是已经解析好的String,也可以是对象,方法中会自动转为json字符串
                第三个参数是自定义Header
        1.3 Post4FormData(url string, parma url.Values, headers map[string]string)
            释义: Http的Post请求-FormData格式
                第一个参数是Url地址
                第二个参数是请求参数,请求参数必须是url.Values类型
                第三个参数是自定义Header
        
    2. Https请求
        2.1 HttpsPost(tlsConfig *tls.Config, url string, parma interface{}, headers map[string]string)
            释义: Https的Post请求-Json格式
                第一个参数是TLS配置
                第二个参数是Url地址
                第三个参数是请求参数,请求参数可以是已经解析好的String,也可以是对象,方法中会自动转为json字符串
                第四个参数是自定义Header
        2.2 HttpsPost4FormData(tlsConfig *tls.Config, url string, parma url.Values, headers map[string]string)
            释义: Http的Post请求-FormData格式
                第一个参数是TLS配置
                第二个参数是Url地址
                第三个参数是请求参数,请求参数必须是url.Values类型
                第四个参数是自定义Header
    3. 公共方法
        3.1 OpenLog()
            释义: 打开日志
        3.2 GetUrlValues(data map[string]string)
            释义: 将map[string]string转换为url.Values, FormData请求参数需要使用
        3.3 DefaultTlsConfig(pemPath string)
            释义: 获取默认的TLS配置, 需传入cacert.pem证书路径, 该返回值为GO的标准TLS配置,所以也可以自行创建后使用
        
```