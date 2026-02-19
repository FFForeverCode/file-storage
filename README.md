# file-storage

轻量级 P2P 分布式文件存储服务器，基于 Go 语言实现。

## 项目简介

file-storage 是一个支持节点间文件共享和传输的分布式存储系统。通过 P2P 网络架构，实现节点之间的文件存储和检索功能，适合构建小型分布式文件存储网络。

## 核心特性

- **P2P 网络通信**：基于 TCP 协议的节点间通信
- **分布式存储**：支持多节点文件存储和检索
- **灵活路径转换**：支持自定义文件存储路径转换逻辑
- **消息编解码**：支持 GOB 和自定义消息解码器
- **高性能日志**：基于 Uber Zap 的结构化日志

## 项目结构

```
file-storage/
├── main.go              # 程序入口，初始化日志和启动服务
├── server.go            # 文件服务器核心逻辑
├── store.go             # 存储数据结构定义
├── p2p/                 # P2P 网络通信模块
│   ├── transport.go     # 传输层接口定义
│   ├── tcp_transport.go # TCP 传输实现
│   ├── message.go       # RPC 消息定义
│   ├── encoding.go      # 消息编解码器
│   └── handleshake.go   # 握手协议
├── logger/              # 日志模块
│   └── log.go          # Zap 日志封装
├── go.mod              # Go 模块定义
├── Makefile            # 构建脚本
└── README.md           # 项目文档
```

## 模块说明

### 核心模块

#### FileServer (server.go)
文件服务器核心实现，支持：
- 节点 ID 生成和管理
- 节点间连接管理
- 文件存储和检索消息处理
- 自定义路径转换函数

#### Store (store.go)
存储抽象层：
- 可配置的存储根目录
- 灵活的路径转换策略
- 默认路径转换实现

#### P2P 通信 (p2p/)
- **Transport 接口**：定义传输层抽象
- **TCPTransport**：TCP 协议实现
- **Peer 接口**：节点连接抽象
- **编解码器**：GOB/Default 解码器支持

#### Logger (logger/)
- 基于 Uber Zap 的生产级日志
- 支持 Info/Error 等多级别日志
- 自动调用位置追踪

## 快速开始

### 环境要求

- Go 1.25+
- 网络：支持 TCP 通信

### 安装依赖

```bash
go mod download
```

### 构建项目

```bash
make build
```

### 运行服务

```bash
make run
# 或直接运行
./bin/fs
```

### 运行测试

```bash
make test
```

## 开发计划

- [ ] 完善 FileServer 的 GET 方法实现
- [ ] 完善 FileServer 的 Store 方法实现
- [ ] 添加文件上传/下载 HTTP 接口
- [ ] 实现文件分片存储
- [ ] 添加节点发现机制
- [ ] 实现数据冗余备份
- [ ] 添加 API 文档

## 技术栈

- **语言**：Go 1.25
- **日志**：go.uber.org/zap v1.27.1
- **错误处理**：go.uber.org/multierr v1.10.0

## 消息协议

### MessageStoreFile
文件存储消息：
```
ID   string  // 节点 ID
Key  string  // 文件键
Size int64   // 文件大小
```

### MessageGetFile
文件获取消息：
```
ID  string  // 节点 ID
Key string  // 文件键
```

## License

MIT

