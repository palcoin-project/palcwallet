// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import (
	"github.com/palcoin-project/palcd/chaincfg"
	"github.com/palcoin-project/palcd/wire"
)

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	RPCClientPort string
	RPCServerPort string
}

// MainNetParams contains parameters specific running palcwallet and
// btcd on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:        &chaincfg.MainNetParams,
	RPCClientPort: "1967", //8334
	RPCServerPort: "1987", //8332
}

// TestNet3Params contains parameters specific running palcwallet and
// btcd on the test network (version 3) (wire.TestNet3).
var TestNet3Params = Params{
	Params:        &chaincfg.TestNet3Params,
	RPCClientPort: "11967", //18334
	RPCServerPort: "11987", //18332
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:        &chaincfg.SimNetParams,
	RPCClientPort: "11918", //18556
	RPCServerPort: "11916", //18554
}

// SigNetParams contains parameters specific to the signet test network
// (wire.SigNet).
var SigNetParams = Params{
	Params:        &chaincfg.SigNetParams,
	RPCClientPort: "31967", //38334
	RPCServerPort: "31987", //38332
}

// SigNetWire is a helper function that either returns the given chain
// parameter's net value if the parameter represents a signet network or 0 if
// it's not. This is necessary because there can be custom signet networks that
// have a different net value.
func SigNetWire(params *chaincfg.Params) wire.BitcoinNet {
	if params.Name == chaincfg.SigNetParams.Name {
		return params.Net
	}

	return 0
}
