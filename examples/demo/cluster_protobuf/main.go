package main

import (
	"flag"
	"fmt"

	"strings"

	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/examples/demo/cluster_protobuf/services"
	"github.com/topfreegames/pitaya/groups"
	"github.com/topfreegames/pitaya/serialize/protobuf"
)

func configureBackend() {
	room := services.NewRoom(app)
	app.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	app.RegisterRemote(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)
}

func configureFrontend(port int) {
	app.Register(&services.Connector{},
		component.WithName("connector"),
		component.WithNameFunc(strings.ToLower),
	)
	app.RegisterRemote(&services.ConnectorRemote{},
		component.WithName("connectorremote"),
		component.WithNameFunc(strings.ToLower),
	)
}

var app pitaya.Pitaya

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")

	flag.Parse()

	builder := pitaya.NewBuilder(*isFrontend, *svType, pitaya.Cluster, map[string]string{})
	if *isFrontend {
		ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", port))
		builder.AddAcceptor(ws)
	}
	builder.Serializer = protobuf.NewSerializer()
	builder.Groups = groups.NewMemoryGroupService(config.NewConfig())
	app := builder.Build()

	defer app.Shutdown()

	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}

	app.Start()
}
