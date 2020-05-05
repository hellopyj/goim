package dmsg

import (
	"github.com/Terry-Mao/goim/dpg/bean/msg/chat"
	"github.com/Terry-Mao/goim/dpg/bean/msg/cmd"
)

type Gmsg struct {
	Chat *chat.ChatMsg `json:"msg,omitempty"`
	Cmd *cmd.Cmd     `json:"cmd,omitempty"`
}

func (g *Gmsg)SetChat(chat *chat.ChatMsg) *Gmsg{
	g.Chat=chat
	return g
}

func (g *Gmsg)SetCmd(cmd *cmd.Cmd) *Gmsg{
	g.Cmd=cmd
	return g
}
