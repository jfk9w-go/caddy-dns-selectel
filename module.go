package selectel

import (
	"errors"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	selectel "github.com/jfk9w-go/libdns-selectel"
)

type Provider struct {
	credentials selectel.Credentials
	*selectel.Provider
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.selectel",
		New: func() caddy.Module {
			return new(Provider)
		},
	}
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider = selectel.NewProvider(selectel.NewClient(p.credentials))
	return nil
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
					p.credentials.Username = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "password":
				if dispenser.NextArg() {
					p.credentials.Password = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "account_id":
				if dispenser.NextArg() {
					p.credentials.AccountID = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}

			case "project_name":
				if dispenser.NextArg() {
					p.credentials.ProjectName = dispenser.Val()
				}

				if dispenser.NextArg() {
					return dispenser.ArgErr()
				}
			}
		}
	}

	if p.credentials.Username == "" {
		return errors.New("no username provided")
	}

	if p.credentials.Password == "" {
		return errors.New("no password provided")
	}

	if p.credentials.AccountID == "" {
		return errors.New("no account id provided")
	}

	if p.credentials.ProjectName == "" {
		return errors.New("no project name provided")
	}

	return nil
}

// type guards

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
