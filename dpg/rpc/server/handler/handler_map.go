package handler

import (
	"fmt"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	"github.com/bilibili/discovery/naming"
	"net/url"
	"time"
	log "github.com/golang/glog"
)

const(
	appid="dp.deal"
)
type HandlerMap struct {
	c *conf.Config
}

func (h *HandlerMap) watchHandler() {
	dis := naming.New(h.c.Discovery)
	resolver := dis.Build(appid)
	event := resolver.Watch()
	select {
	case _, ok := <-event:
		if !ok {
			panic("handler init failed")
		}
		if ins, ok := resolver.Fetch(); ok {
			if err := h.newAddress(ins.Instances); err != nil {
				panic(err)
			}
			log.Infof("handler init newAddress:%+v", ins)
		}
	case <-time.After(10 * time.Second):
		log.Error("handler init instances timeout")
	}
	go func() {
		for {
			if _, ok := <-event; !ok {
				log.Info("handler exit")
				return
			}
			ins, ok := resolver.Fetch()
			if ok {
				if err := h.newAddress(ins.Instances); err != nil {
					log.Errorf("handler newAddress(%+v) error(%+v)", ins, err)
					continue
				}
				log.Infof("handler change newAddress:%+v", ins)
			}
		}
	}()
}

func (h *HandlerMap) newAddress(insMap map[string][]*naming.Instance) error {
	ins := insMap[h.c.Env.Zone]
	if len(ins) == 0 {
		return fmt.Errorf("handler instance is empty")
	}
	clients := make(map[string]*HandlerClient)
	for _, in := range ins {
		key,has:=in.Metadata["tag"]
		if has {
			if old, ok := clients[key]; ok {
				clients[key] = old
				continue
			}
			u, err := url.Parse(in.Addrs[0])
			if err == nil && u.Scheme == "grpc" {
				c, err := newHandler(u.Host)
				if err != nil {
					log.Errorf("handler NewComet(%+v) error(%v)", in, err)
					return err
				}
				clients[key] = c
				log.Info("handler AddHandler grpc:%+v", in)
			}
		}
	}
	//for key, old := range _clients{
	//	if _, ok := clients[key]; !ok {
	//		old.
	//		log.Infof("handler DelHandler:%s", key)
	//	}
	//}
	_clients = clients
	return nil
}
