package operator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	priceFeedAdapter "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/PriceFeedAdapter"
	"github.com/hashicorp/raft"
)

/*
 2 Modules:

 1) Service - HTTP (or rpc) server that allows other nodes to join raft network
	- Security need to enforce that requestor signs a message to verify operator (possibly a /verify endpoint to call)

 2) Price FSM - Interface needed for raft consensus
    - Apply, Snapshot and Restore
	- Join handles new nodes (copy store from example)
	- Initialize starts the node and has it try to join the network
		- If first node then bootstrap cluster
	- Consider triggering leaders to call apply requesting task data and task response will have price responses
*/

const (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

type priceUpdateCommand struct {
	FeedName string `json:"feedName"`
}

type PriceFSM struct {
	raft      *raft.Raft
	mu        sync.Mutex
	priceData map[string]int // past price data
	logger    *log.Logger
	// needed to fetch the price of assets on different on-chain oracle networks
	priceFeedAdapter *priceFeedAdapter.ContractPriceFeedAdapter
}

func New(feedAdapter *priceFeedAdapter.ContractPriceFeedAdapter) *PriceFSM {
	return &PriceFSM{
		priceData:        make(map[string]int),
		logger:           log.New(os.Stderr, "[priceData] ", log.LstdFlags),
		priceFeedAdapter: feedAdapter,
	}
}

func (p *PriceFSM) Initialize(enableSingle bool, localId string) error {
	return nil
}

// Join joins a node, identified by nodeID and located at addr, to this store.
// The node must be ready to respond to Raft communications at that address.
func (p *PriceFSM) Join(nodeID, addr string) error {
	p.logger.Printf("received join request for remote node %s at %s", nodeID, addr)

	configFuture := p.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		p.logger.Printf("failed to get raft configuration: %v", err)
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				p.logger.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return nil
			}

			future := p.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := p.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	p.logger.Printf("node %s at %s joined successfully", nodeID, addr)
	return nil
}

type fsm PriceFSM

type fsmSnapshot struct {
	priceData map[string]int
}

// Triggers operator to fetch the requested price feed and sumbit to leader
func (f *fsm) Apply(l *raft.Log) interface{} {
	// 1 - JSON parse log data and serialize to struct
	// 2 - Look up leader and their http url
	// 3 - Fetch price of feed from multiple sources
	// 4 - Create task response
	// 5 - Submit response to leader with BLS signature
	return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	return &fsmSnapshot{priceData: make(map[string]int)}, nil
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	return nil
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := json.Marshal(f.priceData)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (f *fsmSnapshot) Release() {}
