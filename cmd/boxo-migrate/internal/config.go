package migrate

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	ImportPaths map[string]string
	Modules     []string
}

var DefaultConfig = Config{
	ImportPaths: map[string]string{
		"github.com/ipfs/go-bitswap":                     "github.com/peterargue/boxo/bitswap",
		"github.com/ipfs/go-ipfs-files":                  "github.com/peterargue/boxo/files",
		"github.com/ipfs/tar-utils":                      "github.com/peterargue/boxo/tar",
		"github.com/ipfs/interface-go-ipfs-core":         "github.com/peterargue/boxo/coreiface",
		"github.com/ipfs/go-unixfs":                      "github.com/peterargue/boxo/ipld/unixfs",
		"github.com/ipfs/go-pinning-service-http-client": "github.com/peterargue/boxo/pinning/remote/client",
		"github.com/ipfs/go-path":                        "github.com/peterargue/boxo/path",
		"github.com/ipfs/go-namesys":                     "github.com/peterargue/boxo/namesys",
		"github.com/ipfs/go-mfs":                         "github.com/peterargue/boxo/mfs",
		"github.com/ipfs/go-ipfs-provider":               "github.com/peterargue/boxo/provider",
		"github.com/ipfs/go-ipfs-pinner":                 "github.com/peterargue/boxo/pinning/pinner",
		"github.com/ipfs/go-ipfs-keystore":               "github.com/peterargue/boxo/keystore",
		"github.com/ipfs/go-filestore":                   "github.com/peterargue/boxo/filestore",
		"github.com/ipfs/go-ipns":                        "github.com/peterargue/boxo/ipns",
		"github.com/ipfs/go-blockservice":                "github.com/peterargue/boxo/blockservice",
		"github.com/ipfs/go-ipfs-chunker":                "github.com/peterargue/boxo/chunker",
		"github.com/ipfs/go-fetcher":                     "github.com/peterargue/boxo/fetcher",
		"github.com/ipfs/go-ipfs-blockstore":             "github.com/peterargue/boxo/blockstore",
		"github.com/ipfs/go-ipfs-posinfo":                "github.com/peterargue/boxo/filestore/posinfo",
		"github.com/ipfs/go-ipfs-util":                   "github.com/peterargue/boxo/util",
		"github.com/ipfs/go-ipfs-ds-help":                "github.com/peterargue/boxo/datastore/dshelp",
		"github.com/ipfs/go-verifcid":                    "github.com/peterargue/boxo/verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline":       "github.com/peterargue/boxo/exchange/offline",
		"github.com/ipfs/go-ipfs-routing":                "github.com/peterargue/boxo/routing",
		"github.com/ipfs/go-ipfs-exchange-interface":     "github.com/peterargue/boxo/exchange",
		"github.com/ipfs/go-merkledag":                   "github.com/peterargue/boxo/ipld/merkledag",
		"github.com/boxo/ipld/car":                       "github.com/ipld/go-car",

		// Pre Boxo rename
		"github.com/ipfs/go-libipfs/gateway":               "github.com/peterargue/boxo/gateway",
		"github.com/ipfs/go-libipfs/bitswap":               "github.com/peterargue/boxo/bitswap",
		"github.com/ipfs/go-libipfs/files":                 "github.com/peterargue/boxo/files",
		"github.com/ipfs/go-libipfs/tar":                   "github.com/peterargue/boxo/tar",
		"github.com/ipfs/go-libipfs/coreiface":             "github.com/peterargue/boxo/coreiface",
		"github.com/ipfs/go-libipfs/unixfs":                "github.com/peterargue/boxo/ipld/unixfs",
		"github.com/ipfs/go-libipfs/pinning/remote/client": "github.com/peterargue/boxo/pinning/remote/client",
		"github.com/ipfs/go-libipfs/path":                  "github.com/peterargue/boxo/path",
		"github.com/ipfs/go-libipfs/namesys":               "github.com/peterargue/boxo/namesys",
		"github.com/ipfs/go-libipfs/mfs":                   "github.com/peterargue/boxo/mfs",
		"github.com/ipfs/go-libipfs/provider":              "github.com/peterargue/boxo/provider",
		"github.com/ipfs/go-libipfs/pinning/pinner":        "github.com/peterargue/boxo/pinning/pinner",
		"github.com/ipfs/go-libipfs/keystore":              "github.com/peterargue/boxo/keystore",
		"github.com/ipfs/go-libipfs/filestore":             "github.com/peterargue/boxo/filestore",
		"github.com/ipfs/go-libipfs/ipns":                  "github.com/peterargue/boxo/ipns",
		"github.com/ipfs/go-libipfs/blockservice":          "github.com/peterargue/boxo/blockservice",
		"github.com/ipfs/go-libipfs/chunker":               "github.com/peterargue/boxo/chunker",
		"github.com/ipfs/go-libipfs/fetcher":               "github.com/peterargue/boxo/fetcher",
		"github.com/ipfs/go-libipfs/blockstore":            "github.com/peterargue/boxo/blockstore",
		"github.com/ipfs/go-libipfs/filestore/posinfo":     "github.com/peterargue/boxo/filestore/posinfo",
		"github.com/ipfs/go-libipfs/util":                  "github.com/peterargue/boxo/util",
		"github.com/ipfs/go-libipfs/datastore/dshelp":      "github.com/peterargue/boxo/datastore/dshelp",
		"github.com/ipfs/go-libipfs/verifcid":              "github.com/peterargue/boxo/verifcid",
		"github.com/ipfs/go-libipfs/exchange/offline":      "github.com/peterargue/boxo/exchange/offline",
		"github.com/ipfs/go-libipfs/routing":               "github.com/peterargue/boxo/routing",
		"github.com/ipfs/go-libipfs/exchange":              "github.com/peterargue/boxo/exchange",

		// Unmigrated things
		"github.com/ipfs/go-libipfs/blocks": "github.com/ipfs/go-block-format",
		"github.com/peterargue/boxo/blocks":       "github.com/ipfs/go-block-format",
	},
	Modules: []string{
		"github.com/ipfs/go-bitswap",
		"github.com/ipfs/go-ipfs-files",
		"github.com/ipfs/tar-utils",
		"gihtub.com/ipfs/go-block-format",
		"github.com/ipfs/interface-go-ipfs-core",
		"github.com/ipfs/go-unixfs",
		"github.com/ipfs/go-pinning-service-http-client",
		"github.com/ipfs/go-path",
		"github.com/ipfs/go-namesys",
		"github.com/ipfs/go-mfs",
		"github.com/ipfs/go-ipfs-provider",
		"github.com/ipfs/go-ipfs-pinner",
		"github.com/ipfs/go-ipfs-keystore",
		"github.com/ipfs/go-filestore",
		"github.com/ipfs/go-ipns",
		"github.com/ipfs/go-blockservice",
		"github.com/ipfs/go-ipfs-chunker",
		"github.com/ipfs/go-fetcher",
		"github.com/ipfs/go-ipfs-blockstore",
		"github.com/ipfs/go-ipfs-posinfo",
		"github.com/ipfs/go-ipfs-util",
		"github.com/ipfs/go-ipfs-ds-help",
		"github.com/ipfs/go-verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline",
		"github.com/ipfs/go-ipfs-routing",
		"github.com/ipfs/go-ipfs-exchange-interface",
		"github.com/ipfs/go-libipfs",
	},
}

func ReadConfig(r io.Reader) (Config, error) {
	var config Config
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("reading and decoding config: %w", err)
	}
	return config, nil
}
