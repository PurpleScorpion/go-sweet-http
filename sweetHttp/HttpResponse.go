package sweetHttp

type HttpResponse struct {
	HttpCode int
	Body     []byte
	Error    string
}
