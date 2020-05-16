// Copyright 2016 The go-ethereum Authors
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

package les

import (
	"context"
	"math/big"

	"github.com/taiyuechain/taipublicchain/accounts"
	"github.com/taiyuechain/taipublicchain/common"
	"github.com/taiyuechain/taipublicchain/common/math"
	"github.com/taiyuechain/taipublicchain/core"
	"github.com/taiyuechain/taipublicchain/core/bloombits"
	"github.com/taiyuechain/taipublicchain/core/rawdb"
	"github.com/taiyuechain/taipublicchain/core/state"
	"github.com/taiyuechain/taipublicchain/core/types"
	"github.com/taiyuechain/taipublicchain/core/vm"
	"github.com/taiyuechain/taipublicchain/event"
	"github.com/taiyuechain/taipublicchain/light"
	"github.com/taiyuechain/taipublicchain/params"
	"github.com/taiyuechain/taipublicchain/rpc"
	"github.com/taiyuechain/taipublicchain/tai/downloader"
	"github.com/taiyuechain/taipublicchain/tai/gasprice"
	"github.com/taiyuechain/taipublicchain/taidb"
)

type LesApiBackend struct {
	tai *LightTai
	gpo *gasprice.Oracle
}

func (b *LesApiBackend) ChainConfig() *params.ChainConfig {
	return b.tai.chainConfig
}

func (b *LesApiBackend) CurrentBlock() *types.Block {
	return types.NewBlockWithHeader(b.tai.BlockChain().CurrentHeader())
}

func (b *LesApiBackend) CurrentSnailBlock() *types.SnailBlock {
	return nil
}

func (b *LesApiBackend) SetHead(number uint64) {
	b.tai.protocolManager.downloader.Cancel()
	b.tai.blockchain.SetHead(number)
}

func (b *LesApiBackend) SetSnailHead(number uint64) {
}

func (b *LesApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	if blockNr == rpc.LatestBlockNumber || blockNr == rpc.PendingBlockNumber {
		return b.tai.blockchain.CurrentHeader(), nil
	}

	return b.tai.blockchain.GetHeaderByNumberOdr(ctx, uint64(blockNr))
}
func (b *LesApiBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.tai.blockchain.GetHeaderByHash(hash), nil
}

// TODO: fixed lightchain func.
func (b *LesApiBackend) SnailHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.SnailHeader, error) {
	return nil, nil
}

func (b *LesApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, err
	}
	return b.GetBlock(ctx, header.Hash())
}

// TODO: fixed lightchain func.
func (b *LesApiBackend) SnailBlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.SnailBlock, error) {
	return nil, nil
}

func (b *LesApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	return light.NewState(ctx, header, b.tai.odr), header, nil
}

func (b *LesApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.tai.blockchain.GetBlockByHash(ctx, blockHash)
}

// TODO: fixed lightchain func.
func (b *LesApiBackend) GetFruit(ctx context.Context, fastblockHash common.Hash) (*types.SnailBlock, error) {
	return nil, nil
}

// TODO: fixed lightchain func.
func (b *LesApiBackend) GetSnailBlock(ctx context.Context, blockHash common.Hash) (*types.SnailBlock, error) {
	return nil, nil
}

func (b *LesApiBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.tai.chainDb, hash); number != nil {
		return light.GetBlockReceipts(ctx, b.tai.odr, hash, *number)
	}
	return nil, nil
}

func (b *LesApiBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	if number := rawdb.ReadHeaderNumber(b.tai.chainDb, hash); number != nil {
		return light.GetBlockLogs(ctx, b.tai.odr, hash, *number)
	}
	return nil, nil
}

func (b *LesApiBackend) GetTd(hash common.Hash) *big.Int {
	return b.tai.blockchain.GetTdByHash(hash)
}

func (b *LesApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	context := core.NewEVMContext(msg, header, b.tai.blockchain, nil, nil)
	return vm.NewEVM(context, state, b.tai.chainConfig, vmCfg), state.Error, nil
}

func (b *LesApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.tai.txPool.Add(ctx, signedTx)
}

func (b *LesApiBackend) RemoveTx(txHash common.Hash) {
	b.tai.txPool.RemoveTx(txHash)
}

func (b *LesApiBackend) GetPoolTransactions() (types.Transactions, error) {
	return b.tai.txPool.GetTransactions()
}

func (b *LesApiBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.tai.txPool.GetTransaction(txHash)
}

func (b *LesApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.tai.txPool.GetNonce(ctx, addr)
}

func (b *LesApiBackend) Stats() (pending int, queued int) {
	return b.tai.txPool.Stats(), 0
}

func (b *LesApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.tai.txPool.Content()
}

func (b *LesApiBackend) SubscribeNewTxsEvent(ch chan<- types.NewTxsEvent) event.Subscription {
	return b.tai.txPool.SubscribeNewTxsEvent(ch)
}

func (b *LesApiBackend) SubscribeChainEvent(ch chan<- types.FastChainEvent) event.Subscription {
	return b.tai.blockchain.SubscribeChainEvent(ch)
}

func (b *LesApiBackend) SubscribeChainHeadEvent(ch chan<- types.FastChainHeadEvent) event.Subscription {
	return b.tai.blockchain.SubscribeChainHeadEvent(ch)
}

func (b *LesApiBackend) SubscribeChainSideEvent(ch chan<- types.FastChainSideEvent) event.Subscription {
	return b.tai.blockchain.SubscribeChainSideEvent(ch)
}

func (b *LesApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.tai.blockchain.SubscribeLogsEvent(ch)
}

func (b *LesApiBackend) SubscribeRemovedLogsEvent(ch chan<- types.RemovedLogsEvent) event.Subscription {
	return b.tai.blockchain.SubscribeRemovedLogsEvent(ch)
}

func (b *LesApiBackend) GetReward(number int64) *types.BlockReward {
	//if number < 0 {
	//	return b.tai.blockchain.CurrentReward()
	//}

	//return b.tai.blockchain.GetFastHeightBySnailHeight(uint64(number))
	return nil
}

func (b *LesApiBackend) GetCommittee(number rpc.BlockNumber) (map[string]interface{}, error) {
	return nil, nil
}

func (b *LesApiBackend) GetCurrentCommitteeNumber() *big.Int {
	return nil
}

func (b *LesApiBackend) GetSnailRewardContent(number rpc.BlockNumber) *types.SnailRewardContenet {
	return nil
}

func (b *LesApiBackend) SnailPoolContent() []*types.SnailBlock {
	return nil
}

func (b *LesApiBackend) SnailPoolInspect() []*types.SnailBlock {
	return nil
}

func (b *LesApiBackend) SnailPoolStats() (pending int, unVerified int) {
	return 0, 0
}

func (b *LesApiBackend) Downloader() *downloader.Downloader {
	return b.tai.Downloader()
}

func (b *LesApiBackend) ProtocolVersion() int {
	return b.tai.LesVersion() + 10000
}

func (b *LesApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *LesApiBackend) ChainDb() taidb.Database {
	return b.tai.chainDb
}

func (b *LesApiBackend) EventMux() *event.TypeMux {
	return b.tai.eventMux
}

func (b *LesApiBackend) AccountManager() *accounts.Manager {
	return b.tai.accountManager
}

func (b *LesApiBackend) BloomStatus() (uint64, uint64) {
	if b.tai.bloomIndexer == nil {
		return 0, 0
	}
	sections, _, _ := b.tai.bloomIndexer.Sections()
	return light.BloomTrieFrequency, sections
}

func (b *LesApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.tai.bloomRequests)
	}
}
