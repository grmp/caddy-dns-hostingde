// Package httpnet implements a Caddy DNS provider module for http.net.
//
// This module wraps github.com/libdns/httpnet and registers it with Caddy's
// module system so it can be used to solve ACME DNS-01 challenges.
//
// Caddyfile syntax:
//
//	# Inline arg (auth token)
//	tls {
//	    dns httpnet <auth_token>
//	}
//
//	# Block form
//	tls {
//	    dns httpnet {
//	        auth_token <auth_token>
//	    }
//	}
//
// Credentials may use Caddy placeholder syntax, e.g. {env.HTTPNET_AUTH_TOKEN}.
package httpnet

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/httpnet"
)

// Provider wraps the libdns http.net provider as a Caddy module.
type Provider struct{ *httpnet.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.httpnet",
		New: func() caddy.Module { return &Provider{new(httpnet.Provider)} },
	}
}

// Provision sets up the module. It resolves any Caddy placeholder expressions
// (e.g. environment variable references) in the credentials.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.AuthToken = repl.ReplaceAll(p.Provider.AuthToken, "")
	if p.Provider.AuthToken == "" {
		return fmt.Errorf("httpnet: auth_token is required")
	}
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens.
//
// Syntax:
//
//	httpnet [<auth_token>] {
//	    auth_token <auth_token>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.AuthToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "auth_token":
				if p.Provider.AuthToken != "" {
					return d.Err("auth_token already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.AuthToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	return nil
}

// Interface guards — compile-time checks.
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
