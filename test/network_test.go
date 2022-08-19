//go:build linux

package test_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	g "github.com/stv0g/gont/pkg"
	gopt "github.com/stv0g/gont/pkg/options"
	copt "github.com/stv0g/gont/pkg/options/capture"
	"go.uber.org/zap"

	"riasc.eu/wice/pkg/test"
	"riasc.eu/wice/test/nodes"
)

type Network struct {
	*g.Network

	Name string

	NetworkOptions            []g.Option
	AgentOptions              []g.Option
	WireGuardInterfaceOptions []g.Option

	BasePath string

	SignalingNodes nodes.SignalingList
	RelayNodes     nodes.RelayList
	AgentNodes     nodes.AgentList

	tracer *HandshakeTracer
}

func (n *Network) Start() {
	By("Adding WireGuard peers")

	err := n.AgentNodes.ForEachInterfacePair(func(a, b *nodes.WireGuardInterface) error {
		if a.PeerSelector != nil && a.PeerSelector(a, b) {
			a.AddPeer(b)
		}
		return nil
	})
	Expect(err).To(Succeed(), "Failed to add WireGuard peers: %s", err)

	By("Configuring WireGuard interfaces")

	err = n.AgentNodes.ForEachAgent(func(a *nodes.Agent) error {
		return a.ConfigureWireGuardInterfaces()
	})
	Expect(err).To(Succeed(), "Failed to configure WireGuard interface: %s", err)

	if setup {
		Skip("Aborting test as only network setup has been requested")
	}

	if capture {
		n.StartHandshakeTracer()
	}

	By("Starting relay nodes")

	err = n.RelayNodes.Start(n.BasePath)
	Expect(err).To(Succeed(), "Failed to start relay: %s", err)

	By("Starting signaling nodes")

	err = n.SignalingNodes.Start(binaryPath, n.BasePath)
	Expect(err).To(Succeed(), "Failed to start signaling node: %s", err)

	extraArgs := []any{}

	if len(n.RelayNodes) > 0 {
		// TODO: We currently assume that all relays use the same credentials
		extraArgs = append(extraArgs,
			"--username", n.RelayNodes[0].Username(),
			"--password", n.RelayNodes[0].Password(),
		)
	}

	for _, r := range n.RelayNodes {
		for _, u := range r.URLs() {
			extraArgs = append(extraArgs, "--url", u)
		}
	}

	for _, s := range n.SignalingNodes {
		extraArgs = append(extraArgs, "--backend", s.URL())
	}

	By("Starting agent nodes")

	err = n.AgentNodes.Start(binaryPath, n.BasePath, extraArgs...)
	Expect(err).To(Succeed(), "Failed to start ɯice: %s", err)
}

func (n *Network) Close() {
	By("Stopping agent nodes")

	err := n.AgentNodes.Close()
	Expect(err).To(Succeed(), "Failed to close agent nodes; %s", err)

	By("Stopping signaling nodes")

	err = n.SignalingNodes.Close()
	Expect(err).To(Succeed(), "Failed to close signaling nodes; %s", err)

	By("Stopping relay nodes")

	err = n.RelayNodes.Close()
	Expect(err).To(Succeed(), "Failed to close relay nodes; %s", err)

	By("Stopping network")

	err = n.Network.Close()
	Expect(err).To(Succeed(), "Failed to close network; %s", err)

	n.StopHandshakeTracer()
}

func (n *Network) ConnectivityTests() {
	It("connectivity", func() {
		By("Waiting until all peers are connected")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := n.AgentNodes.WaitConnectionsReady(ctx)
		Expect(err).To(Succeed(), "Failed to wait for peers to connect: %s", err)

		By("Ping between all peers")

		err = n.AgentNodes.PingPeers(ctx)
		Expect(err).To(Succeed(), "Failed to ping peers: %s", err)
	})
}

func (n *Network) Init() {
	*n = Network{}

	n.Name = fmt.Sprintf("wice-%d", rand.Uint32())

	name := GinkgoT().Name()
	n.BasePath = filepath.Join(strings.Split(name, " ")...)
	n.BasePath = filepath.Join("logs", n.BasePath)

	logFilename := filepath.Join(n.BasePath, "test.log")
	pcapFilename := filepath.Join(n.BasePath, "capture.pcapng")

	By("Removing old test case results")

	err := os.RemoveAll(n.BasePath)
	Expect(err).To(Succeed(), "Failed to remove old test case result directory: %s", err)

	By("Creating directory for new test case results")

	err = os.MkdirAll(n.BasePath, 0755)
	Expect(err).To(Succeed(), "Failed to create test case result directory: %s", err)

	// Ginkgo log
	logger = test.SetupLoggingWithFile(logFilename, true)

	// Agent logs
	n.AgentOptions = append(n.AgentOptions,
		gopt.LogToDebug(false),
	)

	n.NetworkOptions = append(n.NetworkOptions,
		gopt.Persistent(persist),
	)

	if capture {
		n.NetworkOptions = append(n.NetworkOptions,
			gopt.CaptureAll(
				copt.Filename(pcapFilename),
			),
		)
	}

	logger.Info("Current test",
		zap.String("name", name),
		zap.String("path", n.BasePath),
		zap.String("executed", time.Now().String()),
	)
}