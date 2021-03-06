// Package observer launches a service attached to the sharding node
// that simply observes activity across the sharded Ethereum network.
package observer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/sharding"
)

// Observer holds functionality required to run an observer service
// in a sharded system. Must satisfy the Service interface defined in
// sharding/service.go.
type Observer struct {
	shardp2p     sharding.ShardP2P
	shardChainDb ethdb.Database
	shardID      int
}

// NewObserver creates a new observer instance.
func NewObserver(shardp2p sharding.ShardP2P, shardChainDb ethdb.Database, shardID int) (*Observer, error) {
	return &Observer{shardp2p, shardChainDb, shardID}, nil
}

// Start the main routine for an observer.
func (o *Observer) Start() {
	log.Info(fmt.Sprintf("Starting observer service in shard %d", o.shardID))
}

// Stop the main loop for observing the shard network.
func (o *Observer) Stop() error {
	log.Info(fmt.Sprintf("Starting observer service in shard %d", o.shardID))
	return nil
}
