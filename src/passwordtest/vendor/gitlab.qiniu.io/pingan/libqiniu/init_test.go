package libqiniu_test

import (
	. "gitlab.qiniu.io/pingan/libqiniu"
	"gitlab.qiniu.io/pingan/libqiniu/fop"
)

var aksk = AccessKeySecretKey{
	AccessKey: "yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_",
	SecretKey: "dlt_zkvBaheHD13UQD-LJNlYBNFq-kVer4No2fMz",
}

var zone = &Zone{
	RsHost:     "http://rs.qiniu.com",
	RsfHost:    "http://rsf.qiniu.com",
	PfopHost:   "http://pfop.qiniu.com",
	StatusHost: "http://status.qiniu.com",
	UpHosts:    []string{"http://up.qiniu.com"},
}

type imageslim struct {
	_ fop.NoValue `fop:"imageslim"`
}
