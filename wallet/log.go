// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"github.com/palcoin-project/palclog"
	"github.com/palcoin-project/palcwallet/waddrmgr"
	"github.com/palcoin-project/palcwallet/walletdb/migration"
	"github.com/palcoin-project/palcwallet/wtxmgr"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log palclog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	UseLogger(palclog.Disabled)
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using palclog.
func UseLogger(logger palclog.Logger) {
	log = logger

	migration.UseLogger(logger)
	waddrmgr.UseLogger(logger)
	wtxmgr.UseLogger(logger)
}

// pickNoun returns the singular or plural form of a noun depending
// on the count n.
func pickNoun(n int, singular, plural string) string {
	if n == 1 {
		return singular
	}
	return plural
}
