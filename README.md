# Aliya RAM

<strong>Aliya记忆体，通过最少的token实现最真实的Aliya，它同时也是一个`AstrBot`友好的记忆持久化智能体</strong>

本地[部署](./docs/deploy.md)或加QQ群体验：761178785

![Aliya GO](./docs/aliya_go.png)

## 它如何运行

- AliyaRAM通过MCP向大模型暴露数据接口，后端使用bleve进行数据的向量索引，LLMs可以通过MCP向AliyaRAM写入记忆，并在必要的时候读取记忆。
- 同时AliyaRAM通过阿里百炼RAG储存了游戏内的完整文本，LLMs可以通过简单的关键词搜索出任何与之相关的内容。

## 构建指南

AliayRAM使用纯go实现，直接编译即可。

## 未来计划

- AliyaRAM目前的能力较为简单，我需要手动整理出完整剧情树结构，而不是依赖RAG的简单整理，那样并不完美。之后我会使用bleve搭建本地知识库。
