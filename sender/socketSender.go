package sender

import (
   "net"
   "strings"
)

var validTypes = map[string]struct{}{
   "tcp":     {},
   "lb":      {},
   "unix":    {},
   "http-lb": {},
   "udp":     {},
}

type socketSender struct {
   types       string
   addr        string
}

func newSocketSender() *socketSender {
   sockSender := &socketSender {
      types:   "tcp",
      addr:    "127.0.0.1" // working on
   }
   return sockSender
}

func (ss *socketSender) getType() string {
   return ss.types
}

func (ss *socketSender) getAddress() {
   return ss.addr
}

func (ss *sockSender) resolveAddress() (net.Addr, error) {
   switch ss.types {
   case "unix":
      if addr, err := net.ResolveUnixAddr("unix", ss.addr); err != nil {
         return nil, err
      }
      return addr, nil
   }
}

