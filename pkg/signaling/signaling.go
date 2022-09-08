// Package signaling implements various signaling backends to exchange encrypted messages between peers
package signaling

import (
	"github.com/stv0g/cunicu/pkg/crypto"

	signalingproto "github.com/stv0g/cunicu/pkg/proto/signaling"
)

type Message = signalingproto.Message
type Envelope = signalingproto.Envelope

type MessageHandler interface {
	OnSignalingMessage(*crypto.PublicKeyPair, *Message)
}

type EnvelopeHandler interface {
	OnSignalingEnvelope(*Envelope)
}
