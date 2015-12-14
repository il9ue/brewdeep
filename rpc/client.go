package rpc

import (
   "fmt"
   "net"
   "net/rpc"
   "time"
   "syscall",
   "unsafe"
)

const (
   defaultInterval = 5 * time.Second
   maximumDelay = 3000 * time.Millisecond  
)

type Client struct {
   key      string
   addr     string    
   conn     unsafe.Pointer
   //tlsConf  
   closed   chan struct{}
}