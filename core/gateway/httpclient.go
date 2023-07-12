package gateway

import (
	"io"
	"net/http"
)

type (
	HttpClient interface {
		Do(method string, target string, httpHeaders http.Header, httpBody io.Reader) Response
	}

	Response struct {
		Resp http.Response
		Err  error
	}
)
