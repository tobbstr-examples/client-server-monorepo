package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	v1 "github.com/tobbstr-examples/client-server-monorepo/shared/pb/v1"
)

// TestClient has the following prerequisites:
//
//  1. Start the gRPC server by running `go run server/cmd/server/server.go` in the root folder of this repo.
func TestClient(t *testing.T) {
	// Given
	require := require.New(t)
	ctx := context.Background()
	cli, closeClient, err := NewGrpcClient(ctx, "localhost:1997", "test-client")
	require.NoError(err)
	defer closeClient() // nolint:errcheck

	// When
	resp, err := cli.LastAccess(ctx, &v1.LastAccessRequest{})

	// Then
	require.NoError(err)
	require.NotNil(resp.GetPerson())
	person := *resp.GetPerson()

	// nolint:copylocks
	// Check whether person has a field called Weight, and if not fail the test
	if _, fieldExists := interface{}(&person).(interface {
		GetWeight() uint32
	}); !fieldExists {
		t.Fatalf("person object in response missing Weight field\n")
	}
}
