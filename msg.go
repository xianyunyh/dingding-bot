package bot


type MarkDownItem struct {
	Title string `json:"title"`
	Text string `json:"text"`
}

// 文本消息
type TextMsg struct {
	MsgType string `json:"msgtype"`
	Text TextContent `json:"text"`
	At MsgAt `json:"at,omitempty"`
}
type TextContent struct {
	Content string `json:"content"`
}

type MsgAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll bool `json:"isAtAll"`
}
// 链接消息
type LinkMsg struct {
	MsgType string `json:"msgtype"`
	Link LinkItem `json:"link"`
}

type LinkItem struct {
	Title string `json:"title"`
	Text string `json:"text"`
	PicUrl string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

// markdown消息
type MarkDownMsg struct {
	MsgType string `json:"msgtype"`
	Markdown MarkDownItem `json:"markdown"`
	At MsgAt `json:"at,omitempty"`
}

type ActionCardMsg struct {
	MsgType string `json:"msgtype"`
	ActionCard ActionCardItem `json:"actionCard"`
}
type ActionCardItem struct {
	Title string `json:"title"`
	Text string `json:"text"`
	SingleTitle string `json:"singleTitle,omitempty"`
	SingleURL string `json:"singleURL,omitempty"`
	HideAvatar string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	Btns []CardBtn `json:"btns,omitempty"`
}

type CardBtn struct {
	Title string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// feed流消息
type FeedMsg struct {
	MsgType string `json:"msgtype"`
	FeedCard FeedCardItem `json:"feedCard"`
}

type FeedCardItem struct {
	Links  []FeedLink `json:"links"`
}
type FeedLink struct {
	Title string `json:"title"`
	MessageUrl string `json:"messageURL"`
	PicUrl string `json:"picURL"`
}