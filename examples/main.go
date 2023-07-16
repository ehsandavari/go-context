package main

import (
	"context"
	"fmt"
	goContext "github.com/ehsandavari/go-context"
)

func main() {

	applicationContext := goContext.NewApplicationContext(context.Background())
	applicationContext.SetValue("asdasd", "w123123123")
	fmt.Println(applicationContext.GetValue("asdasd"))
	applicationContext.SetValue("userId", "a048d082-1fa3-4de6-a23e-1f5be8f369c6")
	fmt.Println(applicationContext.GetUserId())
}
