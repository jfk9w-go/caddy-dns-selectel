package selectel

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	selectel "github.com/jfk9w-go/libdns-selectel"
)

type Provider struct {
	*selectel.Provider
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.selectel",
		New: func() caddy.Module {
			return &Provider{
				Provider: new(selectel.Provider),
			}
		},
	}
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Credentials.Username = caddy.NewReplacer().ReplaceAll(p.Credentials.Username, "")
	p.Credentials.Password = caddy.NewReplacer().ReplaceAll(p.Credentials.Password, "")
	p.Credentials.AccountID = caddy.NewReplacer().ReplaceAll(p.Credentials.AccountID, "")
	p.Credentials.ProjectName = caddy.NewReplacer().ReplaceAll(p.Credentials.ProjectName, "")
	return p.Credentials.Validate()
}

func (p *Provider) UnmarshalCaddyfile(dispenser *caddyfile.Dispenser) error {
	for dispenser.Next() {
		if dispenser.NextArg() {
			return dispenser.ArgErr()
		}

		for nesting := dispenser.Nesting(); dispenser.NextBlock(nesting); {
			switch dispenser.Val() {
			case "username":
				if dispenser.NextArg() {
					p.Credentials.Username = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "password":
				if dispenser.NextArg() {
					p.Credentials.Password = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "account_id":
				if dispenser.NextArg() {
					p.Credentials.AccountID = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "project_name":
				if dispenser.NextArg() {
					p.Credentials.ProjectName = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}
			}
		}
	}

	return p.Credentials.Validate()
}

// type guards

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
