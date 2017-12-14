package libqiniu

import "gitlab.qiniu.io/pingan/libqiniu/op"

type PfopCode int

const (
	PFOP_SUCCESSFUL      = PfopCode(0)
	PFOP_WAITING         = PfopCode(1)
	PFOP_EXECUTING       = PfopCode(2)
	PFOP_FAILED          = PfopCode(3)
	PFOP_CALLBACK_FAILED = PfopCode(4)
)

type PfopStatus struct {
	Id             string     `json:"id"`
	Pipeline       string     `json:"pipeline"`
	Code           PfopCode   `json:"code"`
	Desc           string     `json:"desc"`
	Reqid          string     `json:"reqid"`
	InputBucket    string     `json:"inputBucket"`
	InputKey       string     `json:"inputKey"`
	Items          []PfopItem `json:"items"`
	originalString string
}

type PfopItem struct {
	Cmd       string   `json:"cmd"`
	Code      PfopCode `json:"code"`
	Desc      string   `json:"desc"`
	Error     string   `json:"error,omitempty"`
	Hash      string   `json:"hash,omitempty"`
	Key       string   `json:"key,omitempty"`
	Keys      []string `json:"keys,omitempty"`
	ReturnOld op.Bool  `json:"returnOld"`
}

func (status *PfopStatus) String() string {
	return status.originalString
}
