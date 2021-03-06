package blockrpc

import (
	"fmt"
	"math/rand"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/tclchiam/oxidize-go/blockchain"
	"github.com/tclchiam/oxidize-go/blockchain/engine/mining/proofofwork"
	"github.com/tclchiam/oxidize-go/blockchain/entity"
	"github.com/tclchiam/oxidize-go/blockchain/testdata"
	"github.com/tclchiam/oxidize-go/encoding"
	"github.com/tclchiam/oxidize-go/identity"
	"github.com/tclchiam/oxidize-go/rpc"
)

func TestSyncServer_GetBestHeader(t *testing.T) {
	bc := testdata.NewBlockchainBuilder(t).
		Build().
		ToBlockchain()

	expectedHeader, err := bc.BestHeader()
	if err != nil {
		t.Fatalf("getting best header: %s", err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("starting listener: %s", err)
	}

	server := rpc.NewServer(lis)
	RegisterSyncServer(server, NewSyncServer(bc))
	server.Serve()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	client := NewSyncClient(conn)

	header, err := client.GetBestHeader()
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}

	if !header.IsEqual(expectedHeader) {
		t.Errorf("headers don't match. got - %s, wanted - %s", header, expectedHeader)
	}
}

func TestSyncServer_GetHeaders(t *testing.T) {
	bc := testdata.NewBlockchainBuilder(t).
		Build().
		ToBlockchain()

	saveRandomBlocks(bc, rand.Intn(12))

	startingHeader, err := bc.HeaderByIndex(uint64(rand.Intn(11)))
	if err != nil {
		t.Fatalf("getting header: %s", err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("starting listener: %s", err)
	}

	server := rpc.NewServer(lis)
	RegisterSyncServer(server, NewSyncServer(bc))
	server.Serve()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dialing server: %s", err)
	}
	client := NewSyncClient(conn)

	headers, err := client.GetHeaders(startingHeader.Hash, startingHeader.Index)
	if err != nil {
		t.Fatalf("getting remote headers: %s", err)
	}

	expectedHeaders, err := bc.Headers(startingHeader.Hash, startingHeader.Index)
	if err != nil {
		t.Fatalf("getting local headers: %s", err)
	}

	if !headers.IsEqual(expectedHeaders) {
		t.Errorf("headers don't match. got - %v, wanted - %v", headers, expectedHeaders)
	}
}

func saveRandomBlocks(bc blockchain.Blockchain, num int) error {
	miner := proofofwork.NewDefaultMiner(identity.RandomIdentity().Address())

	for i := 0; i < num; i++ {
		beneficiary := identity.RandomIdentity().Address()
		head, err := bc.BestHeader()
		if err != nil {
			return fmt.Errorf("error reading best header")
		}

		transactions := entity.Transactions{entity.NewRewardTx(beneficiary, encoding.TransactionProtoEncoder())}
		block := miner.MineBlock(head, transactions)
		bc.SaveBlock(block)
	}

	return nil
}
