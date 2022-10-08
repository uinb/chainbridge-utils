module github.com/uinb/chainbridge-utils

go 1.18

// NOTE - this is a specific branch - https://github.com/centrifuge/go-substrate-rpc-client/tree/remove-claims-event,
// that does not have the `Claims_Claimed` event since it is colliding with the one that we have in the claims pallet
// of Centrifuge chain.
require github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.8-0.20220930212708-8350785ebd6f

require (
	github.com/ChainSafe/log15 v1.0.0
	github.com/btcsuite/btcd v0.22.1
	github.com/ethereum/go-ethereum v1.10.17
	github.com/prometheus/client_golang v1.4.1
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56
)

require (
	github.com/ChainSafe/go-schnorrkel v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.1.2 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/decred/base58 v1.0.3 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.9.1 // indirect
	github.com/prometheus/procfs v0.0.10 // indirect
	github.com/vedhavyas/go-subkey v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	google.golang.org/protobuf v1.23.0 // indirect
)
