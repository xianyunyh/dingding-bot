package bot

import (
	"net/http"
	"time"
	"crypto/tls"
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"crypto/hmac"
	"crypto/sha256"
	"strconv"
	"encoding/base64"
	"io/ioutil"
	"log"
)

const API_URI = "https://oapi.dingtalk.com/robot/send"

type DingdingBot struct {
	Key   string
	Token string
}

type Response struct {
	ErrorCode int `json:"errcode"`
	ErrorMsg string `json:"errmsg"`
}
func NewBot(token,key string ) *DingdingBot {
	return &DingdingBot{
		Key:   key,
		Token: token,
	}
}

// http 请求
func (self *DingdingBot) req(data []byte) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	if self.Token == "" {
		return errors.New("access_token不能为空")
	}
	client.Transport = &http.Transport{
		DisableKeepAlives:   true,
		TLSHandshakeTimeout: 1 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true, Certificates: nil},
	}
	params := url.Values{}
	params.Add("access_token",self.Token)
	if self.Key != "" {
		var millisecond int64 = int64(time.Now().Unix() * 1000)
		params.Add("sign",self.sign(millisecond))
		params.Add("timestamp",strconv.FormatInt(millisecond,10))

	}
	reqUrl  := API_URI + "?"+params.Encode()
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type","application/json")
	rep, err := client.Do(req)
	if err != nil {
		return err
	}
	body,err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return err
	}
	defer func() {
		rep.Body.Close()
	}()
	log.Println("req",string(data))
	log.Println("resp",string(body))
	resp := &Response{}
	err = json.Unmarshal(body,resp)
	if err != nil {
		return errors.New("json解码失败")
	}
	if resp.ErrorCode > 0 {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}

// Send 发送消息
func (self *DingdingBot)Send(msg interface{}) error  {
	sendStrctData := func(data interface{}) error {
		postBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return self.req(postBytes)
	}
	switch msg.(type) {
	case TextMsg:
		temp := msg.(TextMsg)
		temp.MsgType = "text"
		return sendStrctData(temp)
	case LinkMsg:
		temp := msg.(LinkMsg)
		temp.MsgType = "link"
		return sendStrctData(temp)
	case MarkDownMsg:
		temp := msg.(MarkDownMsg)
		temp.MsgType = "markdown"
		return sendStrctData(temp)
	case FeedMsg:
		temp := msg.(FeedMsg)
		temp.MsgType = "feedCard"
		return sendStrctData(temp)
	case ActionCardMsg:
		temp := msg.(ActionCardMsg)
		temp.MsgType = "actionCard"
		return sendStrctData(temp)
	case LinkItem:
		temp := LinkMsg{Link: msg.(LinkItem),MsgType:"link"}
		return sendStrctData(temp)
	case MarkDownItem:
		temp := MarkDownMsg{Markdown: msg.(MarkDownItem),MsgType:"markdown"}
		return sendStrctData(temp)
	case ActionCardItem:
		temp := ActionCardMsg{ActionCard:msg.(ActionCardItem),MsgType:"actionCard"}
		return sendStrctData(temp)
	case FeedCardItem:
		temp := FeedMsg{FeedCard:msg.(FeedCardItem),MsgType:"feedCard"}
		return sendStrctData(temp)
	case FeedLink:
		temp := FeedMsg{MsgType:"feedCard"}
		temp.FeedCard.Links = []FeedLink{msg.(FeedLink)}
		return sendStrctData(temp)
	case []byte:
		temp := msg.([]byte)
		return self.req(temp)
	case string:
		temp := []byte(msg.(string))
		return self.req(temp)
	}
	return errors.New("unknown support type")
}

func (self *DingdingBot) SendText(text string ) error {
	msg := TextMsg{
		MsgType:"text",
		Text: TextContent{
			Content:text,
		},
	}
	return self.Send(msg)
}
//签名
func (self *DingdingBot) sign(millisecond int64) string {

	secret := []byte(self.Key)
	message := []byte(strconv.FormatInt(millisecond,10)+"\n"+self.Key)
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)
	// to base64
	return  base64.StdEncoding.EncodeToString(hash.Sum(nil))
}