
package config

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	defaultAddr				= ":28734"
	defaultBalanceMode 	= true

)

var conn = connection.newConnInfo()

type Context struct {
	Address string
	Attributes string
	BalanceMode bool
	ScanPeriod	time.Duration
	ScanMaxIdleTime time.Duration
	TimeTilNodeDead	time.Duration
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.makeDefaults()
	return ctx
}

func (ctx *Context) makeDefaults() {
	ctx.Context.makeDefaults()
	ctx.Address = defaultAddr
	ctx.BalanceMode = defaultBalanceMode 
}

func (ctx *Context) InitNode(() error {
	
	return nil
}

func initEngines(addrs []string, cID string) bool, error {
	elements := len(addrs) - 1
	if elements < 1 {
		return false, err
	}


}

