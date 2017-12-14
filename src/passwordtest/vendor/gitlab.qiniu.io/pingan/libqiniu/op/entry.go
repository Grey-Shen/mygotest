package op

import (
	"encoding/base64"
	"strings"
)

type Entry struct {
	Bucket string
	Key    string
}

func EntryFromString(str string) Entry {
	var entry Entry
	parts := strings.SplitN(str, ":", 2)
	switch len(parts) {
	case 2:
		entry.Key = parts[1]
	case 1:
		entry.Bucket = parts[0]
	}
	return entry
}

func NewEntry(bucket, key string) Entry {
	return Entry{Bucket: bucket, Key: key}
}

func NewEntryWithoutKey(bucket string) Entry {
	return Entry{Bucket: bucket}
}

func (entry Entry) String() string {
	str := entry.Bucket
	if entry.Key != "" {
		str += ":" + entry.Key
	}
	return str
}

func (entry Entry) Encode() string {
	return base64.URLEncoding.EncodeToString([]byte(entry.String()))
}
