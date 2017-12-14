package op

type Bool = uint8

const (
	False = Bool(0)
	True  = Bool(1)
)

type PutPolicy struct {
	Scope               string `json:"scope"`
	Deadline            int64  `json:"deadline"` // 截止时间（以秒为单位）
	IsPrefixalScope     Bool   `json:"isPrefixalScope,omitempty"`
	InsertOnly          Bool   `json:"insertOnly,omitempty"` // 若非0, 即使Scope为 Bucket:Key 的形式也是insert only
	DetectMIME          Bool   `json:"detectMime,omitempty"` // 若非0, 则服务端根据内容自动确定 MimeType
	FsizeLimit          int64  `json:"fsizeLimit,omitempty"`
	FsizeMin            int64  `json:"fsizeMin,omitempty"`
	MIMELimit           string `json:"mimeLimit,omitempty"`
	SaveKey             string `json:"saveKey,omitempty"`
	CallbackURL         string `json:"callbackUrl,omitempty"`
	CallbackHost        string `json:"callbackHost,omitempty"`
	CallbackBody        string `json:"callbackBody,omitempty"`
	CallbackBodyType    string `json:"callbackBodyType,omitempty"`
	ReturnURL           string `json:"returnUrl,omitempty"`
	ReturnBody          string `json:"returnBody,omitempty"`
	PersistentOps       string `json:"persistentOps,omitempty"`
	PersistentNotifyURL string `json:"persistentNotifyUrl,omitempty"`
	PersistentPipeline  string `json:"persistentPipeline,omitempty"`
	EndUser             string `json:"endUser,omitempty"`
	DeleteAfterDays     uint64 `json:"deleteAfterDays,omitempty"`
	FileType            Type   `json:"fileType,omitempty"`
}
