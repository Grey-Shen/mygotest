package libqiniu

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type signedRoundTripper struct {
	transport http.RoundTripper
	aksk      AccessKeySecretKey
}

func (roundTripper signedRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if value := request.Context().Value("donot_sign_request"); value == nil {
		if err := roundTripper.aksk.SignRequest(request); err != nil {
			return nil, err
		}
	}

	transport := roundTripper.transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	return transport.RoundTrip(request)
}

func (aksk AccessKeySecretKey) SignRequest(request *http.Request) error {
	if authorization, err := aksk.requestAuthorization(request); err != nil {
		return err
	} else {
		request.Header.Set("Authorization", authorization)
		return nil
	}
}

func (aksk AccessKeySecretKey) ValidateCallback(request *http.Request) (bool, error) {
	actualAuthorization := request.Header.Get("Authorization")
	if !strings.HasPrefix(actualAuthorization, "QBox ") {
		return false, nil
	}
	if expectedAuthorization, err := aksk.requestAuthorization(request); err != nil {
		return false, err
	} else {
		return actualAuthorization == expectedAuthorization, nil
	}
}

const wwwFormUrlencoded = "application/x-www-form-urlencoded"

func (aksk AccessKeySecretKey) requestAuthorization(request *http.Request) (string, error) {
	var (
		body []byte
		err  error
	)

	pathWithQuery := request.URL.Path
	if request.URL.RawQuery != "" {
		pathWithQuery += "?" + request.URL.RawQuery
	}

	if wwwFormUrlencoded == request.Header.Get("Content-Type") {
		if body, err = ioutil.ReadAll(request.Body); err != nil {
			return "", err
		} else if err = request.Body.Close(); err != nil {
			return "", err
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	return aksk.generateAuthorization(pathWithQuery, body)
}

func (aksk AccessKeySecretKey) generateAuthorization(pathWithQuery string, body []byte) (string, error) {
	var (
		buf = bytes.NewBuffer(make([]byte, 0, 1024))
		err error
	)
	if _, err = buf.WriteString(pathWithQuery); err != nil {
		return "", err
	}
	if _, err = buf.WriteString("\n"); err != nil {
		return "", err
	}
	if body != nil && len(body) > 0 {
		if _, err = buf.Write(body); err != nil {
			return "", err
		}
	}
	sign := aksk.sign(buf.Bytes())
	authorization := "QBox " + aksk.AccessKey + ":" + sign
	return authorization, nil
}
