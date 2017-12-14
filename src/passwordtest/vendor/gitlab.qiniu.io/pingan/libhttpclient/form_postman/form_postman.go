package form_postman

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"gitlab.qiniu.io/pingan/libhttpclient/errors"
	ctx "golang.org/x/net/context"
)

type (
	RequestConstructorError = errors.RequestConstructorError
	JSONDecodeError         = errors.JSONDecodeError
	BodyCloseError          = errors.BodyCloseError
	IOReadError             = errors.IOReadError
	HTTPStatusError         = errors.HTTPStatusError
)

type Postman struct {
	client *http.Client
}

var EmptyParams = make(url.Values)

func NewPostman(client *http.Client) Postman {
	return Postman{
		client: client,
	}
}

func (postman Postman) PostForm(context ctx.Context, url, authorization string, values url.Values, responseBody interface{}) error {
	data := values.Encode()
	_, err := postman.Send(context, "POST", url, authorization, strings.NewReader(data), responseBody)
	return err
}

func (postman Postman) Send(context ctx.Context, method, url, authorization string, reqBodyReader io.Reader, responseBody interface{}) (*http.Response, error) {
	request, err := http.NewRequest(method, url, reqBodyReader)
	if err != nil {
		return nil, RequestConstructorError{Method: method, Url: url, Err: err}
	}
	if context != nil {
		request = request.WithContext(context)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	if authorization != "" {
		request.Header.Set("Authorization", authorization)
	}
	return postman.SendRequest(request, responseBody)
}

func (postman Postman) SendRequest(request *http.Request, responseBody interface{}) (*http.Response, error) {
	var (
		response *http.Response
		body     []byte
		err      error
	)

	if response, err = postman.client.Do(request); err != nil {
		return nil, RequestConstructorError{Method: request.Method, Url: request.URL.String(), Err: err}
	}

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, IOReadError{Err: err}
	} else if err := response.Body.Close(); err != nil {
		return nil, BodyCloseError{Err: err}
	} else {
		response.Body = ioutil.NopCloser(bytes.NewReader(body))

		if responseBody != nil {
			if err = json.NewDecoder(bytes.NewReader(body)).Decode(&responseBody); err != nil {
				return nil, JSONDecodeError{Origin: body, Err: err}
			}
		}
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return response, nil
	} else {
		return response, HTTPStatusError{StatusCode: response.StatusCode, Body: body}
	}
}
