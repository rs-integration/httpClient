package httpClient

const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
)

type HttpClient struct {
	method  string
	url     string
	data    string
	headers string
}

func NewHttpClient(method string, url string, data string, headers string) HttpClient {
	return HttpClient{method, url, data, headers}
}

func (client HttpClient) Request() {

}
