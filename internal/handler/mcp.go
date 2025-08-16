package handler

import (
	"aliya-ram/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type McpHandler struct {
	mcpService *service.MCP
}

func NewMcpHandler() (*McpHandler, error) {
	m, err := service.NewMCP()
	if err != nil {
		log.Printf("failed to create MCP service: %v", err)
		return nil, err
	}
	return &McpHandler{mcpService: m}, nil
}

func (h *McpHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.mcpService.Handler.HandleMCP().ServeHTTP(c.Writer, c.Request)
	}
}
