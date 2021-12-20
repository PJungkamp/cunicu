package p2p

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/pnet"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	maddr "github.com/multiformats/go-multiaddr"
)

const (
	defaultMDNSServiceTag = "wice"
)

var defaultConfig = BackendConfig{
	EnableDHTDiscovery:  true,
	EnableMDNSDiscovery: true,
	MDNSServiceTag:      defaultMDNSServiceTag,
	ListenAddresses:     make(multiAddressList, 0),
	BootstrapPeers:      make(peerAddressList, 0),
}

type peerAddressList []peer.AddrInfo
type multiAddressList []maddr.Multiaddr

type BackendConfig struct {
	URI *url.URL

	// Load some options
	ListenAddresses multiAddressList

	// BootstrapPeers is a list of peers to which we initially connect
	BootstrapPeers peerAddressList

	RendezvousString string

	// PrivateNetwork configures libp2p to use the given private network protector.
	PrivateNetwork pnet.PSK

	// DHTDiscovery enables peer discovery and content routing via the Kadmelia DHT.
	EnableDHTDiscovery bool

	// MDNSDiscovery enables peer discovery via local mDNS.
	EnableMDNSDiscovery bool

	// MDNSServiceTag is used in our mDNS advertisements to discover other chat peers.
	MDNSServiceTag string

	// NATPortMap configures libp2p to use the default NATManager. The default NATManager will attempt to open a port in your network's firewall using UPnP.
	EnableNATPortMap bool

	// Relay enables the relay transport.
	EnableRelay bool

	// EnableAutoRelay configures libp2p to enable the AutoRelay subsystem.
	EnableAutoRelay bool

	// AutoRelayAddresses is a list of relays which should be used
	AutoRelayPeers peerAddressList

	// EnableHolePunching enables NAT traversal by enabling NATT'd peers to both initiate and respond to hole punching attempts to create direct/NAT-traversed connections with other peers.
	EnableHolePunching bool
}

func (al *multiAddressList) Set(option string) error {
	as := strings.Split(option, ":")
	for _, a := range as {
		ma, err := maddr.NewMultiaddr(a)
		if err != nil {
			return err
		}

		*al = append(*al, ma)
	}

	return nil
}

func (al *peerAddressList) Set(option string) error {
	as := strings.Split(option, ":")
	for _, a := range as {
		pi, err := peer.AddrInfoFromString(a)
		if err != nil {
			return err
		}

		*al = append(*al, *pi)
	}

	return nil
}

func (c *BackendConfig) Parse(uri *url.URL, options map[string]string) error {
	if rStr, ok := options["rendevouz-string"]; ok {
		c.RendezvousString = rStr
	} else {
		c.RendezvousString = uri.Host
	}

	if laStr, ok := options["listen-addresses"]; ok {
		if err := c.ListenAddresses.Set(laStr); err != nil {
			return fmt.Errorf("failed to parse listen-address option: %w", err)
		}
	}

	if bpStr, ok := options["bootstrap-peers"]; ok {
		if err := c.BootstrapPeers.Set(bpStr); err != nil {
			return fmt.Errorf("failed to parse listen-address option: %w", err)
		}
	}

	// use the default set of bootstrap peers if none are provided
	if len(c.BootstrapPeers) == 0 {
		for _, s := range dht.DefaultBootstrapPeers {
			if pi, err := peer.AddrInfoFromP2pAddr(s); err != nil {
				c.BootstrapPeers = append(c.BootstrapPeers, *pi)
			}
		}
	}

	if len(c.AutoRelayPeers) == 0 {
		for _, s := range autorelay.DefaultRelays {
			if pi, err := peer.AddrInfoFromString(s); err == nil {
				c.AutoRelayPeers = append(c.AutoRelayPeers, *pi)
			}
		}
	}

	return nil
}
