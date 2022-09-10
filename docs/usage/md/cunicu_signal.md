---
title: cunicu signal
sidebar_label: signal
sidebar_class_name: command-name
slug: /usage/man/signal
hide_title: true
keywords:
    - manpage
---

## cunicu signal

Start gRPC signaling server

```
cunicu signal [flags]
```

### Options

```
  -h, --help            help for signal
  -L, --listen string   listen address (default ":8080")
  -S, --secure          listen with TLS
```

### Options inherited from parent commands

```
  -C, --color string       Enable colorization of output (one of: auto, always, never) (default "auto")
  -l, --log-file string    path of a file to write logs to
  -d, --log-level string   log level (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")
  -v, --verbose int        verbosity level
```

### SEE ALSO

* [cunicu](cunicu.md)	 - cunīcu is a user-space daemon managing WireGuard® interfaces to establish peer-to-peer connections in harsh network environments.
