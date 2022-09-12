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
  
      push     "git.nothamor.cn/NothAmor/GoPusher/ServerChan"
      structs  "git.nothamor.cn/NothAmor/GoPusher/structs"
  )
  
  func main() {
    // 初始化请求结构体
  	serverChanParams := structs.ServerChanRequestStruct{
      // Server酱请求Key，获取方式：https://sct.ftqq.com/
  		Key:   "key",
      // 推送标题
  		Title: "Title",
      // 推送正文内容
  		Desp:  "Content",
  	}
    // 进行推送
  	pusher, err := push.ServerChan(serverChanParams)
    // 错误处理
  	if err != nil {
  		panic(err)
  	}
      fmt.Println(pusher)
  }
  ```

  