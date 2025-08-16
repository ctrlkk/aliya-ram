package service

import (
	"aliya-ram/app/aliyun"
	"aliya-ram/app/ram"
	"context"
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	_ "github.com/mattn/go-sqlite3"
)

type MCP struct {
	Server  *server.Server
	Handler *transport.StreamableHTTPHandler
}

type RecallRequest struct {
	Keyword string `json:"keyword" description:"关键词" required:"true"`
}

type AddMemoriesRequest struct {
	UserID string `json:"userId" description:"用户id" required:"true"`
	Text   string `json:"text" description:"需要存储的内容" required:"true"`
}

type SearchMemoryRequest struct {
	UserID  string `json:"userId" description:"用户id" required:"true"`
	Keyword string `json:"keyword" description:"关键词" required:"true"`
}

type MemoriesRequest struct {
	UserID string `json:"userId" description:"用户id" required:"true"`
}

type CreateIndexsRequest struct {
	Authorize string `json:"authorize" description:"授权码" required:"true"`
	UserID    string `json:"userId" description:"用户id" required:"false"`
}

type UpdateUserIDRequest struct {
	Authorize string `json:"authorize" description:"授权码" required:"true"`
	OldUserID string `json:"oldUserId" description:"原用户id" required:"true"`
	NewUserID string `json:"newUserId" description:"新用户id" required:"true"`
}

var r *ram.RAM
var bailian *aliyun.Bailian

func NewMCP() (*MCP, error) {
	transport, handler, err := transport.NewStreamableHTTPServerTransportAndHandler()
	if err != nil {
		return nil, err
	}
	mcpServer, err := server.NewServer(transport)
	if err != nil {
		return nil, err
	}
	r, err = ram.NewRAM()
	if err != nil {
		return nil, err
	}
	bailian, err = aliyun.NewBalilian()
	if err != nil {
		return nil, err
	}

	var tool *protocol.Tool
	// tool, _ = protocol.NewTool("memory", "根据关键词检索与Aliya相关的记忆。每当用户提出超出当前设置范围的问题时，就需要调用此方法。此外，如果需要进行补充设定，也应调用此方法。查询结果作为数据源的补充，之后结合世界观和Aliya的人设自由发挥。", RecallRequest{})
	// mcpServer.RegisterTool(tool, documentSearch)

	tool, _ = protocol.NewTool("add_memories", "添加新的记忆。每当用户告知自己、他们的偏好或任何与未来对话相关的信息时，就需要调用此方法。此外，如果用户询问您需要记住的内容，也应调用此方法。", AddMemoriesRequest{})
	mcpServer.RegisterTool(tool, addMemories)

	tool, _ = protocol.NewTool("search_memory", "检索存储的记忆。这种方法在用户提出任何请求时都会被调用。", SearchMemoryRequest{})
	mcpServer.RegisterTool(tool, searchMemory)

	tool, _ = protocol.NewTool("list_memories", "列出用户记忆中的所有记忆内容，每当出现新用户或新对话时都应该调用此方法。", MemoriesRequest{})
	mcpServer.RegisterTool(tool, listMemories)

	tool, _ = protocol.NewTool("delete_all_memories", "删除用户记忆中的所有内容", MemoriesRequest{})
	mcpServer.RegisterTool(tool, deleteAllMemories)

	tool, _ = protocol.NewTool("create_indexs", "重新创建索引，这种方法禁止使用", CreateIndexsRequest{})
	mcpServer.RegisterTool(tool, createIndexs)

	tool, _ = protocol.NewTool("update_user_id", "修改用户id，这种方法禁止使用", UpdateUserIDRequest{})
	mcpServer.RegisterTool(tool, updateUserID)

	return &MCP{
		Server:  mcpServer,
		Handler: handler,
	}, nil
}

func addMemories(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq AddMemoriesRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	err := r.AddMemory(recallReq.UserID, recallReq.Text)
	if err != nil {
		log.Printf("add memory error:%s\n", err.Error())
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "error",
				},
			},
		}, nil
	}

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: "succeed",
			},
		},
	}, nil
}

func searchMemory(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq SearchMemoryRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	result := []protocol.Content{}

	memories, err := r.SearchMemory(recallReq.UserID, recallReq.Keyword)
	if err != nil {
		log.Printf("search memory error:%s\n", err.Error())
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "error",
				},
			},
		}, nil
	}
	for _, memory := range memories {
		result = append(result, &protocol.TextContent{
			Type: "text",
			Text: memory,
		})
	}

	nodes, err := bailian.Query(recallReq.Keyword)
	if err != nil {
		log.Printf("查询 Bailian 失败: %v", err)
	} else {
		count := 0
		for _, node := range nodes {
			if *node.GetScore() < 0.5 || count >= 3 {
				continue
			}
			result = append(result, &protocol.TextContent{
				Type: "text",
				Text: *node.GetText(),
			})
			count++
		}
	}

	if len(result) <= 1 {
		result = append(result, &protocol.TextContent{
			Type: "text",
			Text: "没有找到相关内容",
		})
	}

	return &protocol.CallToolResult{
		Content: result,
	}, nil
}

func listMemories(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq MemoriesRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	result := []protocol.Content{}
	memories, err := r.ListMemories(recallReq.UserID)
	if err != nil {
		log.Printf("list memories error:%s\n", err.Error())
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "error",
				},
			},
		}, nil
	}
	for _, memory := range memories {
		result = append(result, &protocol.TextContent{
			Type: "text",
			Text: memory,
		})
	}

	if len(result) <= 1 {
		result = append(result, &protocol.TextContent{
			Type: "text",
			Text: "没有找到相关内容",
		})
	}

	return &protocol.CallToolResult{
		Content: result,
	}, nil
}

func deleteAllMemories(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq MemoriesRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	err := r.DeleteAllMemories(recallReq.UserID)
	if err != nil {
		log.Println("delete all memories error:", err.Error())
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "删除失败",
				},
			},
		}, nil
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: "删除成功",
			},
		},
	}, nil
}

func createIndexs(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq CreateIndexsRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	if recallReq.Authorize != "Ev7x7PMdG8foi9" {
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "无授权",
				},
			},
		}, nil
	}
	err := r.CreateIndexs(recallReq.UserID)
	if err != nil {
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "error",
				},
			},
		}, nil
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: "成功",
			},
		},
	}, nil
}

func updateUserID(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var recallReq UpdateUserIDRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &recallReq); err != nil {
		return nil, err
	}
	if recallReq.Authorize != "Ev7x7PMdG8foi9" {
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "无授权",
				},
			},
		}, nil
	}
	_, err := r.UpdateUserID(recallReq.OldUserID, recallReq.NewUserID)
	if err != nil {
		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: "error",
				},
			},
		}, nil
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: "成功",
			},
		},
	}, nil
}
