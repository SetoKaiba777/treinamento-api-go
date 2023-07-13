package httpclient

import (
	"io"
	"net/http"
	"payments-go/core/gateway"
)

type HttpClient struct {
	client http.Client
}

func NewHttpClient() HttpClient{
	return HttpClient{
		client: *http.DefaultClient,
	}
}

func (c HttpClient) Do(method string, target string, httpHeaders http.Header, httpBody io.Reader) gateway.Response{
		req, err := c.buildRequest(method,target,httpHeaders,httpBody)
		if err != nil{
			return gateway.Response{
				Resp: http.Response{},
				Err: err,
			}
		}
		
		resp, err := c.client.Do(req)
		if err != nil{
			return gateway.Response{
				Resp: http.Response{},
				Err: err,
			}
		}

		return gateway.Response{
			Resp: *resp,
			Err: err,
		}
}

func (c *HttpClient) buildRequest(method,url string, headers http.Header, body io.Reader) (*http.Request, error){
	req, err := http.NewRequest(method,url,body)
	if err!= nil{
		return nil, err
	}

	for header,values := range headers{
		for _, value := range values{
			req.Header.Add(header,value)
		}
	}
	return req, nil
}