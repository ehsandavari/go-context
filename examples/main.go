package main

import (
	"context"
	"fmt"
	"github.com/ehsandavari/go-context-plus"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	ginCtx := &gin.Context{}

	ginContext := contextplus.FromGinContext(ginCtx)
	ginContext.SetRequestId("asdad")

	ginContext.ToGinContext(ginCtx)

	ginContext1 := contextplus.FromGinContext(ginCtx)
	ginContext1.RequestId()

	ctx := contextplus.Background()
	ctx.SetValue("test", "test_value")
	ctx.SetRequestId(uuid.New().String())
	ctx.SetTraceId(uuid.New().String())
	ctx.User.SetId(uuid.New())
	ctx.User.SetPhoneNumber("989215580690")

	fmt.Println(
		ctx.Value("test"),
		ctx.User.Id(),
		ctx.User.PhoneNumber(),
		ctx.RequestId(),
		ctx.TraceId(),
	)

	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")
	ctx.SetValue("test", "test_value")

	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	golangContext(ctx)
	myContext(ctx)
	myContext(contextplus.NewContext(context.Background()))
}

func golangContext(ctx1 context.Context) {
	ctx1 = context.WithValue(ctx1, "1", "asd")
	ctx1 = context.WithValue(ctx1, "2", "asd")
	ctx1 = context.WithValue(ctx1, "3", "asd")
	ctx1 = context.WithValue(ctx1, "5", "asd")
	ctx1 = context.WithValue(ctx1, "4", "asd")
	ctx1 = context.WithValue(ctx1, "213", "asd")
	ctx1 = context.WithValue(ctx1, "a123sd", "asd")
	ctx1 = context.WithValue(ctx1, "qwe", "asd")
	ctx1 = context.WithValue(ctx1, "we", "asd")
	ctx1 = context.WithValue(ctx1, "qw", "asd")
	ctx1 = context.WithValue(ctx1, "g", "asd")
	ctx1 = context.WithValue(ctx1, "da", "asd")
	ctx1 = context.WithValue(ctx1, "xc", "asd")
	ctx1 = context.WithValue(ctx1, "e", "asd")
	ctx1 = context.WithValue(ctx1, "23", "asd")
	ctx1 = context.WithValue(ctx1, "43", "asd")
	ctx1 = context.WithValue(ctx1, "12", "asd")
	ctx1 = context.WithValue(ctx1, "as4d", "asd")
	fmt.Println(ctx1.Value("test"))
}

func myContext(ctx *contextplus.Context) {
	fmt.Println(ctx.Value("test"))
}
