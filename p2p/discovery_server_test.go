package p2p

import (
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/tclchiam/oxidize-go/blockchain/testdata"
	"github.com/tclchiam/oxidize-go/rpc"
)

func TestDiscoveryServer_Ping(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("starting listener: %s", err)
	}

	server := rpc.NewServer(lis)
	RegisterDiscoveryServer(server, NewDiscoveryServer(nil))
	server.Serve()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTimeout(500*time.Millisecond), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	client := NewDiscoveryClient(conn)

	err = client.Ping()

	if err != nil {
		t.Errorf("expected no error, was: %s", err)
	}
}

func TestDiscoveryServer_Ping_TargetIsOffline(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:0", grpc.WithTimeout(500*time.Millisecond), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	client := NewDiscoveryClient(conn)

	err = client.Ping()

	expectedErrorMessage := "rpc error: code = Unavailable desc = all SubConns are in TransientFailure"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("unexpected error message. got - %s, wanted: %s", err, expectedErrorMessage)
	}
}

func TestDiscoveryServer_Version(t *testing.T) {
	bc := testdata.NewBlockchainBuilder(t).
		Build().
		ToBlockchain()

	actualHeader, err := bc.BestHeader()
	if err != nil {
		t.Fatalf("getting best header: %s", err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("starting listener: %s", err)
	}

	server := rpc.NewServer(lis)
	RegisterDiscoveryServer(server, NewDiscoveryServer(bc))
	server.Serve()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTimeout(500*time.Millisecond), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	client := NewDiscoveryClient(conn)

	hash, err := client.Version()

	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	if !hash.IsEqual(actualHeader.Hash) {
		t.Errorf("headers don't match. got - %s, wanted - %s", hash, actualHeader.Hash)
	}
}

func TestDiscoveryServer_Version_TargetIsOffline(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:0", grpc.WithTimeout(500*time.Millisecond), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	client := NewDiscoveryClient(conn)

	_, err = client.Version()

	expectedErrorMessage := "rpc error: code = Unavailable desc = all SubConns are in TransientFailure"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("unexpected error message. got - %s, wanted: %s", err, expectedErrorMessage)
	}
}
