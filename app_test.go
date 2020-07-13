// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package pitaya

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/conn/message"
	"github.com/topfreegames/pitaya/constants"
	e "github.com/topfreegames/pitaya/errors"
	"github.com/topfreegames/pitaya/helpers"
	"github.com/topfreegames/pitaya/logger"
	"github.com/topfreegames/pitaya/logger/logrus"
	"github.com/topfreegames/pitaya/route"
	"github.com/topfreegames/pitaya/router"
	"github.com/topfreegames/pitaya/session/mocks"
	"github.com/topfreegames/pitaya/timer"
)

var (
	tables = []struct {
		isFrontend     bool
		serverType     string
		serverMode     ServerMode
		serverMetadata map[string]string
		cfg            *viper.Viper
	}{
		{true, "sv1", Cluster, map[string]string{"name": "bla"}, viper.New()},
		{false, "sv2", Standalone, map[string]string{}, viper.New()},
	}
)

func TestMain(m *testing.M) {
	exit := m.Run()
	os.Exit(exit)
}

func TestNewApp(t *testing.T) {
	for _, table := range tables {
		t.Run(table.serverType, func(t *testing.T) {
			config := viper.New()
			app := NewDefaultApp(table.isFrontend, table.serverType, table.serverMode, table.serverMetadata, config).(*App)
			assert.Equal(t, table.isFrontend, app.server.Frontend)
			assert.Equal(t, table.serverType, app.server.Type)
			assert.Equal(t, table.serverMode, app.serverMode)
			assert.Equal(t, table.serverMetadata, app.server.Metadata)
		})
	}
}

func TestAddAcceptor(t *testing.T) {
	acc := acceptor.NewTCPAcceptor("0.0.0.0:0")
	for _, table := range tables {
		t.Run(table.serverType, func(t *testing.T) {
			config := viper.New()
			builder := NewBuilder(table.isFrontend, table.serverType, table.serverMode, table.serverMetadata, config)
			builder.AddAcceptor(acc)
			app := builder.Build().(*App)
			if table.isFrontend {
				assert.Equal(t, acc, app.acceptors[0])
			} else {
				assert.Equal(t, 0, len(app.acceptors))
			}
		})
	}
}

func TestSetDebug(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.SetDebug(true)
	assert.Equal(t, true, app.debug)
	app.SetDebug(false)
	assert.Equal(t, false, app.debug)
}

func TestSetLogger(t *testing.T) {
	l := logrus.New()
	SetLogger(l)
	assert.Equal(t, l, logger.Log)
}

func TestGetDieChan(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.Equal(t, app.dieChan, app.GetDieChan())
}

func TestGetConfig(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.Equal(t, app.config, app.GetConfig())
}

func TestGetSever(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.Equal(t, app.server, app.GetServer())
}

func TestGetMetricsReporters(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.Equal(t, app.metricsReporters, app.GetMetricsReporters())
}
func TestGetServerByID(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config)
	s, err := app.GetServerByID("id")
	assert.Nil(t, s)
	assert.EqualError(t, constants.ErrNoServerWithID, err.Error())
}

func TestGetServersByType(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config)
	s, err := app.GetServersByType("id")
	assert.Nil(t, s)
	assert.EqualError(t, constants.ErrNoServersAvailableOfType, err.Error())
}

func TestSetHeartbeatInterval(t *testing.T) {
	inter := 35 * time.Millisecond
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.SetHeartbeatTime(inter)
	assert.Equal(t, inter, app.heartbeat)
}

func TestInitSysRemotes(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.initSysRemotes()
	assert.NotNil(t, app.remoteComp[0])
}

func TestSetDictionary(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)

	dict := map[string]uint16{"someroute": 12}
	err := app.SetDictionary(dict)
	assert.NoError(t, err)
	assert.Equal(t, dict, message.GetDictionary())

	app.running = true
	err = app.SetDictionary(dict)
	assert.EqualError(t, constants.ErrChangeDictionaryWhileRunning, err.Error())
}

func TestAddRoute(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.router = nil
	err := app.AddRoute("somesv", func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		return nil, nil
	})
	assert.EqualError(t, constants.ErrRouterNotInitialized, err.Error())

	app.router = router.New()
	err = app.AddRoute("somesv", func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		return nil, nil
	})
	assert.NoError(t, err)

	app.running = true
	err = app.AddRoute("somesv", func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		return nil, nil
	})
	assert.EqualError(t, constants.ErrChangeRouteWhileRunning, err.Error())
}

func TestShutdown(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	go func() {
		app.Shutdown()
	}()
	<-app.dieChan
}

func TestConfigureDefaultMetricsReporter(t *testing.T) {
	tables := []struct {
		enabled bool
	}{
		{true},
		{false},
	}

	for _, table := range tables {
		t.Run(fmt.Sprintf("%t", table.enabled), func(t *testing.T) {
			config := viper.New()
			config.Set("pitaya.metrics.prometheus.enabled", table.enabled)
			config.Set("pitaya.metrics.statsd.enabled", table.enabled)
			app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
			// if statsd is enabled there are 2 metricsReporters, prometheus and statsd
			assert.Equal(t, table.enabled, len(app.metricsReporters) == 2)
		})
	}
}

func TestDefaultSD(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.NotNil(t, app.serviceDiscovery)

	etcdSD, err := cluster.NewEtcdServiceDiscovery(app.config, app.server, app.dieChan)
	assert.NoError(t, err)
	typeOfetcdSD := reflect.TypeOf(etcdSD)

	assert.Equal(t, typeOfetcdSD, reflect.TypeOf(app.serviceDiscovery))
}

func TestDefaultRPCServer(t *testing.T) {
	ctrl := gomock.NewController(t)

	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.NotNil(t, app.rpcServer)

	sessionPool := mocks.NewMockSessionPool(ctrl)

	natsRPCServer, err := cluster.NewNatsRPCServer(app.config, app.server, nil, app.dieChan, sessionPool)
	assert.NoError(t, err)
	typeOfNatsRPCServer := reflect.TypeOf(natsRPCServer)

	assert.Equal(t, typeOfNatsRPCServer, reflect.TypeOf(app.rpcServer))
}

func TestDefaultRPCClient(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	assert.NotNil(t, app.rpcClient)

	natsRPCClient, err := cluster.NewNatsRPCClient(app.config, app.server, nil, app.dieChan)
	assert.NoError(t, err)
	typeOfNatsRPCClient := reflect.TypeOf(natsRPCClient)

	assert.Equal(t, typeOfNatsRPCClient, reflect.TypeOf(app.rpcClient))
}

func TestStartAndListenStandalone(t *testing.T) {
	config := viper.New()

	acc := acceptor.NewTCPAcceptor("0.0.0.0:0")
	builder := NewBuilder(true, "testtype", Standalone, map[string]string{}, config)
	builder.AddAcceptor(acc)
	app := builder.Build().(*App)

	go func() {
		app.Start()
	}()
	helpers.ShouldEventuallyReturn(t, func() bool {
		return app.running
	}, true)

	assert.NotNil(t, app.handlerService)
	assert.NotNil(t, timer.GlobalTicker)
	// should be listening
	assert.NotEmpty(t, acc.GetAddr())
	helpers.ShouldEventuallyReturn(t, func() error {
		n, err := net.Dial("tcp", acc.GetAddr())
		defer n.Close()
		return err
	}, nil, 10*time.Millisecond, 100*time.Millisecond)
}

func TestStartAndListenCluster(t *testing.T) {
	es, cli := helpers.GetTestEtcd(t)
	defer es.Terminate(t)

	ns := helpers.GetTestNatsServer(t)
	nsAddr := ns.Addr().String()

	cfg := viper.New()
	cfg.Set("pitaya.cluster.rpc.client.nats.connect", fmt.Sprintf("nats://%s", nsAddr))
	cfg.Set("pitaya.cluster.rpc.server.nats.connect", fmt.Sprintf("nats://%s", nsAddr))

	builder := NewBuilder(true, "testtype", Cluster, map[string]string{}, cfg)
	etcdSD, err := cluster.NewEtcdServiceDiscovery(config.NewConfig(cfg), builder.Server, builder.DieChan, cli)
	builder.ServiceDiscovery = etcdSD
	assert.NoError(t, err)
	acc := acceptor.NewTCPAcceptor("0.0.0.0:0")
	builder.AddAcceptor(acc)
	app := builder.Build().(*App)

	go func() {
		app.Start()
	}()
	helpers.ShouldEventuallyReturn(t, func() bool {
		return app.running
	}, true)

	assert.NotNil(t, app.handlerService)
	assert.NotNil(t, timer.GlobalTicker)
	// should be listening
	assert.NotEmpty(t, acc.GetAddr())
	helpers.ShouldEventuallyReturn(t, func() error {
		n, err := net.Dial("tcp", acc.GetAddr())
		defer n.Close()
		return err
	}, nil, 10*time.Millisecond, 100*time.Millisecond)
}

func TestError(t *testing.T) {
	t.Parallel()

	tables := []struct {
		name     string
		err      error
		code     string
		metadata map[string]string
	}{
		{"nil_metadata", errors.New(uuid.New().String()), uuid.New().String(), nil},
		{"empty_metadata", errors.New(uuid.New().String()), uuid.New().String(), map[string]string{}},
		{"non_empty_metadata", errors.New(uuid.New().String()), uuid.New().String(), map[string]string{"key": uuid.New().String()}},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			var err *e.Error
			if table.metadata != nil {
				err = Error(table.err, table.code, table.metadata)
			} else {
				err = Error(table.err, table.code)
			}
			assert.NotNil(t, err)
			assert.Equal(t, table.code, err.Code)
			assert.Equal(t, table.err.Error(), err.Message)
			assert.Equal(t, table.metadata, err.Metadata)
		})
	}
}

func TestGetSessionFromCtx(t *testing.T) {
	ctrl := gomock.NewController(t)
	ss := mocks.NewMockSession(ctrl)

	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, viper.New())
	ctx := context.WithValue(context.Background(), constants.SessionCtxKey, ss)
	s := app.GetSessionFromCtx(ctx)
	assert.Equal(t, ss, s)
}

func TestAddMetricTagsToPropagateCtx(t *testing.T) {
	ctx := AddMetricTagsToPropagateCtx(context.Background(), map[string]string{
		"key": "value",
	})
	val := ctx.Value(constants.PropagateCtxKey)
	assert.Equal(t, map[string]interface{}{
		constants.MetricTagsKey: map[string]string{
			"key": "value",
		},
	}, val)
}

func TestAddToPropagateCtx(t *testing.T) {
	ctx := AddToPropagateCtx(context.Background(), "key", "val")
	val := ctx.Value(constants.PropagateCtxKey)
	assert.Equal(t, map[string]interface{}{"key": "val"}, val)
}

func TestGetFromPropagateCtx(t *testing.T) {
	ctx := AddToPropagateCtx(context.Background(), "key", "val")
	val := GetFromPropagateCtx(ctx, "key")
	assert.Equal(t, "val", val)
}

func TestExtractSpan(t *testing.T) {
	span := opentracing.StartSpan("op", opentracing.ChildOf(nil))
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	spanCtx, err := ExtractSpan(ctx)
	assert.NoError(t, err)
	assert.Equal(t, span.Context(), spanCtx)
}

func TestDescriptor(t *testing.T) {
	bts, err := Descriptor("kick.proto")
	assert.NoError(t, err)
	assert.NotNil(t, bts)

	bts, err = Descriptor("not_exists.proto")
	assert.Nil(t, bts)
	assert.EqualError(t, constants.ErrProtodescriptor, err.Error())
}

func TestDocumentation(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.startupComponents()
	doc, err := app.Documentation(false)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"handlers": map[string]interface{}{},
		"remotes": map[string]interface{}{
			"testtype.sys.bindsession": map[string]interface{}{
				"input": map[string]interface{}{
					"uid":  "string",
					"data": "[]byte",
					"id":   "int64",
				},
				"output": []interface{}{
					map[string]interface{}{
						"error": map[string]interface{}{
							"msg":      "string",
							"code":     "string",
							"metadata": "map[string]string",
						},
						"data": "[]byte",
					},
					"error",
				},
			},
			"testtype.sys.kick": map[string]interface{}{
				"input": map[string]interface{}{
					"userId": "string",
				},
				"output": []interface{}{
					map[string]interface{}{
						"kicked": "bool",
					},
					"error",
				},
			},
			"testtype.sys.pushsession": map[string]interface{}{
				"input": map[string]interface{}{
					"data": "[]byte",
					"id":   "int64",
					"uid":  "string",
				},
				"output": []interface{}{
					map[string]interface{}{
						"error": map[string]interface{}{
							"code":     "string",
							"metadata": "map[string]string",
							"msg":      "string",
						},
						"data": "[]byte",
					},
					"error",
				},
			},
		},
	}, doc)
}

func TestDocumentationTrue(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)
	app.startupComponents()
	doc, err := app.Documentation(true)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"remotes": map[string]interface{}{
			"testtype.sys.bindsession": map[string]interface{}{
				"input": map[string]interface{}{
					"*protos.Session": map[string]interface{}{
						"data": "[]byte",
						"id":   "int64",
						"uid":  "string",
					},
				},
				"output": []interface{}{map[string]interface{}{
					"*protos.Response": map[string]interface{}{
						"data": "[]byte",
						"error": map[string]interface{}{
							"*protos.Error": map[string]interface{}{
								"code":     "string",
								"metadata": "map[string]string",
								"msg":      "string",
							},
						},
					},
				},
					"error",
				},
			},
			"testtype.sys.kick": map[string]interface{}{
				"input": map[string]interface{}{
					"*protos.KickMsg": map[string]interface{}{
						"userId": "string",
					},
				},
				"output": []interface{}{map[string]interface{}{
					"*protos.KickAnswer": map[string]interface{}{
						"kicked": "bool",
					},
				},
					"error",
				},
			},
			"testtype.sys.pushsession": map[string]interface{}{
				"input": map[string]interface{}{
					"*protos.Session": map[string]interface{}{
						"data": "[]byte",
						"id":   "int64",
						"uid":  "string",
					},
				},
				"output": []interface{}{map[string]interface{}{
					"*protos.Response": map[string]interface{}{
						"data": "[]byte",
						"error": map[string]interface{}{
							"*protos.Error": map[string]interface{}{
								"code":     "string",
								"metadata": "map[string]string",
								"msg":      "string",
							},
						},
					},
				},
					"error",
				},
			},
		},
		"handlers": map[string]interface{}{},
	}, doc)
}

func TestAddGRPCInfoToMetadata(t *testing.T) {
	t.Parallel()

	metadata := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	metadata = AddGRPCInfoToMetadata(metadata, "region", "host", "port", "external-host", "external-port")

	assert.Equal(t, map[string]string{
		"key1":                        "value1",
		"key2":                        "value2",
		"key3":                        "value3",
		constants.GRPCHostKey:         "host",
		constants.GRPCPortKey:         "port",
		constants.GRPCExternalHostKey: "external-host",
		constants.GRPCExternalPortKey: "external-port",
		constants.RegionKey:           "region",
	}, metadata)
}

func TestStartWorker(t *testing.T) {
	config := viper.New()
	app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config).(*App)

	app.StartWorker()
	assert.True(t, app.worker.Started())
}

func TestRegisterRPCJob(t *testing.T) {
	t.Run("register_once", func(t *testing.T) {
		config := viper.New()
		app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config)
		app.StartWorker()

		err := app.RegisterRPCJob(nil)
		assert.NoError(t, err)
	})

	t.Run("register_twice", func(t *testing.T) {
		config := viper.New()
		app := NewDefaultApp(true, "testtype", Cluster, map[string]string{}, config)
		app.StartWorker()

		err := app.RegisterRPCJob(nil)
		assert.NoError(t, err)

		err = app.RegisterRPCJob(nil)
		assert.Equal(t, constants.ErrRPCJobAlreadyRegistered, err)
	})
}
