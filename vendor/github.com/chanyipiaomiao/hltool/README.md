# hltool

Go 开发常用工具库


# 安装

使用golang官方 dep 管理依赖
```go
go get github.com/chanyipiaomiao/hltool
```

# 功能列表
- [钉钉机器人通知](#钉钉机器人通知)
- [发送邮件](#发送邮件)
- [JWT Token生成解析](#jwt-token生成解析)
- [Log库](#log库)
- [BoltDB嵌入式KV数据库](#boltdb嵌入式kv数据库)
- [检测图片类型](#检测图片类型)
- [图片转[]byte](#图片转byte数组)
- [[]byte转换为png/jpg](#byte数组转换为png-jpg)

### 钉钉机器人通知
```go
import (
	"log"
	"github.com/chanyipiaomiao/hltool"
)

dingtalk := hltool.NewDingTalkClient("钉钉机器URL", "消息内容", "text|markdown")
ok, err := hltool.SendMessage(dingtalk)
if err != nil {
	log.Fatalf("发送钉钉通知失败了: %s", err)
}

```
[返回到目录](#功能列表)

### 发送邮件
```go
import (
	"log"
	"github.com/chanyipiaomiao/hltool"
)

username := "xxxx@xxx.com"
host := "smtp.exmail.qq.com"
password := "password"
port := 465

subject := "主题"
content := "内容"
contentType := "text/plain|text/html"
attach := "附件路径" 或者 ""
to := []string{"xxx@xxx.com", "xxx@xx.com"}
cc := []string{"xxx@xxx.com", "xxx@xx.com"}

message := hltool.NewEmailMessage(username, subject, contentType, content, attach, to, cc)
email := hltool.NewEmailClient(host, username, password, port, message)
ok, err := hltool.SendMessage(email)
if err != nil {
	log.Fatalf("发送邮件失败了: %s", err)
}
```
[返回到目录](#功能列表)

### JWT Token生成解析
```go
import (
	"fmt"
	"log"

	"github.com/chanyipiaomiao/hltool"
)

func main() {

	// 签名字符串
	sign := "fDEtrkpbQbocVxYRLZrnkrXDWJzRZMfO"

	token := hltool.NewJWToken(sign)

	// -----------  生成jwt token -----------
	tokenString, err := token.GenJWToken(map[string]interface{}{
		"name": "root",
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(tokenString)

	// -----------  解析 jwt token -----------
	r, err := token.ParseJWToken(tokenString)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(r)

}

```
输出
```shell
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicm9vdCJ9.NJMXxkzdBBWrNUO5u2oXFLU9FD18TWiXHqxM2msT1x0

map[name:root]
```
[返回到目录](#功能列表)

### Log库

- 支持按天分割日志
- 不同级别输出到不同文件
- 支持 文本/json日志类型,默认是json类型
- 设置日志最大保留时间

```go
import (
	"github.com/chanyipiaomiao/hltool"
)

func main() {
	
	commonFields := map[string]interface{}{
		"name": "zhangsan",
		"age":  "20",
	}

	hlog, _ := hltool.NewHLog("./test.log")
	// hlog.SetLevel("debug") debug|info|warn|error|fatal|panic
	logger, _ := hlog.GetLogger()

	// Info Warn 会输出到不同的文件
	logger.Info(commonFields, "测试Info消息")
	logger.Warn(commonFields, "测试Warn消息")
	
	// Error Fatal Panic 会输出到一个文件
	logger.Error(commonFields, "测试Error消息")
	logger.Fatal(commonFields, "测试Fatal消息")
	logger.Panic(commonFields, "测试Panic消息")
}
```
日志文件内容:
```shell
{"age":"20","level":"debug","msg":"测试Debug消息","name":"zhangsan","time":"2018-02-08 21:28:29"}
{"age":"20","level":"info","msg":"测试Info消息","name":"zhangsan","time":"2018-02-08 21:28:29"}
{"age":"20","level":"warning","msg":"测试Warn消息","name":"zhangsan","time":"2018-02-08 21:28:29"}
{"age":"20","level":"error","msg":"测试Error消息","name":"zhangsan","time":"2018-02-08 21:28:29"}
```
[返回到目录](#功能列表)

### BoltDB嵌入式KV数据库
```go
import (
	"log"

	"github.com/chanyipiaomiao/hltool"
)

func main() {

	// 数据库文件路径 表名
	db, err := hltool.NewBoltDB("./data/app.db", "token")
	if err != nil {
		log.Fatalf("%s", err)
	}
	db.Set(map[string][]byte{
		"hello": []byte("world"),
		"go":    []byte("golang"),
	})
	r, err := db.Get([]string{"hello", "go"})
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Println(r)
}
```
[返回到目录](#功能列表)


### 检测图片类型

```go
package main

import (
	"fmt"

	"github.com/chanyipiaomiao/hltool"
)

func main() {

	bytes, _ := hltool.ImageToBytes("1.png")
	fmt.Println(hltool.ImageType(bytes))

}
```
输出结果:

```go
image/png
```

[返回到目录](#功能列表)

### 图片转byte数组

```go
package main

import (
	"fmt"

	"github.com/chanyipiaomiao/hltool"
)

func main() {

	bytes, err := hltool.ImageToBytes("1.png")
	if err != nil {
		fmt.Println(err)
	}

}
```

[返回到目录](#功能列表)

### byte数组转换为png jpg
```go
package main

import (
	"fmt"

	"github.com/chanyipiaomiao/hltool"
)

func main() {

	bytes, err := hltool.ImageToBytes("1.png")
	if err != nil {
		log.Fatalln(err)
	}

	err = hltool.BytesToImage(bytes, "111.png")
	if err != nil {
		log.Fatalln(err)
	}

}
```

[返回到目录](#功能列表)