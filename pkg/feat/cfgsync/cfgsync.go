// Package cfgsync synchronizes existing WireGuard configuration files with the kernel/userspace WireGuard device.
package cfgsync

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/stv0g/cunicu/pkg/core"
	"github.com/stv0g/cunicu/pkg/device"
	"github.com/stv0g/cunicu/pkg/watcher"
	"github.com/stv0g/cunicu/pkg/wg"
	"go.uber.org/zap"
	"golang.zx2c4.com/wireguard/wgctrl"
)

// ConfigSync synchronizes the WireGuard device configuration with an on-disk configuration file.
type ConfigSync struct {
	watcher *watcher.Watcher
	client  *wgctrl.Client

	// Settings
	cfgPath string
	user    bool
	filter  func(string) bool

	logger *zap.Logger
}

// New creates a new Syncer
func New(w *watcher.Watcher, client *wgctrl.Client, cfgPath string, watch bool, user bool, filter func(string) bool) *ConfigSync {
	cs := &ConfigSync{
		watcher: w,
		client:  client,
		cfgPath: cfgPath,
		user:    user,
		filter:  filter,
		logger:  zap.L().Named("cfgsync"),
	}

	w.OnInterface(cs)

	if watch {
		go cs.watch()
	}

	return cs
}

func (cs *ConfigSync) Start() error {
	cs.logger.Info("Started configuration file synchronization")

	return nil
}

func (cs *ConfigSync) Close() error {
	return nil
}

// OnInterfaceAdded is a handler which is called whenever an interface has been added
func (cs *ConfigSync) OnInterfaceAdded(i *core.Interface) {
	cfg := filepath.Join(cs.cfgPath, fmt.Sprintf("%s.conf", i.Name()))
	if err := i.SyncConfig(cfg); err != nil && !errors.Is(err, os.ErrNotExist) {
		cs.logger.Fatal("Failed to sync interface configuration",
			zap.Error(err),
			zap.String("intf", i.Name()),
			zap.String("config_file", cfg))
	}
}

func (cs *ConfigSync) OnInterfaceRemoved(i *core.Interface) {}

func (cs *ConfigSync) OnInterfaceModified(i *core.Interface, old *wg.Device, m core.InterfaceModifier) {
}

func (cs *ConfigSync) watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		cs.logger.Fatal("failed to create fsnotify watcher", zap.Error(err))
	}

	if _, err := os.Stat(cs.cfgPath); !os.IsNotExist(err) {
		if err := watcher.Add(cs.cfgPath); err != nil {
			cs.logger.Fatal("Failed to watch WireGuard configuration directory",
				zap.Error(err),
				zap.String("path", cs.cfgPath))
		}
	}

	for {
		select {

		// Fsnotify events
		case event := <-watcher.Events:
			cs.logger.Debug("Received fsnotify event", zap.Any("event", event))

			cs.handleFsnotifyEvent(event)

		// Fsnotify errors
		case err := <-watcher.Errors:
			cs.logger.Error("Error while watching for WireGuard configuration files", zap.Error(err))
		}
	}
}

func (cs *ConfigSync) handleFsnotifyEvent(event fsnotify.Event) {
	cfg := event.Name
	filename := path.Base(cfg)
	extension := path.Ext(filename)
	name := strings.TrimSuffix(filename, extension)

	if extension != ".conf" {
		cs.logger.Warn("Ignoring non-configuration file", zap.String("config_file", cfg))
		return
	}

	if !cs.filter(name) {
		cs.logger.Warn("Ignoring configuration file for interface which does not match the interface filter", zap.String("config_file", cfg))
		return
	}

	i := cs.watcher.InterfaceByName(name)

	if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
		if i == nil {
			if i, err := device.NewDevice(name, cs.user); err != nil {
				cs.logger.Error("Failed to create new device",
					zap.Error(err),
					zap.String("intf", i.Name()),
					zap.String("config_file", cfg))
			}
		}

		if err := i.SyncConfig(cfg); err != nil {
			cs.logger.Error("Failed to sync interface configuration",
				zap.Error(err),
				zap.String("intf", i.Name()),
				zap.String("config_file", cfg))
		}
	} else if event.Op&(fsnotify.Remove) != 0 {
		if i == nil {
			cs.logger.Warn("Ignoring unknown interface", zap.String("intf", name))
			return
		}

		// TODO: Do we really want to delete devices if their config file vanish?
		// Maybe make this configurable?
		if err := i.KernelDevice.Close(); err != nil {
			cs.logger.Error("Failed to close interface", zap.Error(err))
		}
	} else if event.Op&(fsnotify.Rename) != 0 {
		// TODO: This is not supported yet
		cs.logger.Warn("We do not support tracking renamed WireGuard configuration files yet")
	}
}

func (cs *ConfigSync) Sync() error {
	des, err := os.ReadDir(cs.cfgPath)
	if err != nil {
		return fmt.Errorf("failed to list config files in '%s': %w", cs.cfgPath, err)
	}

	for _, de := range des {
		cs.handleFsnotifyEvent(fsnotify.Event{
			Name: filepath.Join(cs.cfgPath, de.Name()),
			Op:   fsnotify.Write,
		})
	}

	return nil
}
