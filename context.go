package lingo

import base "github.com/lingo-dev/lingo-base"

type Context struct {
	baseCtx *base.Context
}

func NewContext(baseCtx *base.Context) *Context {
	return &Context{baseCtx: baseCtx}
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// Env read env
func (ctx *Context) Env(key string) string {
	return ctx.baseCtx.Envs[key]
}

// Secret read secret
func (ctx *Context) Secret(key string) string {
	return ctx.baseCtx.Secrets[key]
}
