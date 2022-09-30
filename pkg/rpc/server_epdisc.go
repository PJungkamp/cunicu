package rpc

import (
	"context"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stv0g/cunicu/pkg/core"
	"github.com/stv0g/cunicu/pkg/crypto"
	"github.com/stv0g/cunicu/pkg/daemon"
	"github.com/stv0g/cunicu/pkg/daemon/feature/epdisc"
	"github.com/stv0g/cunicu/pkg/proto"

	icex "github.com/stv0g/cunicu/pkg/ice"
	protoepdisc "github.com/stv0g/cunicu/pkg/proto/feature/epdisc"
	rpcproto "github.com/stv0g/cunicu/pkg/proto/rpc"
)

type EndpointDiscoveryServer struct {
	rpcproto.UnimplementedEndpointDiscoverySocketServer

	*Server
}

func NewEndpointDiscoveryServer(s *Server) *EndpointDiscoveryServer {
	eps := &EndpointDiscoveryServer{
		Server: s,
	}

	rpcproto.RegisterEndpointDiscoverySocketServer(s.grpc, eps)

	s.daemon.OnInterface(eps)

	return eps
}

func (s *EndpointDiscoveryServer) InterfaceByCore(ci *core.Interface) *epdisc.Interface {
	if i := s.daemon.InterfaceByCore(ci); i != nil {
		if f, ok := i.Features["epdisc"]; ok {
			if ep, ok := f.(*epdisc.Interface); ok {
				return ep
			}
		}
	}

	return nil
}

func (s *EndpointDiscoveryServer) OnInterfaceAdded(ci *core.Interface) {
	if ep := s.InterfaceByCore(ci); ep != nil {
		ep.OnConnectionStateChange(s)
	}
}

func (s *EndpointDiscoveryServer) RestartPeer(ctx context.Context, params *rpcproto.RestartPeerParams) (*proto.Empty, error) {
	di := s.daemon.InterfaceByName(params.Intf)
	if di == nil {
		return nil, status.Errorf(codes.NotFound, "unknown interface %s", params.Intf)
	}

	i := s.InterfaceByCore(di.Interface)
	if i == nil {
		return nil, status.Errorf(codes.NotFound, "interface %s has endpoint discovery not enabled", params.Intf)
	}

	pk, err := crypto.ParseKeyBytes(params.Peer)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse peer public key: %s", err)
	}

	p := i.PeerByPublicKey(pk)
	if p == nil {
		return nil, status.Errorf(codes.NotFound, "unknown peer %s/%s", params.Intf, pk)
	}

	if err = p.Restart(); err != nil {
		return &proto.Empty{}, status.Errorf(codes.Unknown, "failed to restart peer session: %s", err)
	}

	return &proto.Empty{}, nil
}

func (s *EndpointDiscoveryServer) SendConnectionStates(stream rpcproto.Daemon_StreamEventsServer) {
	s.daemon.ForEachInterface(func(di *daemon.Interface) error {
		i := s.InterfaceByCore(di.Interface)

		for _, p := range i.Peers {
			e := &rpcproto.Event{
				Type:      rpcproto.EventType_PEER_CONNECTION_STATE_CHANGED,
				Interface: p.Interface.Name(),
				Peer:      p.Peer.PublicKey().Bytes(),
				Event: &rpcproto.Event_PeerConnectionStateChange{
					PeerConnectionStateChange: &rpcproto.PeerConnectionStateChangeEvent{
						NewState: protoepdisc.NewConnectionState(p.ConnectionState()),
					},
				},
			}

			if err := stream.Send(e); err == io.EOF {
				continue
			} else if err != nil {
				s.logger.Error("Failed to send", zap.Error(err))
			}
		}

		return nil
	})
}

func (s *EndpointDiscoveryServer) OnConnectionStateChange(p *epdisc.Peer, new, prev icex.ConnectionState) {
	s.events.Send(&rpcproto.Event{
		Type: rpcproto.EventType_PEER_CONNECTION_STATE_CHANGED,

		Interface: p.Interface.Name(),
		Peer:      p.PublicKey().Bytes(),

		Event: &rpcproto.Event_PeerConnectionStateChange{
			PeerConnectionStateChange: &rpcproto.PeerConnectionStateChangeEvent{
				NewState:  protoepdisc.NewConnectionState(new),
				PrevState: protoepdisc.NewConnectionState(prev),
			},
		},
	})
}