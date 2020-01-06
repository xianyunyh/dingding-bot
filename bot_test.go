package bot

import (
	"testing"
)
// 用户的access_token
const token = ""
// 加密验证签名的密钥
const key = ""

func TestMarkDownMsg(t *testing.T) {
	bot := NewBot(token,key)
	MarkItem := MarkDownItem {
		Title:"hello this is markdown",
		Text:"### 111 \n > 111",
	}
	err := bot.Send(MarkItem)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendTextMsg(t *testing.T)  {
	bot := NewBot(token,key)
	err := bot.SendText("hello world")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendLinkMsg(t *testing.T)  {
	msg := LinkItem{
		Text:"hello link",
		Title:"this is title",
		PicUrl:"https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.197/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png",
		MessageUrl:"https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.197/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png",
	}
	bot := NewBot(token,key)
	err := bot.Send(msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionCardMsg(t *testing.T)  {
	bot := NewBot(token,key)
	cardMsg := ActionCardItem{
		Title:"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
		Text:`### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`,
		SingleTitle:"阅读全文",
		SingleURL:"https://www.dingtalk.com/",
		HideAvatar:"0",
		BtnOrientation:"0",
	}
	if err := bot.Send(cardMsg);err != nil {
		t.Fatal(err)
	}
}

func TestSendFeedMsg(t *testing.T)  {
	bot := NewBot(token,key)
	links := make([]FeedLink,0)
	links = append(links,FeedLink{
		Title:"时代的火车向前开",
		MessageUrl:"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==",
		PicUrl:"https://www.dingtalk.com/",
	})
	links = append(links,FeedLink{
		Title:"时代的火车向前开2",
		MessageUrl:"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==",
		PicUrl:"https://img.alicdn.com/tfs/TB1nhWCiBfH8KJjy1XbXXbLdXXa-547-379.png",
	})
	msg := FeedCardItem{Links:links}
	if err := bot.Send(msg);err != nil {
		t.Fatal(err)
	}
}
func TestSendFeedMsg2(t *testing.T)  {
	bot := NewBot(token,key)
	msg := FeedLink{
		Title:"时代的火车向前开",
		MessageUrl:"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==",
		PicUrl:"https://img.alicdn.com/tfs/TB1nhWCiBfH8KJjy1XbXXbLdXXa-547-379.png",
	}
	if err := bot.Send(msg);err != nil {
		t.Fatal(err)
	}
}
