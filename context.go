package context

import (
	"context"
)

type ApplicationContext struct {
	context.Context
	User      User
	requestId string
	traceId   string
}

func NewApplicationContext(ctx context.Context) *ApplicationContext {
	return &ApplicationContext{
		Context: ctx,
	}
}

func (r *ApplicationContext) SetValue(key, val any) *ApplicationContext {
	r.Context = context.WithValue(r.Context, key, val)
	return r
}

func (r *ApplicationContext) RequestId() string {
	return r.requestId
}

func (r *ApplicationContext) SetRequestId(requestId string) *ApplicationContext {
	r.requestId = requestId
	return r
}

func (r *ApplicationContext) TraceId() string {
	return r.traceId
}

func (r *ApplicationContext) SetTraceId(traceId string) *ApplicationContext {
	r.traceId = traceId
	return r
}
