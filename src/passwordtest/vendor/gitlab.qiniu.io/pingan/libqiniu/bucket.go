package libqiniu

import (
	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	"gitlab.qiniu.io/pingan/libqiniu/op"
)

type Bucket struct {
	postman    form_postman.Postman
	name       string
	domain     string
	pfopClient PfopClient
	AccessKeySecretKey
	*Zone
}

func (bucket *Bucket) Name() string {
	return bucket.name
}

func (bucket *Bucket) Domain() string {
	return bucket.domain
}

func (bucket *Bucket) SetDomain(domain string) {
	bucket.domain = domain
}

func (bucket *Bucket) Entry(key string) Entry {
	return Entry{bucket: bucket, key: key}
}

func (bucket *Bucket) Batch() *Batch {
	return &Batch{bucket: bucket, ops: op.NewBatchOps()}
}

func (bucket *Bucket) CreatePutPolicyForBucket() PutPolicyWithoutDeadline {
	return NewPutPolicyForBucket(bucket.name)
}

func (bucket *Bucket) CreatePutPolicyForPrefix(prefix string) PutPolicyWithoutDeadline {
	return NewPutPolicyForPrefix(bucket.name, prefix)
}
