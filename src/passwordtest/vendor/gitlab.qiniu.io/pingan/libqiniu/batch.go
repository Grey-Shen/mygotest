package libqiniu

import (
	"strings"

	"gitlab.qiniu.io/pingan/libqiniu/op"
	ctx "golang.org/x/net/context"
)

type Batch struct {
	bucket *Bucket
	ops    *op.BatchOps
}

func (self *Batch) Stat(key string) *Batch {
	self.ops.AddOp(op.Stat{Entry: op.Entry{Bucket: self.bucket.name, Key: key}})
	return self
}

func (self *Batch) ChangeMime(key, mime string) *Batch {
	self.ops.AddOp(op.Chgm{Entry: op.NewEntry(self.bucket.name, key), Mime: mime})
	return self
}

func (self *Batch) ChangeType(key string, t op.Type) *Batch {
	self.ops.AddOp(op.Chtype{Entry: op.NewEntry(self.bucket.name, key), Type: t})
	return self
}

func (self *Batch) MoveTo(key, destBucket string, force bool) *Batch {
	return self.Move(key, destBucket, key, force)
}

func (self *Batch) Move(srcKey, destBucket, destKey string, force bool) *Batch {
	self.ops.AddOp(op.Move{Src: op.NewEntry(self.bucket.name, srcKey), Dest: op.NewEntry(destBucket, destKey), Force: force})
	return self
}

func (self *Batch) CopyTo(key, destBucket string, force bool) *Batch {
	return self.Copy(key, destBucket, key, force)
}

func (self *Batch) Copy(srcKey, destBucket, destKey string, force bool) *Batch {
	self.ops.AddOp(op.Copy{Src: op.NewEntry(self.bucket.name, srcKey), Dest: op.NewEntry(destBucket, destKey), Force: force})
	return self
}

func (self *Batch) Delete(key string) *Batch {
	self.ops.AddOp(op.Delete{Entry: op.NewEntry(self.bucket.name, key)})
	return self
}

func (self *Batch) DeleteAfterDays(key string, afterDays uint64) *Batch {
	self.ops.AddOp(op.DeleteAfterDays{Entry: op.NewEntry(self.bucket.name, key), AfterDays: afterDays})
	return self
}

func (self *Batch) Do(context ctx.Context) (op.BatchOpsRet, error) {
	var ret op.BatchOpsRet
	body := self.ops.FormEncode()
	if auth, err := self.bucket.generateAuthorization("/batch", []byte(body)); err != nil {
		return nil, err
	} else {
		_, err := self.bucket.postman.Send(context, "POST", self.bucket.RsHost+"/batch", auth, strings.NewReader(body), &ret)
		return ret, err
	}
}
