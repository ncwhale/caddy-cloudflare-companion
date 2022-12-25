# caddy-cloudflare-companion
Automate the CNAME CRUD operations follow the caddy config file.

## Build

Best way to build `caddy` with this plugin is using [`xcaddy`](https://github.com/caddyserver/xcaddy) tool:

0. install `golang` with whatever you like;
1. install `xcaddy`: `go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest`
2. build using `xcaddy`: `xcaddy build --with "github.com/ncwhale/caddy-cloudflare-companion"`

    You can using `--with <other plugin module> `  to add more plugin you like.

## Reference

* [Cloudflare CNAME automation](https://caddy.community/t/cloudflare-cname-automation/13628)
* [libdns/cloudflare](https://github.com/libdns/cloudflare)