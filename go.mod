module github.com/palcoin-project/palcwallet

go 1.16

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/protobuf v1.4.3
	github.com/jessevdk/go-flags v1.5.0
	github.com/jrick/logrotate v1.0.0
	github.com/lightninglabs/gozmq v0.0.0-20191113021534-d20a764486bf
	github.com/lightningnetwork/lnd/ticker v1.0.0
	github.com/palcoin-project/neutrino v0.12.11
	github.com/palcoin-project/palcd v0.22.0-beta
	github.com/palcoin-project/palcutil v1.0.3
	github.com/palcoin-project/palcutil/psbt v1.0.3-0.20210810200445-1144b2ef60c5
	github.com/palcoin-project/palcwallet/wallet/txauthor v1.0.2-0.20210811110650-52bb517d9666
	github.com/palcoin-project/palcwallet/wallet/txrules v1.0.1-0.20210811110650-52bb517d9666
	github.com/palcoin-project/palcwallet/wallet/txsizes v1.0.2-0.20210811110650-52bb517d9666
	github.com/palcoin-project/palcwallet/walletdb v1.4.1
	github.com/palcoin-project/palcwallet/wtxmgr v1.3.1-0.20210811110650-52bb517d9666
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	google.golang.org/grpc v1.39.1
)
