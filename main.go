package main

import (
	"aliya-ram/app/mcp"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mcp, err := mcp.NewMCP()
	if err != nil {
		log.Fatal(err)
	}
	defer mcp.Server.Shutdown(context.Background())

	go func() {
		mcp.Server.Run()
	}()

	r := gin.Default()

	r.GET("/mcp", func(ctx *gin.Context) {
		mcp.Handler.HandleMCP().ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.POST("/mcp", func(ctx *gin.Context) {
		mcp.Handler.HandleMCP().ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err = r.Run(); err != nil {
		return
	}
}
