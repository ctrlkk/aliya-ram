# Aliya RAM

Aliya 记忆体，通过最少的 token 实现最真实的 Aliya，它同时也是一个`AstrBot`友好的记忆持久化 MCP Servers

[本地部署](./docs/deploy.md)或加 QQ 群体验：761178785

> aliya 提示词已迁移到  [aliya-prompt](https://github.com/ctrlkk/aliya-prompt)

<img src="./docs/images/aliya_go.png" style="max-height:200px; width:auto;" alt="Aliya GO" />

## 它如何运行

- AliyaRAM 通过 MCP 向 LLMs 暴露数据接口，后端使用 bleve 进行数据的向量索引，LLMs 可以通过 MCP 向 AliyaRAM 写入记忆，并在必要的时候读取记忆。

- 同时 AliyaRAM 通过阿里百炼 RAG 储存了游戏内的完整文本，LLMs 可以通过简单的关键词搜索出任何与之相关的内容。

## 构建指南

编译 AliayRAM，你需要拥有 c 环境和 go 环境。

- go >= 1.24
- gcc >= 13.3

```shell
gcc -v
go version
```

## 未来计划

- [ ] AliyaRAM 目前的能力较为简单，我需要手动整理出完整剧情树结构，而不是依赖 RAG 的简单整理，那样并不完美。之后我会使用 bleve 搭建本地知识库。
- [x] 使用小参数模型与 AstrBot 插件协同工作，提升效率。
