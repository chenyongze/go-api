package controllers

import (
"fmt"

"github.com/silenceper/wechat"
"github.com/silenceper/wechat/message"
)


type WeChatController struct {
	BaseController
}

func (self *WeChatController) Run() {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "your app id",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(self.Ctx.Request, self.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("x")
	//发送回复的消息
	server.Send()

}
