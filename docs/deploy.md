# 部署指南

Ubuntu24

## 部署前准备

首先在[阿里百炼](https://bailian.console.aliyun.com/?tab=app#/knowledge-base)中创建一个知识库，使用默认配置，导入对话数据，创建成功后进行测试。
关于[RAG](https://bailian.console.aliyun.com/?spm=a2c4g.11186623.0.0.22b32562p5Vmiq&tab=doc#/doc/?type=app&url=https%3A%2F%2Fhelp.aliyun.com%2Fdocument_detail%2F2807740.html&renderType=iframe)

[知识库数据](./对话.csv)

## 部署

### 临时环境变量

在软件根目录创建.env文件，内容参考[.env.example](./.env.example)

### 运行

运行编译后的二进制文件

### 使用

AliyaRAM通过Streamable HTTP对外提供服务，下面是在AstrBot中的配置案例

```json
{
  "url": "http://localhost:8080/mcp",
  "transport": "streamable_http"
}
```
