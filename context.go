package context

import (
	"context"
	"github.com/google/uuid"
)

type ApplicationContext struct {
	context.Context
}

func NewApplicationContext(ctx context.Context) *ApplicationContext {
	return &ApplicationContext{
		ctx,
	}
}

func (r *ApplicationContext) SetValue(key, val any) {
	r.Context = context.WithValue(r.Context, key, val)
}

func (r *ApplicationContext) GetValue(key any) any {
	return r.Value(key)
}

func (r *ApplicationContext) SetUserId(userId uuid.UUID) {
	r.Context = context.WithValue(r.Context, "userId", userId)
}

func (r *ApplicationContext) GetUserId() uuid.UUID {
	return uuid.MustParse(r.Value("userId").(string))
}

func (r *ApplicationContext) SetRequestId(requestId string) {
	r.Context = context.WithValue(r.Context, "requestId", requestId)
}

func (r *ApplicationContext) GetRequestId() string {
	return r.Value("requestId").(string)
}

func (r *ApplicationContext) SetTraceId(traceId string) {
	r.Context = context.WithValue(r.Context, "traceId", traceId)
}

func (r *ApplicationContext) GetTraceId() string {
	return r.Value("traceId").(string)
}

func (r *ApplicationContext) SetPhoneNumber(phoneNumber string) {
	r.Context = context.WithValue(r.Context, "phoneNumber", phoneNumber)
}

func (r *ApplicationContext) GetPhoneNumber() string {
	return r.Value("phoneNumber").(string)
}
