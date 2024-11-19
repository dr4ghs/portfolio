package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type CustomContext struct {
	context.Context
	StartTime time.Time
}

type (
	CustomHandler    func(ctx *CustomContext, w http.ResponseWriter, r *http.Request)
	CustomMiddleware func(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error
)

func Chain(
	w http.ResponseWriter,
	r *http.Request,
	middleware ...CustomMiddleware,
) {
	ctx := &CustomContext{
		Context:   context.Background(),
		StartTime: time.Now(),
	}

	for _, mw := range middleware {
		if err := mw(ctx, w, r); err != nil {
			return
		}
	}

	Log(ctx, w, r)
}

func Log(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error {
	elapsedTime := time.Since(ctx.StartTime)
	formattedTime := time.Now().Format("2024-11-19 22:45:00")
	fmt.Printf("[%s] [%s] [%s] [%s]\n", formattedTime, r.Method, r.URL.Path, elapsedTime)

	return nil
}
