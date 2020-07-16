package builder

import (
	"github.com/google/uuid"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/agent"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/conn/codec"
	"github.com/topfreegames/pitaya/conn/message"
	"github.com/topfreegames/pitaya/groups"
	"github.com/topfreegames/pitaya/logger"
	"github.com/topfreegames/pitaya/metrics"
	"github.com/topfreegames/pitaya/pipeline"
	"github.com/topfreegames/pitaya/router"
	"github.com/topfreegames/pitaya/serialize"
	"github.com/topfreegames/pitaya/service"
	"github.com/topfreegames/pitaya/session"
	"github.com/topfreegames/pitaya/worker"
)

//Builder ...
type Builder interface {
	Build() pitaya.Pitaya

	SetSerializer(serializer serialize.Serializer)
	SetRouter(router *router.Router)
	SetRPCClient(client cluster.RPCClient)
	SetRPCServer(server cluster.RPCServer)
	SetServiceDiscovery(serviceDiscovery cluster.ServiceDiscovery)
	SetGroups(groupService groups.GroupService)
	SetWorker(worker *worker.Worker)
	AddMetricsReporter(reporter metrics.Reporter)
	AddAcceptor(acceptor acceptor.Acceptor)
}

// BuilderImpl ...
type BuilderImpl struct {
	config           pitaya.PitayaConfig
	PacketDecoder    codec.PacketDecoder
	PacketEncoder    codec.PacketEncoder
	MessageEncoder   *message.MessagesEncoder
	Serializer       serialize.Serializer
	Router           *router.Router
	RPCClient        cluster.RPCClient
	RPCServer        cluster.RPCServer
	ServiceDiscovery cluster.ServiceDiscovery
	Groups           groups.GroupService
	Worker           *worker.Worker
	DieChan          chan bool
	SessionPool      session.SessionPool
	HandlerHooks     *pipeline.HandlerHooks
	metricsReporters []metrics.Reporter
	acceptors        []acceptor.Acceptor
	server           *cluster.Server
	serverMode       pitaya.ServerMode
}

// NewBuilder ...
func NewBuilder(isFrontend bool, serverType string, serverMode pitaya.ServerMode, config pitaya.PitayaConfig, serverMetadata map[string]string) Builder {
	return NewBuilderImpl(isFrontend, serverType, serverMode, config, serverMetadata)
}

// NewBuilderImpl ...
func NewBuilderImpl(isFrontend bool, serverType string, serverMode pitaya.ServerMode, config pitaya.PitayaConfig, serverMetadata map[string]string) *BuilderImpl {
	return &BuilderImpl{
		server:           cluster.NewServer(uuid.New().String(), serverType, isFrontend, serverMetadata),
		serverMode:       serverMode,
		config:           config,
		metricsReporters: []metrics.Reporter{},
		DieChan:          make(chan bool),
		HandlerHooks:     pipeline.NewHandlerHooks(),
		SessionPool:      session.NewSessionPool(),
		PacketDecoder:    codec.NewPomeloPacketDecoder(),
		PacketEncoder:    codec.NewPomeloPacketEncoder(),
		MessageEncoder:   message.NewMessagesEncoder(config.MessageCompression),
	}
}

// SetSerializer ...
func (builder *BuilderImpl) SetSerializer(serializer serialize.Serializer) {
	builder.Serializer = serializer
}

// SetRouter ...
func (builder *BuilderImpl) SetRouter(router *router.Router) {
	builder.Router = router
}

// SetRPCClient ...
func (builder *BuilderImpl) SetRPCClient(client cluster.RPCClient) {
	if builder.serverMode == pitaya.Standalone {
		panic("Standalone mode can't have RPC or service discovery instances")
	}
	builder.RPCClient = client
}

// SetRPCServer ...
func (builder *BuilderImpl) SetRPCServer(server cluster.RPCServer) {
	if builder.serverMode == pitaya.Standalone {
		panic("Standalone mode can't have RPC or service discovery instances")
	}
	builder.RPCServer = server
}

// SetServiceDiscovery ...
func (builder *BuilderImpl) SetServiceDiscovery(serviceDiscovery cluster.ServiceDiscovery) {
	if builder.serverMode == pitaya.Standalone {
		panic("Standalone mode can't have RPC or service discovery instances")
	}
	builder.ServiceDiscovery = serviceDiscovery
}

// SetGroups ...
func (builder *BuilderImpl) SetGroups(groupService groups.GroupService) {
	builder.Groups = groupService
}

// SetWorker ...
func (builder *BuilderImpl) SetWorker(worker *worker.Worker) {
	builder.Worker = worker
}

// AddMetricsReporter ...
func (builder *BuilderImpl) AddMetricsReporter(reporter metrics.Reporter) {
	builder.metricsReporters = append(builder.metricsReporters, reporter)
}

// AddAcceptor adds a new acceptor
func (builder *BuilderImpl) AddAcceptor(ac acceptor.Acceptor) {
	if !builder.server.Frontend {
		logger.Log.Error("tried to add an acceptor to a backend server, skipping")
		return
	}
	builder.acceptors = append(builder.acceptors, ac)
}

// Build ...
func (builder *BuilderImpl) Build() pitaya.Pitaya {
	if builder.Serializer == nil {
		panic("You must define a message serializer")
	}

	if builder.Router == nil {
		panic("You must define a router")
	}

	if builder.server.Frontend {
		if len(builder.acceptors) == 0 {
			panic("Frontend servers must have at least one acceptor")
		}
	} else {
		if len(builder.acceptors) > 0 {
			panic("Backend servers can't have acceptors")
		}
	}

	var remoteService *service.RemoteService
	if builder.serverMode == pitaya.Standalone {
		if builder.ServiceDiscovery != nil || builder.RPCClient != nil || builder.RPCServer != nil {
			panic("Standalone mode can't have RPC or service discovery instances")
		}
	} else {
		if !(builder.ServiceDiscovery != nil && builder.RPCClient != nil && builder.RPCServer != nil) {
			panic("Cluster mode must have RPC and service discovery instances")
		}

		builder.Router.SetServiceDiscovery(builder.ServiceDiscovery)

		remoteService = service.NewRemoteService(
			builder.RPCClient,
			builder.RPCServer,
			builder.ServiceDiscovery,
			builder.PacketEncoder,
			builder.Serializer,
			builder.Router,
			builder.MessageEncoder,
			builder.server,
			builder.SessionPool,
			builder.HandlerHooks,
		)

		builder.RPCServer.SetPitayaServer(remoteService)
	}

	agentFactory := agent.NewAgentFactory(builder.DieChan,
		builder.PacketDecoder,
		builder.PacketEncoder,
		builder.Serializer,
		builder.config.HearbeatInterval,
		builder.MessageEncoder,
		builder.config.BufferAgentMessages,
		builder.SessionPool,
		builder.metricsReporters,
	)

	handlerService := service.NewHandlerService(
		builder.PacketDecoder,
		builder.Serializer,
		builder.config.BufferHandlerLocalProcess,
		builder.config.BufferHandlerRemoteProcess,
		builder.server,
		remoteService,
		agentFactory,
		builder.metricsReporters,
		builder.HandlerHooks,
	)

	return pitaya.NewApp(
		builder.serverMode,
		builder.Serializer,
		builder.acceptors,
		builder.DieChan,
		builder.Router,
		builder.server,
		builder.RPCClient,
		builder.RPCServer,
		builder.Worker,
		builder.ServiceDiscovery,
		remoteService,
		handlerService,
		builder.Groups,
		builder.SessionPool,
		builder.metricsReporters,
		builder.config,
	)
}
