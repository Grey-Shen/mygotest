package signature

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"strings"
)

type (
	InvalidAuthorizationFormat struct {
	}

	InvalidAuthorizationType struct {
		ActualType   string
		ExpectedType string
	}
)

func (err InvalidAuthorizationFormat) Error() string {
	return "Invalid authorization format or authorization is missing"
}

func (err InvalidAuthorizationType) Error() string {
	return fmt.Sprintf("Invalid authorization type, should be %s, but gives %s", err.ExpectedType, err.ActualType)
}

func ParseAuthorization(authorization, authType string) (id, signature string, err error) {
	if splits := strings.SplitN(authorization, " ", 2); len(splits) != 2 {
		err = InvalidAuthorizationFormat{}
		return
	} else if splits[0] != authType {
		err = InvalidAuthorizationType{ActualType: splits[0], ExpectedType: authType}
		return
	} else {
		authorization = splits[1]
	}

	if splits := strings.SplitN(authorization, ":", 2); len(splits) != 2 {
		err = InvalidAuthorizationFormat{}
		return
	} else {
		id, signature = splits[0], splits[1]
	}
	return
}

func CalcSignature(data, key string) []byte {
	return CalcSignatureForBinary([]byte(data), key)
}

func CalcSignatureForBinary(data []byte, key string) []byte {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(data)
	return mac.Sum(nil)
}
