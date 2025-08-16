package router

import (
	"aliya-ram/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	r := gin.Default()

	mcpHandler, err := handler.NewMcpHandler()
	if err != nil {
		return nil, err
	}

	r.Any("/mcp", mcpHandler.Handle())

	return r, nil
}
