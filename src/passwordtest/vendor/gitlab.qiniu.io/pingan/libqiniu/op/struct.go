package op

type FileInfo struct {
	Hash     string `json:"hash"`
	Fsize    int64  `json:"fsize"`
	PutTime  int64  `json:"putTime"`
	MimeType string `json:"mimeType"`
	Type     Type   `json:"type"`
}

type BatchOpRet struct {
	Code int `json:"code,omitempty"`
	Data struct {
		FileInfo
		Error string `json:"error"`
	} `json:"data,omitempty"`
}

type BatchOpsRet []*BatchOpRet

func (rets BatchOpsRet) AnyError() bool {
	for _, ret := range rets {
		if ret.Code >= 300 {
			return true
		}
	}
	return false
}
