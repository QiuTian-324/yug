# YUG

## 项目简介
这是一个基于 **Golang** 和 **Electron** 的即时通讯（IM）软件，旨在提供高效的消息交流与办公协作功能。该平台使用 **WebSocket** 进行通信，支持单聊和群聊，消息类型涵盖文本、文件、图片、语音、视频等。此外，项目集成了 **AI 助手**、**工作台**、以及 **审批工作流** 等功能，致力于提升团队协作效率。

## 功能概述
### 即时通讯
- **单聊与群聊**：支持个人和群组之间的实时消息通信。
- **多种消息类型**：
  - 文本消息
  - 文件发送与接收
  - 图片消息
  - 语音消息
  - 视频消息
- **WebSocket 通信**：确保消息传输的实时性与可靠性。

### 工作台
- **任务与申请提交**：用户可以在工作台中提交任务与申请。
- **审批与驳回**：审核人可以对提交的任务或申请进行审批或驳回操作。
- **表单与工作流管理**：
  - 后台管理员可以创建自定义表单。
  - 支持工作流配置，自动化审批流程。

### AI 助手
- **智能助手**：提供基于 AI 的智能问答、文本处理等辅助功能，以提升用户效率。

## 技术栈
- **前端**：
  - Electron
  - Vue.js（可选，用于渲染 UI 界面）
- **后端**：
  - Golang
  - WebSocket（实时通信）
  - RESTful API
- **数据库**：
  - MySQL / PostgreSQL（可根据需求调整）
- **其他**：
  - AI 接口集成（如 OpenAI、GPT 系列 API）
  - 配置文件管理（YAML）

## 系统架构

```
GinFramework/
├── api/                # 控制器层，处理请求与响应
│   ├── chat.go         # 处理聊天相关的API逻辑
│   ├── router.go       # 定义API的路由
│   └── user.go         # 处理用户相关的API逻辑
├── cmd/                # 命令行工具和启动脚本
│   └── start.go        # 项目启动入口，包含初始化配置、启动HTTP服务器等逻辑
├── configs/            # 配置文件目录
│   └── configs.yaml    # 项目的主要配置文件，包括数据库连接、端口号、日志配置等
├── global/             # 全局变量管理
│   ├── constant.go     # 定义项目中使用的常量
│   └── global.go       # 定义全局变量，如配置对象、日志对象等
├── internal/           # 内部模块，通常不会被外部项目引用
│   ├── config/         # 配置管理模块，负责加载和解析配置文件
│   ├── data/           # 数据访问层，处理数据库操作
│   ├── dto/            # 数据传输对象，用于API请求和响应的数据结构
│   ├── handlers/       # 处理业务逻辑的处理器
│   ├── libs/           # 辅助库，包含一些通用工具函数
│   ├── middleware/     # 中间件，处理跨请求的通用逻辑
│   └── services/       # 业务服务层，封装业务逻辑
├── log/                # 日志管理
├── pkg/                # 第三方包或自定义包
├── test/               # 测试文件
├── tmp/                # 临时文件目录
└── utils/              # 工具函数
```


## 环境配置
### 前置条件
- Node.js >= 18.19
- Golang >= 1.22.1
- MySQL
- Git

### 本地开发环境搭建
1. **克隆项目**

   ```bash
   git clone https://github.com/QiuTian-324/FoxIM.git
   cd FoxIM

2. **安装前端依赖**

   ```bash
   cd FoxIM-UI
   pnpm install
   pnpm run dev
   ```
3. **配置后端**
 - 编辑 configs/configs.yaml 文件，填写数据库连接信息和其他配置项。

4. **启动后端服务**

  ```bash
  cd FoxIM-Server
  go mod tidy
  go run main.go
  ```
  
5. **访问应用**

```bash
前端：http://localhost:8080
后端：http://localhost:9000
 ```

## 贡献指南

欢迎大家为本项目贡献代码！请遵循以下步骤：

Fork 本仓库
  - 创建 feature 分支：git checkout -b feature/my-feature
  - 提交修改：git commit -m 'Add my feature'
  - 推送到分支：git push origin feature/my-feature
  - 发起 Pull Request

## 许可证

- 本项目基于 MIT License 开源。

## 联系我

- Email: yjj4872@gmail.com
- **Issue Tracker**: [GitHub Issues](https://github.com/QiuTian-324/FoxIM.git/issues)
