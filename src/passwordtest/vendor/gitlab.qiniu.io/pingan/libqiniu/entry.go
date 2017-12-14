package libqiniu

import (
	"net/http"
	"net/url"

	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	"gitlab.qiniu.io/pingan/libqiniu/op"
	ctx "golang.org/x/net/context"
)

type Entry struct {
	bucket *Bucket
	key    string
}

func (entry Entry) BucketName() string {
	return entry.bucket.Name()
}

func (entry Entry) BucketDomain() string {
	return entry.bucket.Domain()
}

func (entry Entry) CreatePutPolicy() PutPolicyWithoutDeadline {
	return NewPutPolicy(entry.bucket.name, entry.key)
}

func (entry Entry) toEntry() op.Entry {
	return op.NewEntry(entry.bucket.name, entry.key)
}

func (entry Entry) URL() *urlWithoutDeadline {
	u, err := url.Parse("http://" + entry.BucketDomain() + "/" + entry.key)
	if err != nil {
		panic(err)
	}
	return entry.bucket.NewURL(u)
}

func (entry Entry) Pfop(context ctx.Context, params PfopParams) (PersistentID, error) {
	return entry.bucket.pfopClient.Pfop(context, op.Entry{Bucket: entry.BucketName(), Key: entry.key}, params)
}

func (entry Entry) Stat(context ctx.Context) (op.FileInfo, error) {
	var fileInfo op.FileInfo
	err := send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Stat{Entry: entry.toEntry()}, &fileInfo)
	return fileInfo, err
}

func (entry Entry) ChangeMime(context ctx.Context, mime string) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Chgm{Entry: entry.toEntry(), Mime: mime}, nil)
}

func (entry Entry) ChangeType(context ctx.Context, t Type) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Chtype{Entry: entry.toEntry(), Type: t}, nil)
}

func (entry Entry) MoveTo(context ctx.Context, destBucket string, force bool) error {
	return entry.Move(context, destBucket, entry.key, force)
}

func (entry Entry) Move(context ctx.Context, destBucket, destKey string, force bool) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Move{Src: entry.toEntry(), Dest: op.NewEntry(destBucket, destKey), Force: force}, nil)
}

func (entry Entry) CopyTo(context ctx.Context, destBucket string, force bool) error {
	return entry.Copy(context, destBucket, entry.key, force)
}

func (entry Entry) Copy(context ctx.Context, destBucket, destKey string, force bool) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Copy{Src: entry.toEntry(), Dest: op.NewEntry(destBucket, destKey), Force: force}, nil)
}

func (entry Entry) Delete(context ctx.Context) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.Delete{Entry: entry.toEntry()}, nil)
}

func (entry Entry) DeleteAfterDays(context ctx.Context, afterDays uint64) error {
	return send(entry.bucket.postman, entry.bucket.AccessKeySecretKey, context, entry.bucket.RsHost, op.DeleteAfterDays{Entry: entry.toEntry(), AfterDays: afterDays}, nil)
}

type encodable interface {
	Encode() string
}

func send(postman form_postman.Postman, aksk AccessKeySecretKey, context ctx.Context, host string, op encodable, responseBody interface{}) error {
	encodedOp := op.Encode()
	if auth, err := aksk.generateAuthorization(encodedOp, nil); err != nil {
		return err
	} else {
		_, err := postman.Send(context, "POST", host+encodedOp, auth, http.NoBody, responseBody)
		return err
	}
}
