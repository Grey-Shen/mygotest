package libqiniu

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"gitlab.qiniu.io/pingan/libqiniu/op"
)

type Uptoken struct {
	aksk      *AccessKeySecretKey
	putPolicy *op.PutPolicy
	token     string
}

func (uptoken Uptoken) GetPutPolicy() (PutPolicy, error) {
	if uptoken.putPolicy != nil {
		return PutPolicy{policy: uptoken.putPolicy}, nil
	} else {
		var err error
		uptoken.putPolicy, err = uptoken.getPutPolicy()
		return PutPolicy{policy: uptoken.putPolicy}, err
	}
}

func (uptoken Uptoken) getPutPolicy() (*op.PutPolicy, error) {
	putPolicy := new(op.PutPolicy)
	items := strings.SplitN(uptoken.token, ":", 3)
	if len(items) != 3 {
		return putPolicy, UptokenFormatError{Uptoken: uptoken}
	}
	if policyBytes, err := base64.URLEncoding.DecodeString(items[2]); err != nil {
		return putPolicy, Base64DecodeError{Uptoken: uptoken, Err: err}
	} else if err = json.Unmarshal(policyBytes, putPolicy); err != nil {
		return putPolicy, JsonDecodeError{Uptoken: uptoken, Err: err}
	} else {
		return putPolicy, nil
	}
}

func (uptoken Uptoken) String() string {
	if uptoken.token != "" {
		return uptoken.token
	} else if uptoken.aksk != nil && uptoken.putPolicy != nil {
		putPolicy := PutPolicy{policy: uptoken.putPolicy}
		if base64PutPolicy, err := putPolicy.Base64Encode(); err != nil {
			panic("Uptoken internal error")
		} else {
			sign := uptoken.aksk.sign([]byte(base64PutPolicy))
			uptoken.token = uptoken.aksk.AccessKey + ":" + sign + ":" + base64PutPolicy
			return uptoken.token
		}
	} else {
		panic("Uptoken internal error")
	}
}

type UptokenFormatError struct {
	Uptoken Uptoken
}

func (err UptokenFormatError) Error() string {
	return "Invalid uptoken, format error"
}

type Base64DecodeError struct {
	Uptoken Uptoken
	Err     error
}

func (err Base64DecodeError) Error() string {
	return "Invalid uptoken, base64 decode error"
}

type JsonDecodeError struct {
	Uptoken Uptoken
	Err     error
}

func (err JsonDecodeError) Error() string {
	return "Invalid uptoken, json decode error"
}
