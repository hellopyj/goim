package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Terry-Mao/goim/dpg/bean"
	"github.com/gomodule/redigo/redis"
)

//xlz===============
const (
	_prefixChatsRoom = _prekey+"chats:room:%s"
	_prefixChatsUser = _prekey+"chats:user:%s"
	_chatlist_length=100
	_user_info=_prekey+"info:%v"

)
//读扩散
func (d *Dao)SaveChat(keys []string,score int64,msgs interface{},prefix string)(err error)  {
	if msgs=="" {
		err=errors.New("msg is empty")
		return
	}
	conn := d.redis.Get()
	defer conn.Close()
	for _,key:=range keys {
		if key=="" {
			continue
		}
		key=fmt.Sprintf(prefix,key)
		if err=conn.Send("ZADD",key,score,msgs);err != nil {
			return
		}
		if err=conn.Send("LTRIM",key,-_chatlist_length,-1);err != nil {
			return
		}
	}
	if err=conn.Flush();err != nil {
		return
	}
	return
}
//读扩散
func (d *Dao)SaveGroupChat(keys []string,score int64,msgs interface{})(err error)  {
	return d.SaveChat(keys,score,msgs,_prefixChatsRoom)
}
//写扩散
func (d *Dao)SaveUserChat(keys []string,score int64,msgs interface{})(err error)  {
	return d.SaveChat(keys,score,msgs,_prefixChatsUser)
}
func (d *Dao)GetConInfo(uid int)(info bean.ConInfo,err error)  {
	conn := d.redis.Get()
	defer conn.Close()
	rkey:=fmt.Sprintf(_user_info,uid)
	infob,err:=redis.Bytes(conn.Do("GET",rkey))
	if err==nil {
		info=bean.ConInfo{}
		json.Unmarshal(infob,&info)
	}
	return
}
func (d *Dao)SetConInfo(uid int,info bean.ConInfo)(err error)  {
	conn := d.redis.Get()
	defer conn.Close()
	rkey:=fmt.Sprintf(_user_info,uid)
	b, _ := json.Marshal(info)
	_,err=conn.Do("SET",rkey,b)
	return
}
//xlz===============
