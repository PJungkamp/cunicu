// Package rtsync synchronizes the kernel routing table with the AllowedIPs of each WireGuard peer
package rtsync

import (
	"net/netip"

	"github.com/stv0g/cunicu/pkg/core"
	"github.com/stv0g/cunicu/pkg/daemon"
	"go.uber.org/zap"
)

func init() {
	daemon.Features["rtsync"] = &daemon.FeaturePlugin{
		New:         New,
		Description: "Route synchronization",
		Order:       30,
	}
}

type Interface struct {
	*daemon.Interface

	gwMap map[netip.Addr]*core.Peer

	stop chan struct{}

	logger *zap.Logger
}

func New(i *daemon.Interface) (daemon.Feature, error) {
	if !i.Settings.RouteSync.Enabled {
		return nil, nil
	}

	rs := &Interface{
		Interface: i,
		gwMap:     map[netip.Addr]*core.Peer{},
		stop:      make(chan struct{}),
		logger:    zap.L().Named("rtsync").With(zap.String("intf", i.Name())),
	}

	i.OnPeer(rs)
	// i.Daemon.Config.OnInterfaceChanged("rtsync", rs)

	if i.Settings.RouteSync.Watch {
		go rs.watchKernel()
	}

	return rs, nil
}

func (rs *Interface) Start() error {
	rs.logger.Info("Started route synchronization")

	return nil
}

func (rs *Interface) Close() error {
	// TODO: Remove Kernel routes added by us

	close(rs.stop)

	return nil
}

func (rs *Interface) OnConfigChanged(key string, old, new any) {
	rs.logger.Warn("Config changed", zap.String("key", key), zap.Any("old", old), zap.Any("new", new))
}