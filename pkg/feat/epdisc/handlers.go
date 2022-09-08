package epdisc

import icex "github.com/stv0g/cunicu/pkg/feat/epdisc/ice"

type OnConnectionStateHandler interface {
	OnConnectionStateChange(p *Peer, new, prev icex.ConnectionState)
}

func (e *EndpointDiscovery) OnConnectionStateChange(h OnConnectionStateHandler) {
	e.onConnectionStateChange = append(e.onConnectionStateChange, h)
}
