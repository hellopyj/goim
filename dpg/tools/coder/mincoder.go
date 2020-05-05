package coder

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)
type (
	messageType uint8
)

func (m messageType) String() string {
	return strconv.Itoa(int(m))
}
func (m messageType) Name() string {
	switch m {
	case messageTypeString:
		return "string"
	case messageTypeInt:
		return "int"
	case messageTypeBool:
		return "bool"
	case messageTypeBytes:
		return "[]byte"
	case messageTypeJSON:
		return "json"
	default:
		return "Invalid(" + m.String() + ")"
	}
}
const (
	messageTypeString messageType = iota
	messageTypeInt
	messageTypeBool
	messageTypeBytes
	messageTypeJSON
)

type coderdata struct {
	Event string
	Data cdata
}

func (cd *coderdata)RawStr()string  {
	return string(cd.Data.Raw)
}
type cdata struct {
	Raw []byte
	Type int
}
var messagehead=[]byte("hello:")
var headlen=len(messagehead)
var messageSeparatorByte=byte(';')
var maxhead=50
func Encode(evt string,data interface{})(ret string)  {
	b:=bytes.Buffer{}
	b.Write(messagehead)
	b.WriteString(evt)
	b.WriteByte(messageSeparatorByte)
	switch v := data.(type) {
	case string:
		b.WriteString(messageTypeString.String())
		b.WriteByte(messageSeparatorByte)
		b.WriteString(v)
	case int:
		b.WriteString(messageTypeInt.String())
		b.WriteByte(messageSeparatorByte)
		binary.Write(&b, binary.LittleEndian, v)
	case bool:
		b.WriteString(messageTypeBool.String())
		b.WriteByte(messageSeparatorByte)
		if v {
			b.Write([]byte("true"))
		} else {
			b.Write([]byte("false"))
		}
	case []byte:
		b.WriteString(messageTypeBytes.String())
		b.WriteByte(messageSeparatorByte)
		b.Write(v)
	default:
		res, err := json.Marshal(data)
		if err != nil {
			return ""
		}
		b.WriteString(messageTypeJSON.String())
		b.WriteByte(messageSeparatorByte)
		b.Write(res)
	}
	return b.String()
}
func Decode(content []byte)(cd *coderdata,err error)  {
	if len(content) < headlen {
		return nil,errors.New("长度太小")
	}
	defer func() {
		if e:=recover();e!=nil {
			err=fmt.Errorf("错误:%v",e)
		}
	}()
	cd=new(coderdata)
	count:=0
	mark:=headlen
	lcontent:=len(content)
	for index:=headlen;index<lcontent;index++{
		if  index>maxhead{
			break
		}
		if content[index]==messageSeparatorByte {
			if count==0 {
				cd.Event=string(content[mark:index])
			}else if count==1 {
				cd.Data.Type,_=strconv.Atoi(string(content[mark+1:index]))
				cd.Data.Raw=content[index+1:]
				break
			}
			mark=index
			count++
		}
	}
	return
}
