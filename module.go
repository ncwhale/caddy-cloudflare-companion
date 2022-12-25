package cfc

import (
	caddy "github.com/caddyserver/caddy/v2"
)

func init() {
	caddy.RegisterModule(CfCompanion{})
}

type CfCompanion struct{}

func (CfCompanion) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.proxy.cloudflare",
		New: func() caddy.Module { return new(CfCompanion) },
	}
}
