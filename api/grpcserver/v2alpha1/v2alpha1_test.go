package v2alpha1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/spacemeshos/go-spacemesh/api/grpcserver"
)

func launchServer(tb testing.TB, services ...grpcserver.ServiceAPI) (grpcserver.Config, func()) {
	cfg := grpcserver.DefaultTestConfig()
	grpcService, err := grpcserver.NewWithServices(cfg.PublicListener, zaptest.NewLogger(tb).Named("grpc"), cfg, services)
	require.NoError(tb, err)

	// start gRPC server
	require.NoError(tb, grpcService.Start())

	// update config with bound addresses
	cfg.PublicListener = grpcService.BoundAddress

	return cfg, func() { assert.NoError(tb, grpcService.Close()) }
}

func dialGrpc(ctx context.Context, tb testing.TB, cfg grpcserver.Config) *grpc.ClientConn {
	tb.Helper()
	conn, err := grpc.DialContext(ctx,
		cfg.PublicListener,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	require.NoError(tb, err)
	tb.Cleanup(func() { require.NoError(tb, conn.Close()) })
	return conn
}
