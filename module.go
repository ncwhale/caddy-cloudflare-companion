package cloudflare

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/cloudflare"
)

func init() {
	caddy.RegisterModule(Companion{})
}

type Companion struct{ *cloudflare.Provider }

func (Companion) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.proxy.cloudflare",
		New: func() caddy.Module { return &Companion{new(cloudflare.Provider)} },
	}
}

func (this *Companion) Provision(ctx caddy.Context) error {
	this.Provider.APIToken = caddy.NewReplacer().ReplaceAll(this.Provider.APIToken, "")
	return nil
}

func (p *Companion) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.APIToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

var (
	_ caddyfile.Unmarshaler = (*Companion)(nil)
	_ caddy.Provisioner     = (*Companion)(nil)
)
