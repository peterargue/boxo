package mdutils

import (
	dag "github.com/peterargue/boxo/ipld/merkledag"

	bsrv "github.com/peterargue/boxo/blockservice"
	blockstore "github.com/peterargue/boxo/blockstore"
	offline "github.com/peterargue/boxo/exchange/offline"
	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	ipld "github.com/ipfs/go-ipld-format"
)

// Mock returns a new thread-safe, mock DAGService.
func Mock() ipld.DAGService {
	return dag.NewDAGService(Bserv())
}

// Bserv returns a new, thread-safe, mock BlockService.
func Bserv() bsrv.BlockService {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	return bsrv.New(bstore, offline.Exchange(bstore))
}
