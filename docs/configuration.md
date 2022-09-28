---
sidebar_position: 7
---

# Configuration

This page describes the ways of configuring the cunīcu daemon (`cunicu daemon`).

## Command Line Flags

The `cunicu daemon` can almost fully be configured by passing command line arguments.
A full overview is available in its [manpage](./usage/md/cunicu_daemon.md).

## Configuration File

Alternatively a configuration file can be used for a persistent configuration:

```yaml title="cunicu.yaml"
# An interval at which cunīcu will periodically check for added,
# removed or modified WireGuard interfaces.
watch_interval: 1s

## Signaling backends
#
# These backends are used for exchanging control-plane messages
# between the peers.
# E.g. ICE candidates, Peer information
backends:
- grpc://signal.cunicu.li
# - grpc://localhost:8080?insecure=true&skip_verify=true
# - k8s:///path/to/your/kubeconfig.yaml?namespace=default


# RPC control socket settings
rpc:
  # Path to a Unix socket for management
  # and monitoring of the cunicu daemon.
  socket: /var/run/cunicu.sock

  # Start of cunīcu daemon will block until
  # its unblocked via the control socket.
  # Mostly useful for testing automation
  wait: false


## Hook callbacks
#
# Hook callback can be used to invoke subprocesses
# or web-hooks on certain events within cunīcu.
hooks:

# An 'exec' hook spawn a subprocess for each event.
- type: exec
  command: ../../scripts/hook.sh

  # Prepend additional arguments
  args: []

  # Pass JSON object via Stdin to command
  stdin: true

  # Set environment variables for invocation
  env:
    COLOR: "1"

# A 'web' hook performs HTTP requests for each event.
- type: web

  # URL of the webhook endpoint
  url: https://my-webhook-endpoint.com/api/v1/webhook
  
  # HTTP method of the request
  method: POST

  # Additional HTTP headers which are used for the requests
  headers:
    User-Agent: ahoi
    Authorization: Bearer XXXXXX


#### Interface settings start here
# The following settings can be overwritten for each interface
# using the 'interfaces' settings (see below).
# The following settings will be used as default.

## WireGuard interface settings
#
# These settings configure WireGuard specific settings
# of the interface.
wireguard:
  # A base64 private key generated by wg genkey.
  # Will be automatically generated if not provided.
  private_key: KLoqDLKgoqaUkwctTd+Ov3pfImOfadkkvTdPlXsuLWM=

  # Create WireGuard interfaces using bundled wireguard-go
  # user space implementation. This will be the default
  # if there is no WireGuard kernel module present.
  userspace: false

  # A range constraint for an automatically assigned
  # selected listen port.
  # If the interface has no listen port specified, cunīcu
  # will use the first available port from this range.
  listen_port_range:
    min: 52820
    max: 65535

  # A 16-bit port for listening. Optional;
  # If not specified, first available port from listen_port_range
  # will be used.
  listen_port: 51825

  # A 32-bit fwmark for outgoing packets which can be used
  # for Netfilter or TC classification.
  # If set to 0 or "off", this option is disabled.
  # May be specified in hexadecimal by prepending "0x". Optional.
  fwmark: 0x1000

  # A list of peers.
  peers:
    
  - # A base64 public key calculated by wg pubkey from a private key,
    # and usually transmitted out of band
    # to the author of the configuration file.
    public_key: FlKHqqQQx+bTAq7+YhwEECwWRg2Ih7NQ48F/SeOYRH8=
    
    # A base64 preshared key generated by wg genpsk.
    # Optional, and may be omitted.
    # This option adds an additional layer of symmetric-key
    # cryptography to be mixed into the already existing
    # public-key cryptography, for post-quantum resistance.
    preshared_key: zu86NBVsWOU3cx4UKOQ6MgNj3gv8GXsV9ATzSemdqlI=

    # A preshared passphrase which is used to derive a preshared key.
    # cunīcu is using Argon2id as the key derivation function.
    preshared_key_passphrase:

    # An endpoint IP or hostname, followed by a colon,
    # and then a port number. This endpoint will be updated
    # automatically to the most recent source IP address and
    # port of correctly authenticated packets from the peer.
    endpoint: vpn.example.com:51820

    # A time duration, between 1 and 65535s inclusive, of how
    # often to send an authenticated empty packet to the peer
    # for the purpose of keeping a stateful firewall or NAT mapping
    # valid persistently. For example, if the interface very rarely
    # sends traffic, but it might at anytime receive traffic from a
    # peer, and it is behind NAT, the interface might benefit from
    # having a persistent keepalive interval of 25 seconds.
    # If set to 0 or "off", this option is disabled.
    # By default or when unspecified, this option is off.
    # Most users will not need this. Optional.
    persistent_keepalive: 120s

    # A comma-separated list of IP (v4 or v6) addresses with
    # CIDR masks from which incoming traffic for this peer is
    # allowed and to which outgoing  traffic for this peer is directed.
    # The catch-all 0.0.0.0/0 may be specified for matching
    # all IPv4 addresses, and ::/0 may be specified for matching
    # all IPv6 addresses. May be specified multiple times.
    allowed_ips:
    - 192.168.5.0/24

## Auto configuration
#
autocfg:
  # The Maximum Transfer Unit of the WireGuard interface.
  mtu: 1420

  # IPv4 / IPv6 addresses for the WireGuard interface.
  addresses:
  - 10.10.0.1/24

  # Assign link-local addresses to the WireGuard interface.
  link_local: true

## Config file synchronization
#
# Synchronize local WireGuard interface configuration with wg(8) config-files.
cfgsync:
  enabled: false
  
  # Directory where Wireguard configuration files are located.
  # We expect the same format as used by wg(8) and wg-quick(8).
  # Filenames must match the interface name with a '.conf' suffix.
  path: /etc/wireguard

  # Watch the configuration files via inotify(7) for changes and apply them accordingly.
  watch: false


## Route Synchronization
#
# Synchronize the kernel routing table with WireGuard's AllowedIPs setting
# 
# It checks for routes in the kernel routing table which have a peers link-local address
# as next-hop and adds those routes to the AllowedIPs setting of the respective peer.
#
# In reverse, also networks listed in a peers AllowedIPs setting will be installed as a
# kernel route with the peers link-local address as the routes next-hop. 
rtsync:
  enabled: true

  table: 254 # See /etc/iproute2/rt_tables for table ids

  # Keep watching the for changes in the kernel routing table via netlink multicast group.
  watch: true


## /etc/hosts synchronization
#
# Synchronizes the local /etc/hosts file with host names and link-local IP addresses of connected peers. 
hsync:
  enabled: true

  # The domain name which is appended to each of the peer host names
  domain: wg-local


## Peer discovery
#
# Peer discovery finds new peers within the same community and adds them to the respective interface
pdisc:
  enabled: true

  # The hostname which gets advertised to remote peers
  hostname: my-node

  # Networks which are reachable via this peer and get advertised to remote peers
  # These will be part of this interfaces AllowedIPs at the remote peers.
  networks:
  - 192.168.1.0/24
  - 10.2.0.0/24

  # A list of WireGuard public keys which are accepted peers
  # If not configured, all peers will be accepted.
  whitelist:
  - coNsGPwVPdpahc8U+dbbWGzTAdCd6+1BvPIYg10wDCI=
  - AOZzBaNsoV7P8vo0D5UmuIJUQ7AjMbHbGt2EA8eAuEc=

  # A passphrase shared among all peers of the same community
  community: "some-common-password"


## Endpoint discovery
#
# Endpoint discovery uses Interactive Connectivity Establishment (ICE) as used by WebRTC to
# gather a list of candidate endpoints and performs connectivity checks to find a suitable
# endpoint address which can be used by WireGuard
epdisc:
  enabled: true

  # Interactive Connectivity Establishment (ICE) parameters
  ice:
    # A list of STUN and TURN servers used by ICE.
    urls:
    # Community provided STUN/TURN servers
    - grpc://relay.cunicu.li

    # Public STUN servers
    - stun:stun3.l.google.com:19302
    - stun:relay.webwormhole.io
    - stun:stun.sipgate.net
    - stun:stun.ekiga.net
    - stun:stun.services.mozilla.com

    # Caution: OpenRelay servers are located in Ontario, Canada.
    # Beware of the latency!
    # See also: https://www.metered.ca/tools/openrelay/
    # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:80
    # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:443
    # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:443?transport=tcp

    # Credentials for STUN/TURN servers configured above.
    username: ""
    password: ""

    # Allow connections to STUNS/TURNS servers for which we can not validate TLS certificates.
    insecure_skip_verify: false

    # Limit available network and candidate types.
    # network_types: [udp4, udp6, tcp4, tcp6]
    # candidate_types: [host, srflx, prflx, relay]

    # A glob(7) pattern to match interfaces against which are used to gather ICE candidates (e.g. \"eth[0-9]\").
    interface_filter: "*"

    # Lite agents do not perform connectivity check and only provide host candidates.
    lite: false

    # Enable local Multicast DNS discovery.
    mdns: false

    # Sets the max amount of binding requests the agent will send over a candidate pair for validation or nomination.
    # If after the the configured number, the candidate is yet to answer a binding request or a nomination we set the pair as failed.
    max_binding_requests: 7

    # SetNAT1To1IPs sets a list of external IP addresses of 1:1 (D)NAT and a candidate type for which the external IP address is used.
    # This is useful when you are host a server using Pion on an AWS EC2 instance which has a private address, behind a 1:1 DNAT with a public IP (e.g. Elastic IP).
    # In this case, you can give the public IP address so that Pion will use the public IP address in its candidate instead of the private IP address.
    # nat_1to1_ips:
    # - 10.10.2.3

    # Limit the port range used by ICE
    port_range:
        # Minimum port for allocation policy for ICE sockets (range: 0-65535)
        min: 49152

        # Maximum port for allocation policy for ICE sockets (range: 0-65535)
        max: 65535

    # Interval at which the agent performs candidate checks in the connecting phase
    check_interval: 200ms
    
    # Time until an Agent transitions disconnected.
    # If the duration is 0, the ICE Agent will never go to disconnected
    disconnected_timeout: 5s

    # Time until an Agent transitions to failed after disconnected
    # If the duration is 0, we will never go to failed.
    failed_timeout: 5s

    # Time to wait before ICE restart
    restart_timeout: 5s

    # Interval between STUN keepalives (should be less then connection timeout above).
    # Af the interval is 0, we never send keepalive packets
    keepalive_interval: 2s


## Interface specific settings / overwrites.
#
# Most of the top-level settings of this configuration file can be overwritten
# with settings specific to a single or a group of interfaces.
# This includes the following settings (see below):
# - wireguard
# - cfgsync
# - rtsync
# - hsync
# - pdisc
# - epdisc
# 
# The keys of this mapping are glob(7) patterns which are matched against the
# interface names.
# Settings are overlayed in the order in which the keys are provided in the
# interface map.
#
# Keys which are not a glob(8) pattern, will be created as new interfaces if
# they do not exist already in the system.
interfaces:
  # A simple interface specific setting
  # cunicu will set the private key of interface 'wg0' to the provided value.
  wg0:
    epdisc:
      enabled: false

  # No settings are overwritten. But since this is not a glob pattern,
  # A new interface named 'wg1' will be created if it does not exist yet.
  # The same applies to the previous interface 'wg0'
  wg1: {}

  # Create a new interface using the wireguard-go user-space implementation.
  wg2:
    wireguard:
      userspace: true

  # This pattern configuration will be applied to all interfaces which match the pattern.
  # This rule will not create any new interfaces.
  wg-work-*:
    pdisc:
      community: "mysecret-pass" 
    
    epdisc:
      ice:
        urls:
        - turn:mysecret.turn-server.com

  # Multiple patterns are supported and evaluated in the order they a defined in the configuration file.
  # 
  wg-work-external-*:
    epdisc:
      ice:
        network_types: [ udp6 ]
```

## Environment Variables

All the settings from the configuration file can also be passed via environment variables by following the following rules:

-   Convert the setting name to uppercase
-   Prefixing the setting name with `CUNICU_`
-   Nested settings are separated by underscores

**Example:** The setting `epdisc.ice.max_binding_requests` can be set by the environment variable `CUNICU_ENDPOINT_DISC_ICE_MAX_BINDING_REQUESTS`

:::note
Setting lists such as `epdisc.ice.urls` or `backends` can currently not be set via environment variables.
:::

## At Runtime

cunīcu's configuration can also be updated at runtime, elevating the need to restart the daemon to avoid interruption of connectivity.

Please have a look at the [`cunicu config`](./usage/md/cunicu_config.md) commands.

## DNS Auto-configuration

cunīcu als supports retrieving parts of the configuration via DNS lookups.

When `cunicu daemon` is started with a `--domain example.com` parameter it will look for the following DNS records to obtain its configuration.

STUN and TURN servers used for ICE are retrieved by SVR lookups and other cunīcu settings are retrieved via TXT lookups: 

```text
_stun._udp.example.com.  3600 IN SRV 10 0 3478 stun.example.com.
_stuns._tcp.example.com. 3600 IN SRV 10 0 3478 stun.example.com.
_turn._udp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.
_turn._tcp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.
_turns._tcp.example.com. 3600 IN SRV 10 0 5349 turn.example.com.

example.com.             3600 IN TXT "cunicu-backend=p2p"
example.com.             3600 IN TXT "cunicu-peer-disc-community=my-community-password"
example.com.             3600 IN TXT "cunicu-endpoint-disc-ice-username=user1"
example.com.             3600 IN TXT "cunicu-endpoint-disc-ice-password=pass1"
example.com.             3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"
```

:::note
The `cunicu-backend` and `cunicu-config` TXT records can be provided multiple times. Others not.
:::

## Remote Configuration File

When `cunicu daemon` can be started with `--config` options pointing to HTTPS URIs.
cunīcu will download all configuration files in the order they are specified on the command line and merge them subsequently.

This feature can be combined with the DNS auto-configuration method by providing a TXT record pointing to the configuration file:

```text
example.com.             3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"
```

:::note
Remote configuration files must be fetched via HTTPS if they are not hosted locally and required a trusted server certificate.
:::