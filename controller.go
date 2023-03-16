package main

import (
	"context"
	"fmt"
	"my_webserver/framework"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	durationCtx, cancle := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancle()
	finish := make(chan struct{}, 1)
	panicchan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicchan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")
		finish <- struct{}{}
	}()

	select {
	case <-panicchan:
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}
