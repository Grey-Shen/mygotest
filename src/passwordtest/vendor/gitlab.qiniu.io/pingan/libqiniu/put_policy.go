package libqiniu

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"gitlab.qiniu.io/pingan/libqiniu/op"
)

type Bool = op.Bool
type Type = op.Type

const (
	False          = op.False
	True           = op.True
	TypeNormal     = op.TypeNormal
	TypeInfrequent = op.TypeInfrequent
)

type PutPolicy struct {
	policy *op.PutPolicy
}

type PutPolicyWithoutDeadline struct {
	entry           op.Entry
	isPrefixalScope Bool
}

func NewPutPolicyForBucket(bucket string) PutPolicyWithoutDeadline {
	return PutPolicyWithoutDeadline{
		entry: op.NewEntryWithoutKey(bucket),
	}
}

func NewPutPolicy(bucket, key string) PutPolicyWithoutDeadline {
	return PutPolicyWithoutDeadline{
		entry: op.NewEntry(bucket, key),
	}
}

func NewPutPolicyForPrefix(bucket, prefix string) PutPolicyWithoutDeadline {
	return PutPolicyWithoutDeadline{
		entry:           op.NewEntry(bucket, prefix),
		isPrefixalScope: True,
	}
}

func (self PutPolicyWithoutDeadline) SetLifetime(lifetime time.Duration) PutPolicy {
	putPolicy := PutPolicy{
		policy: &op.PutPolicy{
			Scope:           self.entry.String(),
			Deadline:        time.Now().Add(lifetime).Unix(),
			IsPrefixalScope: self.isPrefixalScope,
			SaveKey:         self.entry.Key,
		}}
	return putPolicy
}

func (self PutPolicyWithoutDeadline) SetDeadline(deadline time.Time) PutPolicy {
	putPolicy := PutPolicy{
		policy: &op.PutPolicy{
			Scope:           self.entry.String(),
			Deadline:        deadline.Unix(),
			IsPrefixalScope: self.isPrefixalScope,
			SaveKey:         self.entry.Key,
		}}
	return putPolicy
}

func (self PutPolicy) InsertOnly() PutPolicy {
	self.policy.InsertOnly = True
	return self
}

func (self PutPolicy) DetectMIME() PutPolicy {
	self.policy.DetectMIME = True
	return self
}

func (self PutPolicy) SetSaveKey(key string) PutPolicy {
	self.policy.SaveKey = key
	return self
}

func (self PutPolicy) LimitMIME(mime string) PutPolicy {
	self.policy.MIMELimit = mime
	return self
}

func (self PutPolicy) SetMinFileSize(size int64) PutPolicy {
	self.policy.FsizeMin = size
	return self
}

func (self PutPolicy) SetMaxFileSize(size int64) PutPolicy {
	self.policy.FsizeLimit = size
	return self
}

func (self PutPolicy) Infrequent() PutPolicy {
	self.policy.FileType = TypeInfrequent
	return self
}

func (self PutPolicy) DeleteAfterDays(days uint64) PutPolicy {
	self.policy.DeleteAfterDays = days
	return self
}

func (self PutPolicy) SetReturnURL(url string) PutPolicy {
	self.policy.ReturnURL = url
	return self
}

func (self PutPolicy) SetReturnBody(body string) PutPolicy {
	self.policy.ReturnBody = body
	return self
}

func (self PutPolicy) SetCallback(params PutPolicyCallbackParams) PutPolicy {
	self.policy.CallbackURL = strings.Join(params.urls, ";")
	self.policy.CallbackHost = params.host
	self.policy.CallbackBodyType = params.contentType
	self.policy.CallbackBody = params.body
	return self
}

func (self PutPolicy) SetPfop(params PutPolicyPfopParams) PutPolicy {
	self.policy.PersistentOps = params.cmd
	self.policy.PersistentNotifyURL = params.notifyURL
	self.policy.PersistentPipeline = params.pipeline
	return self
}

func (self PutPolicy) Entry() op.Entry {
	return op.EntryFromString(self.policy.Scope)
}

func (self PutPolicy) GetLifetime() time.Duration {
	return self.GetDeadline().Sub(time.Now())
}

func (self PutPolicy) GetDeadline() time.Time {
	return time.Unix(self.policy.Deadline, 0)
}

func (self PutPolicy) GetPutPolicy() (PutPolicy, error) {
	return self, nil
}

func (self PutPolicy) Base64Encode() (string, error) {
	if putPolicyBytes, err := json.Marshal(self.policy); err != nil {
		return "", err
	} else {
		base64PutPolicy := base64.URLEncoding.EncodeToString(putPolicyBytes)
		return base64PutPolicy, nil
	}
}
