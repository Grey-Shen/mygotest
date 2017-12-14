package op

import (
	"net/url"
)

type encodable interface {
	Encode() string
}

type BatchOps struct {
	ops []encodable
}

func NewBatchOps(ops ...encodable) *BatchOps {
	return &BatchOps{ops: ops}
}

func (batch *BatchOps) AddOp(op encodable) {
	batch.ops = append(batch.ops, op)
}

func (batch *BatchOps) FormEncode() string {
	values := make(url.Values)
	for _, op := range batch.ops {
		values.Add("op", op.Encode())
	}
	return values.Encode()
}
