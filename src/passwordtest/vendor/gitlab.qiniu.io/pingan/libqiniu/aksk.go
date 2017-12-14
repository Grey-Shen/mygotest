package libqiniu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"
)

type AccessKeySecretKey struct {
	AccessKey string
	SecretKey string
}

func NewAccessKeySecretKey(accessKey, secretKey string) AccessKeySecretKey {
	return AccessKeySecretKey{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (aksk AccessKeySecretKey) NewURL(baseURL *url.URL) *urlWithoutDeadline {
	return &urlWithoutDeadline{baseURL: baseURL, aksk: aksk}
}

type PutPolicyGetter interface {
	GetPutPolicy() (PutPolicy, error)
}

func (aksk AccessKeySecretKey) Uptoken(getter PutPolicyGetter) (Uptoken, error) {
	if putPolicy, err := getter.GetPutPolicy(); err != nil {
		return Uptoken{}, err
	} else {
		return Uptoken{aksk: &aksk, putPolicy: putPolicy.policy}, nil
	}
}

func (aksk AccessKeySecretKey) SignedTransport(transport http.RoundTripper) http.RoundTripper {
	return signedRoundTripper{transport: transport, aksk: aksk}
}

func (aksk AccessKeySecretKey) sign(data []byte) string {
	h := hmac.New(sha1.New, []byte(aksk.SecretKey))
	h.Write(data)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
