package chat

import (
	"encoding/json"
	"fmt"
	"github.com/Terry-Mao/goim/dpg/tools/snowflake"
	"github.com/json-iterator/go"
	"sort"
	"time"
)
var _snode *snowflake.Node
func init()  {
	_snode, _ = snowflake.NewNode(1)
}
//错误
const (
	ERROR_FORMAT="格式错误"
)
type Chat struct {
	ID         	int64 `json:"id"`//全局唯一ID
	MID        	int64 `json:"mid"` //消息ID
	From       	string `json:"from"`//来源ID
	To         	string `json:"to"`//目标ID
	MsgType    	int `json:"msgType"`//消息类型
	ChatType   	int `json:"chatType"`
	Group 	   	string `json:"group,omitempty"`//群组ID
	Content    	interface{}`json:"content,omitempty"`//内容
	Extras     	interface{} `json:"extras,omitempty"`//扩展
	Ship		bool `json:"-"` //不存储
	CreateTime 	int64 `json:"createTime"`//创建时间
	UpdateTime 	int64 `json:"updateTime"`//更新时间
}
//id生成
func (cm *Chat)_id() (id int64){
	return _snode.Generate().Int64()
}
//设置内容
func (cm *Chat)SetContent(content interface{}) (newcm *Chat){
	cm.Content=content
	return cm
}
//设置扩展
func (cm *Chat)SetExtra(extra interface{}) (newcm *Chat){
	cm.Extras=extra
	return cm
}
//msgid生成
func (cm *Chat)SetID(id int64) (newcm *Chat){
	cm.ID=id
	return cm
}
//createtime生成
func (cm *Chat)SetCreateTime(crtime int64) (newcm *Chat){
	cm.CreateTime=crtime
	return cm
}
//createtime生成
func (cm *Chat)SetUpdateTime(uptime int64) (newcm *Chat){
	cm.UpdateTime=uptime
	return cm
}
//设置聊天类别
func (cm *Chat)SetChatType(chat_type int) (newcm *Chat){
	cm.ChatType=chat_type
	return cm
}
//设置信息类别
func (cm *Chat)SetMsgType(msg_type int) (newcm *Chat){
	cm.MsgType=msg_type
	return cm
}
//初始化
func (cm *Chat)Init() (newcm *Chat){
	cm.SetID(cm._id())
	cm.SetCreateTime(time.Now().Unix())
	return cm
}
//格式化
func (cm *Chat)String() (cmstr string){
	cmstr,_=jsoniter.MarshalToString(cm)
	return
}
//是否是单聊
func (cm *Chat)IsUserChat() (bool){
	return cm.ChatType==CHAT_USER&&cm.To!=""
}
//是否跳过存储
func (cm *Chat)IsSkip() (bool){
	return cm.Ship
}
//是否是单聊
func (cm *Chat)IsGroupChat() (bool){
	return cm.ChatType==CHAT_GROUP&&cm.Group!=""
}
//是否可以发送
func (cm *Chat)CanSend() (bool){
	return true
}
//格式化
func (cm *Chat)GetRKeys() (cmstr []string){
	switch cm.ChatType {
	case CHAT_USER :
		ids:=[]string{cm.From,cm.To}
		return []string{ChatKey(ids)}
	case CHAT_GROUP:
		return []string{cm.Group}
	}
	return
}
func (cm *Chat)Bytes() (bs []byte){
	bs,_=jsoniter.Marshal(cm)
	return
}
func Decode(b []byte)(msg Chat,err error)  {
	err=json.Unmarshal(b,&msg)
	//if msg.ChatType==CHAT_USER{
	//	msg.Group=""
	//}
	return
}
func InitFromByte(b []byte)(msg Chat,err error)  {
	msg,err=Decode(b)
	if err!=nil {
		return
	}
	msg.Init()
	return
}
func Init() (*Chat){
	cm:=Chat{}
	cm.Init()
	return &cm
}
//聊天key
func ChatKey(ids [] string) (onekey string){
	sort.Strings(ids)
	onekey=fmt.Sprintf("%v-%v",ids[0],ids[1])
	return
}
