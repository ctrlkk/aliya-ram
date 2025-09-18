# Aliya RAM

Aliya 记忆体，通过最少的 token 实现最真实的 Aliya，它同时也是一个`AstrBot`友好的记忆持久化 MCP Servers

[本地部署](./docs/deploy.md)或加 QQ 群体验：761178785

> aliya 提示词已迁移到 [aliya-prompt](https://github.com/ctrlkk/aliya-prompt)

<div style="text-align:center;">
  <img src="./docs/images/aliya_go.png" width="200px" alt="Aliya GO" />
</div>

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

## ❤️ 赞助与支持

如果你喜欢这个项目，并愿意为我提供继续开发的动力，欢迎赞助一杯咖啡 ☕️ ！

| 平台 | 链接 / 二维码 |
|------|---------------|
| 爱发电 | [![爱发电](https://img.shields.io/badge/爱发电-赞助-red?style=flat-square&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGNUQ1RCIgZD0iTTEyIDIxLjM1bC0xLjQ1LTEuMzJDNS40IDE1LjM2IDIgMTIuMjggMiA4LjUgMiA1LjQyIDQuNDIgMyA3LjUgM2MxLjc0IDAgMy40MS44MSA0LjUgMi4wOUMxMy4wOSAzLjgxIDE0Ljc2IDMgMTYuNSAzIDE5LjU4IDMgMjIgNS40MiAyMiA4LjVjMCAzLjc4LTMuNCA2Ljg2LTguNTUgMTEuNTRMMTIgMjEuMzV6Ii8+PC9zdmc+)](https://afdian.com/a/ctrlkk) |
| 微信支付 | <img src="/docs/pay/wechatpay.png" width="180" alt="微信赞助码" /> |
| 支付宝 | <img src="/docs/pay/alipay.jpg" width="180" alt="支付宝赞助码" /> |

> **感谢你的每一份支持！**  
> 赞助本身不会带来额外权益，但你可以在付款备注留下昵称或名字，我会将其添加到贡献者名单中，以表感谢 🙏

## 未来计划

- [ ] AliyaRAM 目前的能力较为简单，我需要手动整理出完整剧情树结构，而不是依赖 RAG 的简单整理，那样并不完美。之后我会使用 bleve 搭建本地知识库。
- [ ] 使用小参数模型与 AstrBot 插件协同工作，提升效率。

> 当前方案因效果不佳已被废弃

# 相关插件

用于限制用户图片尺寸：
https://github.com/ctrlkk/astrbot_plugin_image_size_limit

让bot可以发送贴图增加互动性：
https://github.com/ctrlkk/astrbot_plugin_meme_manager_lite

显示bot的输入状态：
https://github.com/ctrlkk/astrbot_plugin_input_state_by_nc

下面是一些还处于Private状态的插件：

日程表插件，支持ai自动日程安排与日程冲突处理：
https://github.com/ctrlkk/astrbot_plugin_calendar

专为aliya定制的画图插件，后续也可能会增加泛用性，市面上已经有很多类似插件了：
https://github.com/ctrlkk/astrbot_plugin_canvas_tool

aliya agent与astrbot的绑定：
https://github.com/ctrlkk/astrbot_plugin_aliya_agent_bridge

aliya agent，由nestjs驱动的agent服务，用于管理知识库和记忆：
https://github.com/ctrlkk/aliya-ram-agent

# 赞助者名单，感谢以下用户的鼎力支持，排名不分前后
名单截止至：2025/9/18

暮の色
幻械
Cierra
野兽仙贝
冈上岩村
酸汤点点头
ck567
灰之魔女
泥甘玛

还有更多匿名赞助者，感谢大家的支持
