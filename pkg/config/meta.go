// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"reflect"
	"slices"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pion/ice/v2"

	"cunicu.li/cunicu/pkg/types"
)

const delim = "."

type Meta struct {
	Fields map[string]*Meta
	Parent *Meta
	Type   reflect.Type

	onChanged []ChangedHandler
}

func Metadata() *Meta {
	settingsType := reflect.TypeOf(Settings{})
	return metadata(settingsType)
}

func metadata(typ reflect.Type) *Meta {
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	m := &Meta{
		Type: typ,
	}

	if typ.Kind() == reflect.Struct {
		m.Fields = map[string]*Meta{}

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if tag, ok := field.Tag.Lookup("koanf"); ok {
				name := strings.Split(tag, ",")[0]
				n := metadata(field.Type)
				if name != "" {
					n.Parent = m
					m.Fields[name] = n
				} else {
					for k, v := range n.Fields {
						m.Fields[k] = v
					}
				}
			}
		}
	}

	return m
}

func (m *Meta) Keys() []string {
	keys := []string{}

	for k, v := range m.Fields {
		if v.Fields == nil {
			keys = append(keys, k)
		} else {
			for _, p := range v.Keys() {
				keys = append(keys, k+delim+p)
			}
		}
	}

	return keys
}

func (m *Meta) Lookup(key string) *Meta {
	parts := strings.Split(key, delim)
	return m.lookup(parts)
}

func (m *Meta) lookup(key []string) *Meta {
	if len(key) == 0 {
		return m
	} else if m.Fields != nil {
		if n, ok := m.Fields[key[0]]; ok {
			return n.lookup(key[1:])
		}
	}

	return nil
}

func (m *Meta) AddChangedHandler(key string, h ChangedHandler) {
	if n := m.Lookup(key); n != nil && !slices.Contains(n.onChanged, h) {
		m.onChanged = append(m.onChanged, h)
	}
}

func (m *Meta) InvokeChangedHandlers(key string, change types.Change) error {
	for n := m.Lookup(key); n != nil; n = n.Parent {
		for _, h := range n.onChanged {
			if err := h.OnConfigChanged(key, change.Old, change.New); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Meta) Parse(str string) (any, error) {
	out := reflect.New(m.Type)

	dec, err := mapstructure.NewDecoder(DecoderConfig(out.Interface()))
	if err != nil {
		return nil, err
	}

	if err := dec.Decode(str); err != nil {
		return nil, err
	}

	return out.Elem().Interface(), nil
}

func (m *Meta) CompletionOptions() []string {
	var options []string

	switch {
	case m.Type.Kind() == reflect.Bool:
		options = append(options, "true", "false")
	case m.Type == reflect.TypeOf([]ice.CandidateType{}):
		options = append(options, "host", "srflx", "prflx", "relay")
	case m.Type == reflect.TypeOf([]ice.NetworkType{}):
		options = append(options, "udp4", "udp6", "tcp4", "tcp6")
	}

	return options
}
