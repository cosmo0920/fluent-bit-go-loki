# fluent-bit loki output plugin

This plugin works with fluent-bit's go plugin interface. You can use fluent-bit loki to ship logs into grafana datasource with loki.

The configuration typically looks like:

```graphviz
fluent-bit --> loki --> graphana <-- other grafhana sources
```

# Usage

```bash
$ fluent-bit -e /path/to/built/out_loki.so -c fluent-bit.conf
```

# Prerequisites

* Go 1.11+
* gcc (for cgo)

## Building

```bash
$ make
```

### Configuration Options

| Key           | Description                      | Default                             |
| --------------|----------------------------------|-------------------------------------|
| Url           | Url of loki server API endpoint  | http://localhost:3100/api/prom/push |

Example:

add this section to fluent-bit.conf

```properties
[Output]
    Name loki
    Match *
    Url http://localhost:3100/api/prom/push
```

## Useful links

* [fluent-bit-go](https://github.com/fluent/fluent-bit-go)
