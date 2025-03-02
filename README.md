# Go-Dmsoft

大漠插件的 Go 语言封装，提供完整的接口封装和多线程支持。支持窗口操作、键鼠模拟、图色处理和 AI 功能。

## 版本信息

- 大漠插件版本: 7.2450 
  - 支持 Win11 26100(24h2) 版本
  - 驱动和 LoadAi 功能支持最新系统

## 下载地址

- 蓝奏云: https://wwpk.lanzouo.com/iZobX2hu84wj
- 123云盘: https://vip.123pan.cn/1824865303/dm/dm%E5%AF%86%E7%A0%811234.rar
- 百度网盘: https://pan.baidu.com/s/1WT0inczc0CFWuQ5pDR-iyg?pwd=1234 (提取码: 1234)

## ⚠️ 重要提示

大漠插件为32位程序，使用前必须设置：

$env:GOARCH="386"  # PowerShell
set GOARCH=386     # CMD
```

## 功能特性

- 窗口操作和绑定
- 键鼠模拟
- 图色找点找图
- 屏幕截图
- AI 目标检测
- 多线程支持
```
## 项目结构

```
dmsoft/
├── dll/                   # DLL 文件目录
│   ├── DmReg.dll          # 注册模块
│   └── dm.dll             # 主模块
├── example/               # 示例代码
│   ├── basic/             # 基础使用示例
│   └── multi_window/      # 多线程示例
├── dmsoft.go              # 核心实现
├── ai.go                  # AI 相关功能
└── go.mod                 # 模块定义
```

## 安装

```bash
go get github.com/angerige/go-dmsoft
```

## 使用示例

```go
package main

import (
    "log"
    "github.com/angerige/go-dmsoft"
)

func main() {
    dm, err := dmsoft.CreateDm(&dmsoft.DmConfig{
        RegCode:   "你的注册码",   // 必须提供
        ExtraCode: "你的附加码",  // 必须提供
    })
    if err != nil {
        log.Fatal(err)
    }
    defer dm.Release()

    log.Printf("插件版本:%s", dm.Ver())
}
```

## 注意事项

1. 32位编译环境
```powershell
$env:GOARCH="386"  # PowerShell
set GOARCH=386     # CMD
```

2. 管理员权限运行
3. DLL文件放在 dll 目录下

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request

## 联系方式

如有问题或建议，请通过以下方式联系：
- GitHub Issues
- Email: your.email@example.com