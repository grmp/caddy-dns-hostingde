# grmp/caddy-dns-hostingde

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy).
It can be used to manage DNS records with [hosting.de](https://www.hosting.de).

## Caddy module name

```
dns.providers.hostingde
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "hostingde",
        "auth_token": "YOUR_HOSTINGDE_AUTH_TOKEN"
      }
    }
  }
}
```

Or with the Caddyfile:

```
# globally
{
    acme_dns hostingde <auth_token>
}
```

```
# one site, inline token
tls {
    dns hostingde <auth_token>
}
```

```
# one site, block form with env-var placeholder
tls {
    dns hostingde {
        auth_token {env.HOSTINGDE_AUTH_TOKEN}
    }
}
```

## Building with xcaddy

```bash
xcaddy build --with github.com/grmp/caddy-dns-hostingde
```
