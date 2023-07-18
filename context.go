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

func NewContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
	}
}

func (r *Context) SetValue(key, val any) *Context {
	r.Context = context.WithValue(r.Context, key, val)
	return r
}

func (r *Context) RequestId() string {
	return r.requestId
}

func (r *Context) SetRequestId(requestId string) *Context {
	r.requestId = requestId
	return r
}

func (r *Context) TraceId() string {
	return r.traceId
}

func (r *Context) SetTraceId(traceId string) *Context {
	r.traceId = traceId
	return r
}
