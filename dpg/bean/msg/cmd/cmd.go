package cmd

import (
	"encoding/json"
	"github.com/Terry-Mao/goim/dpg/bean/msg/cmd/toast"
	"github.com/json-iterator/go"
	"time"
)

type CmdRaw struct {
	Code int `json:"code"`
	Content jsoniter.RawMessage `json:"content"`
	Extra  jsoniter.RawMessage `json:"extra"`
}

func (cm * CmdRaw)GetContent(c interface{})(error)  {
	return jsoniter.Unmarshal(cm.Content,&c)
}
func (cm * CmdRaw)GetExtra(c interface{})(error)  {
	return jsoniter.Unmarshal(cm.Extra,&c)
}

type Cmd struct {
	Code int `json:"code"`
	Content interface{} `json:"content,omitempty"`
	Extra interface{} `json:"extra,omitempty"`
	CreateTime 	int64 `json:"createTime"`//创建时间
}
func (cm *Cmd)SetCode(code int)(newcm *Cmd)  {
	cm.Code=code
	return cm
}
func (cm *Cmd)SetContent(content interface{})(newcm *Cmd)  {
	switch content.(type) {
	case toast.Toast:
		cm.SetCode(CMD_TOAST)
	default:
		//cm.SetCode(0)
	}
	cm.Content=content
	return cm
}
func (cm *Cmd)SetExtra(extra interface{})(newcm *Cmd)  {
	cm.Extra=extra
	return cm
}
func (cm *Cmd)SetCreateTime(crtime int64) (newcm *Cmd){
	cm.CreateTime=crtime
	return cm
}
func (cm *Cmd)Init() (newcm *Cmd){
	cm.SetCreateTime(time.Now().Unix())
	return cm
}
//格式化
func (cm *Cmd)String() (cmstr string){
	cmstr,_=jsoniter.MarshalToString(cm)
	return
}
func (c *Cmd)Bytes() (bs []byte){
	bs,_=jsoniter.Marshal(c)
	return
}
func Decode(b []byte)(cmd Cmd,err error)  {
	err=json.Unmarshal(b,&cmd)
	return
}
func InitFromByte(b []byte)(cmd Cmd,err error)  {
	cmd,err=Decode(b)
	if err!=nil {
		return
	}
	cmd.Init()
	return
}
func Init() (*Cmd){
	cm:=Cmd{}
	cm.Init()
	return &cm
}
