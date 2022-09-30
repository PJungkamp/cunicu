package watcher

import (
	"github.com/stv0g/cunicu/pkg/core"
	"github.com/stv0g/cunicu/pkg/wg"
)

// All handler

type allHandler struct {
	core.AllHandler
}

func (h *allHandler) OnInterfaceAdded(i *core.Interface) {
	i.OnModified(h)
	i.OnPeer(h)

	h.AllHandler.OnInterfaceAdded(i)
}

func (h *allHandler) OnInterfaceRemoved(i *core.Interface) {
	h.AllHandler.OnInterfaceRemoved(i)
}

func (h *allHandler) OnInterfaceModified(i *core.Interface, old *wg.Device, m core.InterfaceModifier) {
	h.AllHandler.OnInterfaceModified(i, old, m)
}

func (h *allHandler) OnPeerAdded(p *core.Peer) {
	p.OnModified(h)

	h.AllHandler.OnPeerAdded(p)
}

// Peer handler

type peerHandler struct {
	core.PeerHandler
}

func (h *peerHandler) OnInterfaceAdded(i *core.Interface) {
	i.OnPeer(h)
}

func (h *peerHandler) OnInterfaceRemoved(i *core.Interface) {}

// OnAll adds a new handler to all the events observed by the watcher.
func (w *Watcher) OnAll(h core.AllHandler) {
	w.OnInterface(&allHandler{h})
}

// OnPeer registers an handler for peer-related events
func (w *Watcher) OnPeer(h core.PeerHandler) {
	w.OnInterface(&peerHandler{h})
}

// OnInterface registers an handler for interface-related events
func (w *Watcher) OnInterface(h core.InterfaceHandler) {
	w.onInterface = append(w.onInterface, h)
}
