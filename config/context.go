
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