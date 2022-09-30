package config

import (
	"strings"

	"golang.org/x/exp/slices"
)

type ChangedHandler interface {
	OnConfigChanged(key string, old, new any)
}

func (c *Config) OnInterfaceChanged(name, key string, h ChangedHandler) {

}

func (c *Config) InvokeHandlers(key string, change Change) {
	c.Meta.InvokeHandlers(key, change)

	if keyParts := strings.Split(key, "."); len(keyParts) > 0 && keyParts[0] == "interfaces" {
		pattern := keyParts[1]

		for name, meta := range c.onInterfaceChanged {
			pats := c.InterfaceOrderByName(name)

			if slices.Contains(pats, pattern) {
				key := strings.Join(keyParts[2:], ".")
				meta.InvokeHandlers(key, change)
			}
		}
	}
}