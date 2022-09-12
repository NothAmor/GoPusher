# GoPusher  

基于Golang的消息推送轮子  

## 支持推送服务如下：  
* ServerChan 微信公众号Server酱  

  请求结构体如下，其中必填项为Key, Title

  ```go
  type ServerChanRequestStruct struct {
  	Key     string
  	Title   string
  	Desp    string `default: ""`
  	Short   string `default: ""`
  	Channel string `default: ""`
  	Openid  string `default: ""`
  }
  ```

  使用方法如下：

  ```go
  package main
  
  import (
  	"fmt"
  
  	GoPusher "git.nothamor.cn/NothAmor/GoPusher"
  	structs "git.nothamor.cn/NothAmor/GoPusher/structs"
  )
  
  func main() {
  	// 初始化请求结构体
  	serverChanParams := structs.ServerChanRequestStruct{
  		// Server酱请求Key，获取方式：https://sct.ftqq.com/
  		Key: "Key",
  		// 推送标题
  		Title: "Title",
  		// 推送正文内容
  		Desp: "Content",
  	}
  	// 进行推送
  	pusher, err := GoPusher.ServerChan(serverChanParams)
  	// 错误处理
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(pusher)
  }
  ```

* SMTP邮件服务

  请求结构体如下，其中除Nickname外都是必填项

  ```go
  type SmtpRequestStruct struct {
  	Host     string
  	Account  string
  	Password string
  	Port     int    `default: 0`
  	Nickname string `default: "GoPusher"`
  	MailType string
  	Sender   string
  	SendTo   []string
  	Title    string
  	Content  string
  }
  ```
  
  使用demo如下：
  
  ```go
  package main
  
  import (
  	"fmt"
  
  	GoPusher "git.nothamor.cn/NothAmor/GoPusher"
  	structs "git.nothamor.cn/NothAmor/GoPusher/structs"
  )
  
  func main() {
  	mailParams := structs.SmtpRequestStruct{
  		Host:     "Host",     // SMTP主机
  		Account:  "Account",  // SMTP用户名
  		Password: "Password", // SMTP密码
  		Port:     25,         // SMTP端口
  		MailType: "html",     // 邮件格式，支持值：plain，html
  		Sender:   "Sender",   // 发送者，可以与Account值相同
  		SendTo: []string{ // 接收者，数组形式
  			"Receiver1",
  			"Receiver2",
  		},
  		Title:   "Title",   // 邮件标题
  		Content: "Content", // 邮件正文，如果正文是html内容，请将MailType改为html，否则不会解析
  	}
  	// 发送邮件
  	smtpResponse, err := GoPusher.Smtp(mailParams)
  	if err != nil {
  		panic(err)
  	}
  	// 打印结果
  	fmt.Println(smtpResponse)
  }
  ```
  
  