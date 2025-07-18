package apiclient

import (
	"context"
	"math"
	"time"

	"github.com/argoproj/argo-cd/v3/common"
	"github.com/argoproj/argo-cd/v3/util/env"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpc_util "github.com/argoproj/argo-cd/v3/util/grpc"
	utilio "github.com/argoproj/argo-cd/v3/util/io"
)

// MaxGRPCMessageSize contains max grpc message size
var MaxGRPCMessageSize = env.ParseNumFromEnv(common.EnvGRPCMaxSizeMB, 100, 0, math.MaxInt32) * 1024 * 1024

// Clientset represents config management plugin server api clients
type Clientset interface {
	NewConfigManagementPluginClient() (utilio.Closer, ConfigManagementPluginServiceClient, error)
}

type clientSet struct {
	address string
}

func (c *clientSet) NewConfigManagementPluginClient() (utilio.Closer, ConfigManagementPluginServiceClient, error) {
	conn, err := NewConnection(c.address)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewConfigManagementPluginServiceClient(conn), nil
}

func NewConnection(address string) (*grpc.ClientConn, error) {
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1000 * time.Millisecond)),
	}
	unaryInterceptors := []grpc.UnaryClientInterceptor{grpc_retry.UnaryClientInterceptor(retryOpts...)}
	dialOpts := []grpc.DialOption{
		grpc.WithStreamInterceptor(grpc_util.RetryOnlyForServerStreamInterceptor(retryOpts...)),
		grpc.WithChainUnaryInterceptor(unaryInterceptors...),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxGRPCMessageSize), grpc.MaxCallSendMsgSize(MaxGRPCMessageSize)),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}

	dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc_util.BlockingDial(context.Background(), "unix", address, nil, dialOpts...)
	if err != nil {
		log.Errorf("Unable to connect to config management plugin service with address %s", address)
		return nil, err
	}
	return conn, nil
}

// NewConfigManagementPluginClientSet creates new instance of config management plugin server Clientset
func NewConfigManagementPluginClientSet(address string) Clientset {
	return &clientSet{address: address}
}
