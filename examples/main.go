package main

import (
	"context"
	"fmt"
	"github.com/ehsandavari/go-context-plus"
	"github.com/google/uuid"
)

func main() {

	ctx := contextplus.NewContext(context.Background()).
		SetValue("test", "test_value").
		SetRequestId(uuid.New().String()).
		SetTraceId(uuid.New().String())

	ctx.User.SetId(uuid.New()).SetPhoneNumber("989215580690")
	fmt.Println(
		ctx.Value("test"),
		ctx.User.Id(),
		ctx.User.PhoneNumber(),
		ctx.RequestId(),
		ctx.TraceId(),
	)

	fmt.Println(
		ctx.SetValue("test", "test_value").Value("test"),
		ctx.User.SetId(uuid.New()).Id(),
		ctx.User.SetPhoneNumber("09215580690").PhoneNumber(),
		ctx.SetRequestId(uuid.New().String()).RequestId(),
		ctx.SetTraceId(uuid.New().String()).TraceId(),
	)

	golangContext(ctx)
	myContext(ctx)
	myContext(contextplus.NewContext(context.Background()))
}

func golangContext(ctx context.Context) {
	fmt.Println(ctx.Value("test"))
}

func myContext(ctx *contextplus.Context) {
	fmt.Println(ctx.Value("test"))
}
