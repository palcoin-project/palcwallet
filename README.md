palcwallet
=========
palcwallet is a daemon handling palcoin wallet functionality for a
single user.  It acts as both an RPC client to palcd and an RPC server
for wallet clients.

Public and private keys are derived using the hierarchical
deterministic format described by
[BIP0032](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki).
Unencrypted private keys are not supported and are never written to
disk.  palcwallet uses the
`m/44'/<coin type>'/<account>'/<branch>/<address index>`
HD path for all derived addresses, as described by
[BIP0044](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki).

Due to the sensitive nature of public data in a BIP0032 wallet,
palcwallet provides the option of encrypting not just private keys, but
public data as well.  This is intended to thwart privacy risks where a
wallet file is compromised without exposing all current and future
addresses (public keys) managed by the wallet. While access to this
information would not allow an attacker to spend or steal coins, it
does mean they could track all transactions involving your addresses
and therefore know your exact balance.  In a future release, public data
encryption will extend to transactions as well.

palcwallet is not an SPV client and requires connecting to a local or
remote palcd instance for asynchronous blockchain queries and
notifications over websockets.  

## Requirements

[Go](http://golang.org) 1.12 or newer.

### Windows/Linux/BSD/POSIX - Build from source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain palcwallet, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/palcoin-project/palcwallet
$ GO111MODULE=on go install -v . ./cmd/...
```

- palcwallet (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update btcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/palcoin-project/palcwallet
$ git pull
$ GO111MODULE=on go install -v . ./cmd/...
```

## Getting Started

The following instructions detail how to get started with palcwallet connecting
to a localhost btcd.  Commands should be run in `cmd.exe` or PowerShell on
Windows, or any terminal emulator on *nix.

- Run the following command to start palcd:

```
palcd -u rpcuser -P rpcpass
```

- Run the following command to create a wallet:

```
palcwallet -u rpcuser -P rpcpass --create
```

- Run the following command to start palcwallet:

```
palcwallet -u rpcuser -P rpcpass
```

If everything appears to be working, it is recommended at this point to
copy the sample palcd and palcwallet configurations and update with your
RPC username and password.

PowerShell (Installed from MSI):
```
PS> cp "$env:ProgramFiles\palcoin-project\Palcd\sample-palcd.conf" $env:LOCALAPPDATA\Palcd\palcd.conf
PS> cp "$env:ProgramFiles\palcoin-project\Palcwallet\sample-palcwallet.conf" $env:LOCALAPPDATA\Palcwallet\palcwallet.conf
PS> $editor $env:LOCALAPPDATA\Palcd\palcd.conf
PS> $editor $env:LOCALAPPDATA\Palcwallet\palcwallet.conf
```

PowerShell (Installed from source):
```
PS> cp $env:GOPATH\src\github.com\palcoin-project\palcd\sample-palcd.conf $env:LOCALAPPDATA\Palcd\palcd.conf
PS> cp $env:GOPATH\src\github.com\palcoin-project\palcwallet\sample-palcwallet.conf $env:LOCALAPPDATA\Palcwallet\palcwallet.conf
PS> $editor $env:LOCALAPPDATA\Palcd\palcd.conf
PS> $editor $env:LOCALAPPDATA\Palcwallet\palcwallet.conf
```

Linux/BSD/POSIX (Installed from source):
```bash
$ cp $GOPATH/src/github.com/palcoin-project/palcd/sample-palcd.conf ~/.palcd/palcd.conf
$ cp $GOPATH/src/github.com/palcoin-project/palcwallet/sample-palcwallet.conf ~/.palcwallet/palcwallet.conf
$ $EDITOR ~/.palcd/palcd.conf
$ $EDITOR ~/.palcwallet/palcwallet.conf
```


palcwallet is licensed under the liberal ISC License.
