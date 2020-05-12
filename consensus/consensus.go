// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package consensus implements different Ethereum consensus engines.
package consensus

import (
	"math/big"

	"github.com/taiyuechain/taipublicchain/common"
	"github.com/taiyuechain/taipublicchain/core/state"
	"github.com/taiyuechain/taipublicchain/core/types"
	"github.com/taiyuechain/taipublicchain/params"
	"github.com/taiyuechain/taipublicchain/rpc"
)

// ChainReader defines a small collection of methods needed to access the local
// blockchain during header and/or uncle verification.
type ChainReader interface {
	// Config retrieves the blockchain's chain configuration.
	Config() *params.ChainConfig

	// CurrentHeader retrieves the current header from the local chain.
	CurrentHeader() *types.Header

	// GetHeader retrieves a block header from the database by hash and number.
	GetHeader(hash common.Hash, number uint64) *types.Header

	// GetHeaderByNumber retrieves a block header from the database by number.
	GetHeaderByNumber(number uint64) *types.Header

	// GetHeaderByHash retrieves a block header from the database by its hash.
	GetHeaderByHash(hash common.Hash) *types.Header

	// GetBlock retrieves a block from the database by hash and number.
	GetBlock(hash common.Hash, number uint64) *types.Block
}

// ChainSnailReader defines a small collection of methods needed to access the local
// block chain during header and/or uncle verification.
// Temporary interface for snail
type SnailChainReader interface {
	// Config retrieves the blockchain's chain configuration.
	Config() *params.ChainConfig

	// CurrentHeader retrieves the current header from the local chain.
	CurrentHeader() *types.SnailHeader

	// GetHeader retrieves a block header from the database by hash and number.
	GetHeader(hash common.Hash, number uint64) *types.SnailHeader

	// GetHeaderByNumber retrieves a block header from the database by number.
	GetHeaderByNumber(number uint64) *types.SnailHeader

	// GetHeaderByHash retrieves a block header from the database by its hash.
	GetHeaderByHash(hash common.Hash) *types.SnailHeader

	// GetBlock retrieves a block from the database by hash and number.
	GetBlock(hash common.Hash, number uint64) *types.SnailBlock
}

// Engine is an algorithm agnostic consensus engine.
type Engine interface {
	SetElection(e CommitteeElection)

	GetElection() CommitteeElection

	SetSnailChainReader(scr SnailChainReader)

	// Author retrieves the Ethereum address of the account that minted the given
	// block, which may be different from the header's coinbase if a consensus
	// engine is based on signatures.
	Author(header *types.Header) (common.Address, error)
	AuthorSnail(header *types.SnailHeader) (common.Address, error)

	// VerifyHeader checks whether a header conforms to the consensus rules of a
	// given engine. Verifying the seal may be done optionally here, or explicitly
	// via the VerifySeal method.
	VerifyHeader(chain ChainReader, header *types.Header) error
	VerifySnailHeader(chain SnailChainReader, fastchain ChainReader, header *types.SnailHeader, seal bool, isFruit bool) error

	// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
	// concurrently. The method returns a quit channel to abort the operations and
	// a results channel to retrieve the async verifications (the order is that of
	// the input slice).
	VerifyHeaders(chain ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error)

	// VerifySnailHeaders is similar to VerifySnailHeader, but verifies a batch of headers concurrently.
	// VerifySnailHeaders only verifies snail header rather than fruit header.
	// The method returns a quit channel to abort the operations and
	// a results channel to retrieve the async verifications (the order is that of
	// the input slice).
	VerifySnailHeaders(chain SnailChainReader, headers []*types.SnailHeader, seals []bool) (chan<- struct{}, <-chan error)

	// VerifySeal checks whether the crypto seal on a header is valid according to
	// the consensus rules of the given engine.
	VerifySnailSeal(chain SnailChainReader, header *types.SnailHeader, isFruit bool) error

	VerifyFreshness(chain SnailChainReader, fruit *types.SnailHeader, headerNumber *big.Int, canonical bool) error

	VerifySigns(fastnumber *big.Int, fastHash common.Hash, signs []*types.PbftSign) error

	VerifySwitchInfo(fastnumber *big.Int, info []*types.CommitteeMember) error

	// Prepare initializes the consensus fields of a block header according to the
	// rules of a particular engine. The changes are executed inline.
	Prepare(chain ChainReader, header *types.Header) error
	PrepareSnail(chain ChainReader, snailchain SnailChainReader, header *types.SnailHeader) error
	PrepareSnailWithParent(chain ChainReader, snailchain SnailChainReader, header *types.SnailHeader, parents []*types.SnailHeader) error

	// Finalize runs any post-transaction state modifications (e.g. block rewards)
	// and assembles the final block.
	// Note: The block header and state database might be updated to reflect any
	// consensus rules that happen at finalization (e.g. block rewards).
	Finalize(chain ChainReader, header *types.Header, state *state.StateDB,
		txs []*types.Transaction, receipts []*types.Receipt, feeAmount *big.Int) (*types.Block, error)
	FinalizeSnail(chain SnailChainReader, header *types.SnailHeader,
		uncles []*types.SnailHeader, fruits []*types.SnailBlock, signs []*types.PbftSign) (*types.SnailBlock, error)

	FinalizeCommittee(block *types.Block) error

	// Seal generates a new block for the given input block with the local miner's
	Seal(chain SnailChainReader, block *types.SnailBlock, stop <-chan struct{}) (*types.SnailBlock, error)

	// ConSeal generates a new block for the given input block with the local miner's
	// seal place on top.
	ConSeal(chain SnailChainReader, block *types.SnailBlock, stop <-chan struct{}, send chan *types.SnailBlock)

	CalcSnailDifficulty(chain SnailChainReader, time uint64, parents []*types.SnailHeader) *big.Int

	GetDifficulty(header *types.SnailHeader, isFruit bool) (*big.Int, *big.Int)

	// APIs returns the RPC APIs this consensus engine provides.
	APIs(chain ChainReader) []rpc.API

	DataSetHash(epoch uint64) string

	GetRewardContentBySnailNumber(sBlock *types.SnailBlock) *types.SnailRewardContenet
}

//Election module implementation committee interface
type CommitteeElection interface {
	// VerifySigns verify the fast chain committee signatures in batches
	VerifySigns(pvs []*types.PbftSign) ([]*types.CommitteeMember, []error)

	// VerifySwitchInfo verify committee members and it's state
	VerifySwitchInfo(fastnumber *big.Int, info []*types.CommitteeMember) error

	FinalizeCommittee(block *types.Block) error

	//Get a list of committee members
	//GetCommittee(FastNumber *big.Int, FastHash common.Hash) (*big.Int, []*types.CommitteeMember)
	GetCommittee(fastNumber *big.Int) []*types.CommitteeMember

	GenerateFakeSigns(fb *types.Block) ([]*types.PbftSign, error)
}

// PoW is a consensus engine based on proof-of-work.
type PoW interface {
	Engine

	// Hashrate returns the current mining hashrate of a PoW consensus engine.
	Hashrate() float64
}
