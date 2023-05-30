package cylinder

import "time"

// Config data structure for Cylinder process.
type Config struct {
	ChainID          string        `mapstructure:"chain-id"`          // ChainID of the target chain
	NodeURI          string        `mapstructure:"node"`              // Remote RPC URI of BandChain node to connect to
	Granter          string        `mapstructure:"granter"`           // The granter address
	GasPrices        string        `mapstructure:"gas-prices"`        // Gas prices of the transaction
	LogLevel         string        `mapstructure:"log-level"`         // Log level of the logger
	BroadcastTimeout time.Duration `mapstructure:"broadcast-timeout"` // The time that cylinder will wait for tx commit
	RPCPollInterval  time.Duration `mapstructure:"rpc-poll-interval"` // The duration of rpc poll interval
	MaxTry           uint64        `mapstructure:"max-try"`           // The maximum number of tries to submit a report transaction
	MinDE            uint64        `mapstructure:"min-DE"`            // The minimum number of DE
}
