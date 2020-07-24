package main

import (
	"flag"
	"fmt"

	"strings"

	"github.com/felippeduran/pitaya/v2"
	acceptor "github.com/felippeduran/pitaya/v2/acceptor"
	component "github.com/felippeduran/pitaya/v2/component"
	"github.com/felippeduran/pitaya/v2/examples/demo/cluster_protobuf/services"
	groups "github.com/felippeduran/pitaya/v2/groups"
	"github.com/felippeduran/pitaya/v2/serialize/protobuf"
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

	builder := pitaya.NewDefaultBuilder(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, pitaya.NewDefaultBuilderConfig())
	if *isFrontend {
		ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", port))
		builder.AddAcceptor(ws)
	}
	builder.Serializer = protobuf.NewSerializer()
	builder.Groups = groups.NewMemoryGroupService(groups.NewDefaultMemoryGroupConfig())
	app := builder.Build()

	defer app.Shutdown()

	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}

	app.Start()
}
