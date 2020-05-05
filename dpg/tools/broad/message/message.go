package message

type Message struct {
	Name string `json:"name,omitempty"`
	Time int `json:"time,omitempty"`
	Msg interface{} `json:"msg,omitempty"`
	Type string  `json:"type,omitempty"`//0.普通消息,1.是红包,2是提示
	Avatar string `json:"avatar,omitempty"` //头像
	Uid int `json:"uid,omitempty"` //用户ID
	Aid int `json:"aid,omitempty"` //艾特对象
}
const(
	TXT="text" //文字
	IMG="image"//图片
	BET="bet"
)