package libqiniu

import (
	"net/http"
	"net/url"

	"gitlab.qiniu.io/pingan/libqiniu/op"
	ctx "golang.org/x/net/context"
)

type File struct {
	op.FileInfo
	Entry
	Key string `json:"key"`
}

type listFilesRet struct {
	Marker string `json:"marker"`
	Files  []File `json:"items"`
}

type FilesIterator struct {
	bucket    *Bucket
	prefix    string
	marker    string
	restFiles []File
	context   ctx.Context
	hasMore   bool
	err       error
}

func (bucket *Bucket) List(context ctx.Context, prefix string) *FilesIterator {
	return &FilesIterator{bucket: bucket, prefix: prefix, hasMore: true, context: context}
}

func (iterator *FilesIterator) Next(file *File) bool {
	for {
		if len(iterator.restFiles) > 0 {
			*file = iterator.restFiles[0]
			file.Entry = Entry{bucket: iterator.bucket, key: iterator.restFiles[0].Key}
			iterator.restFiles = iterator.restFiles[1:]
			return true
		} else if iterator.hasMore && iterator.askMoreBatch() {
			continue
		}
		return false
	}
}

func (iterator *FilesIterator) Err() error {
	return iterator.err
}

func (iterator *FilesIterator) askMoreBatch() bool {
	query := make(url.Values)
	query.Add("bucket", iterator.bucket.name)
	query.Add("limit", "1000")
	if iterator.prefix != "" {
		query.Add("prefix", iterator.prefix)
	}
	if iterator.marker != "" {
		query.Add("marker", iterator.marker)
	}

	var ret listFilesRet
	pathWithQuery := "/list?" + query.Encode()
	url := iterator.bucket.RsfHost + pathWithQuery
	if auth, err := iterator.bucket.generateAuthorization(pathWithQuery, nil); err != nil {
		iterator.err = err
		return false
	} else if _, err = iterator.bucket.postman.Send(iterator.context, "POST", url, auth, http.NoBody, &ret); err != nil {
		iterator.err = err
		return false
	}
	if len(ret.Files) == 0 {
		return false
	} else {
		if ret.Marker != "" {
			iterator.marker = ret.Marker
		} else {
			iterator.hasMore = false
		}
		iterator.restFiles = ret.Files
		return true
	}
}
