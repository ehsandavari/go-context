package contextplus

import (
	"context"
)

type Context struct {
	context.Context
	User      User
	requestId string
	traceId   string
}

func NewContext(ctx context.Context) Context {
	return Context{
		Context: ctx,
	}
}

func (r *Context) SetValue(key, val any) {
	r.Context = context.WithValue(r.Context, key, val)
}

func (r *Context) RequestId() string {
	return r.requestId
}

func (r *Context) SetRequestId(requestId string) {
	r.requestId = requestId
}

func (r *Context) TraceId() string {
	return r.traceId
}

func (r *Context) SetTraceId(traceId string) {
	r.traceId = traceId
}
