# caddy-dns/httpnet

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy).
It can be used to manage DNS records with [http.net](https://www.http.net).

## Caddy module name

```
dns.providers.httpnet
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "httpnet",
        "auth_token": "YOUR_HTTPNET_AUTH_TOKEN"
      }
    }
  }
}
```

Or with the Caddyfile:

```
# globally
{
    acme_dns httpnet <auth_token>
}
```

```
# one site, inline token
tls {
    dns httpnet <auth_token>
}
```

```
# one site, block form with env-var placeholder
tls {
    dns httpnet {
        auth_token {env.HTTPNET_AUTH_TOKEN}
    }
}
```

## Building with xcaddy

```bash
xcaddy build --with github.com/caddy-dns/httpnet
```
