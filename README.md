# hc-install

**NOTE (!)** This is a fork of HashiCorp's [hc-install](github.com/hashicorp/hc-install), modified for specific use in CI.

This repository will become defunct when upstream supports using `hc-install` as a command.

# Getting Started

The `hc-install` command can be installed by running

```bash
go install gophers.dev/cmds/hc-install/cmd/hc-install@latest
```

# Usage

```bash
hc-install [product] [version]
```

Currently only supports installing `consul` and `vault`.

# License

The upstream `github.com/hashicorp/hc-install` repository is open source under the [MPL](LICENSE)
