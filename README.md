# ztctl - ZeroTier网络管理工具

`ztctl`是一个命令行工具，用于管理ZeroTier网络。它可以帮助你创建、删除、修改网络，以及管理网络成员。

## 功能特点

- 网络管理：创建、列出、删除ZeroTier网络
- 成员管理：查看网络成员、授权/取消授权成员
- 友好的命令行界面
- 清晰的输出格式

## 环境要求

- Go 1.16或更高版本
- ZeroTier Central API Token

## 安装

```bash
go install github.com/lfshao/ztctl@latest
```

## 配置

在使用`ztctl`之前，你需要设置ZeroTier Central API Token。

1. 登录[ZeroTier Central](https://my.zerotier.com/)
2. 在账户设置中生成API Token
3. 设置环境变量：

```bash
export ZEROTIER_CENTRAL_TOKEN="your-api-token"
```

## 使用方法

### 网络管理

1. 列出所有网络：
```bash
ztctl network list
```

2. 创建新网络：
```bash
ztctl network create --name "我的网络" --description "这是一个测试网络"
```

3. 删除网络：
```bash
ztctl network delete --id "network-id"
```

### 成员管理

1. 查看网络成员：
```bash
ztctl network members --id "network-id"
```

2. 授权成员加入网络：
```bash
ztctl network authorize --network-id "network-id" --member-id "member-id"
```

3. 取消成员授权：
```bash
ztctl network deauthorize --network-id "network-id" --member-id "member-id"
```

## 输出示例

1. 列出网络：
```
网络ID                    网络名称    成员数量
abc123456789             测试网络    3
def987654321             开发网络    5
```

2. 查看成员：
```
网络: 测试网络 (abc123456789)
成员ID              名称       IP地址          状态    最后在线时间
123456789          设备1    10.147.1.1      在线    2023-12-01 12:00:00
987654321          设备2    10.147.1.2      离线    2023-12-01 10:30:00
```

## 错误处理

工具会提供清晰的错误信息和建议：

- API Token无效时会提示检查token
- 网络ID不存在时会提示检查ID
- 网络连接问题时会提供相应的故障排除建议

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License