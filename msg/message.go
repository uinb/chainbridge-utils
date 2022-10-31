// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type ChainId uint8
type TransferType string
type ResourceId [32]byte
type TxHash [32]byte
func (r ResourceId) Hex() string {
	return fmt.Sprintf("%x", r)
}

type Nonce uint64

func (n Nonce) Big() *big.Int {
	return big.NewInt(int64(n))
}

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"
var GenericTransfer TransferType = "GenericTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source       ChainId      // Source where message was initiated
	Destination  ChainId      // Destination chain of message
	Type         TransferType // type of bridge transfer
	DepositNonce Nonce        // Nonce for the deposit
	Depositer    common.Address
	ResourceId   ResourceId
	TxHash       TxHash
	Payload      []interface{} // data associated with event sequence
}

func NewFungibleTransfer(source,dest ChainId, depositer common.Address, nonce Nonce, amount *big.Int, resourceId ResourceId, txHash TxHash, recipient []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         FungibleTransfer,
		DepositNonce: nonce,
		Depositer:    depositer,
		ResourceId:   resourceId,
		TxHash:       txHash,
		Payload: []interface{}{
			amount.Bytes(),
			recipient,
		},
	}
}

func NewNonFungibleTransfer(source, dest ChainId, depositer common.Address, nonce Nonce, resourceId ResourceId, tokenId *big.Int, txHash TxHash, recipient, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         NonFungibleTransfer,
		DepositNonce: nonce,
		Depositer:    depositer,
		ResourceId:   resourceId,
		TxHash:       txHash,
		Payload: []interface{}{
			tokenId.Bytes(),
			recipient,
			metadata,
		},
	}
}

func NewGenericTransfer(source, dest ChainId, depositer common.Address, nonce Nonce, resourceId ResourceId, txHash TxHash, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         GenericTransfer,
		DepositNonce: nonce,
		Depositer:    depositer,
		ResourceId:   resourceId,
		TxHash:       txHash,
		Payload: []interface{}{
			metadata,
		},
	}
}

func ResourceIdFromSlice(in []byte) ResourceId {
	var res ResourceId
	copy(res[:], in)
	return res
}
