// Copyright 2015 The go-ethereum Authors
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

package tai

import (
	"github.com/taiyuechain/taipublicchain/metrics"
	"github.com/taiyuechain/taipublicchain/p2p"
)

var (
	propTxnInPacketsMeter     = metrics.NewRegisteredMeter("etai/prop/txns/in/packets", nil)
	propTxnInTxsMeter         = metrics.NewRegisteredMeter("etai/prop/txns/in/txs", nil)
	propTxnInTrafficMeter     = metrics.NewRegisteredMeter("etai/prop/txns/in/traffic", nil)
	propTxnOutPacketsMeter    = metrics.NewRegisteredMeter("etai/prop/txns/out/packets", nil)
	propTxnOutTrafficMeter    = metrics.NewRegisteredMeter("etai/prop/txns/out/traffic", nil)
	propFtnInPacketsMeter     = metrics.NewRegisteredMeter("etai/prop/ftns/in/packets", nil)
	propFtnInTrafficMeter     = metrics.NewRegisteredMeter("etai/prop/ftns/in/traffic", nil)
	propFtnOutPacketsMeter    = metrics.NewRegisteredMeter("etai/prop/ftns/out/packets", nil)
	propFtnOutTrafficMeter    = metrics.NewRegisteredMeter("etai/prop/ftns/out/traffic", nil)
	propFHashInPacketsMeter   = metrics.NewRegisteredMeter("etai/prop/fhashes/in/packets", nil)
	propFHashInTrafficMeter   = metrics.NewRegisteredMeter("etai/prop/fhashes/in/traffic", nil)
	propFHashOutPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/fhashes/out/packets", nil)
	propFHashOutTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/fhashes/out/traffic", nil)
	propSHashInPacketsMeter   = metrics.NewRegisteredMeter("etai/prop/shashes/in/packets", nil)
	propSHashInTrafficMeter   = metrics.NewRegisteredMeter("etai/prop/shashes/in/traffic", nil)
	propSHashOutPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/shashes/out/packets", nil)
	propSHashOutTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/shashes/out/traffic", nil)
	propFBlockInPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/fblocks/in/packets", nil)
	propFBlockInTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/fblocks/in/traffic", nil)
	propFBlockOutPacketsMeter = metrics.NewRegisteredMeter("etai/prop/fblocks/out/packets", nil)
	propFBlockOutTrafficMeter = metrics.NewRegisteredMeter("etai/prop/fblocks/out/traffic", nil)
	propSBlockInPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/sblocks/in/packets", nil)
	propSBlockInTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/sblocks/in/traffic", nil)
	propSBlockOutPacketsMeter = metrics.NewRegisteredMeter("etai/prop/sblocks/out/packets", nil)
	propSBlockOutTrafficMeter = metrics.NewRegisteredMeter("etai/prop/sblocks/out/traffic", nil)

	propNodeInfoInPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/nodeinfo/in/packets", nil)
	propNodeInfoInTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/nodeinfo/in/traffic", nil)
	propNodeInfoOutPacketsMeter = metrics.NewRegisteredMeter("etai/prop/nodeinfo/out/packets", nil)
	propNodeInfoOutTrafficMeter = metrics.NewRegisteredMeter("etai/prop/nodeinfo/out/traffic", nil)

	propNodeInfoHashInPacketsMeter  = metrics.NewRegisteredMeter("etai/prop/nodeinfohash/in/packets", nil)
	propNodeInfoHashInTrafficMeter  = metrics.NewRegisteredMeter("etai/prop/nodeinfohash/in/traffic", nil)
	propNodeInfoHashOutPacketsMeter = metrics.NewRegisteredMeter("etai/prop/nodeinfohash/out/packets", nil)
	propNodeInfoHashOutTrafficMeter = metrics.NewRegisteredMeter("etai/prop/nodeinfohash/out/traffic", nil)

	reqFHeaderInPacketsMeter  = metrics.NewRegisteredMeter("etai/req/headers/in/packets", nil)
	reqFHeaderInTrafficMeter  = metrics.NewRegisteredMeter("etai/req/headers/in/traffic", nil)
	reqFHeaderOutPacketsMeter = metrics.NewRegisteredMeter("etai/req/headers/out/packets", nil)
	reqFHeaderOutTrafficMeter = metrics.NewRegisteredMeter("etai/req/headers/out/traffic", nil)
	reqSHeaderInPacketsMeter  = metrics.NewRegisteredMeter("etai/req/sheaders/in/packets", nil)
	reqSHeaderInTrafficMeter  = metrics.NewRegisteredMeter("etai/req/sheaders/in/traffic", nil)
	reqSHeaderOutPacketsMeter = metrics.NewRegisteredMeter("etai/req/sheaders/out/packets", nil)
	reqSHeaderOutTrafficMeter = metrics.NewRegisteredMeter("etai/req/sheaders/out/traffic", nil)

	reqFBodyInPacketsMeter  = metrics.NewRegisteredMeter("etai/req/fbodies/in/packets", nil)
	reqFBodyInTrafficMeter  = metrics.NewRegisteredMeter("etai/req/fbodies/in/traffic", nil)
	reqFBodyOutPacketsMeter = metrics.NewRegisteredMeter("etai/req/fbodies/out/packets", nil)
	reqFBodyOutTrafficMeter = metrics.NewRegisteredMeter("etai/req/fbodies/out/traffic", nil)
	reqSBodyInPacketsMeter  = metrics.NewRegisteredMeter("etai/req/sbodies/in/packets", nil)
	reqSBodyInTrafficMeter  = metrics.NewRegisteredMeter("etai/req/sbodies/in/traffic", nil)
	reqSBodyOutPacketsMeter = metrics.NewRegisteredMeter("etai/req/sbodies/out/packets", nil)
	reqSBodyOutTrafficMeter = metrics.NewRegisteredMeter("etai/req/sbodies/out/traffic", nil)

	reqStateInPacketsMeter    = metrics.NewRegisteredMeter("etai/req/states/in/packets", nil)
	reqStateInTrafficMeter    = metrics.NewRegisteredMeter("etai/req/states/in/traffic", nil)
	reqStateOutPacketsMeter   = metrics.NewRegisteredMeter("etai/req/states/out/packets", nil)
	reqStateOutTrafficMeter   = metrics.NewRegisteredMeter("etai/req/states/out/traffic", nil)
	reqReceiptInPacketsMeter  = metrics.NewRegisteredMeter("etai/req/receipts/in/packets", nil)
	reqReceiptInTrafficMeter  = metrics.NewRegisteredMeter("etai/req/receipts/in/traffic", nil)
	reqReceiptOutPacketsMeter = metrics.NewRegisteredMeter("etai/req/receipts/out/packets", nil)
	reqReceiptOutTrafficMeter = metrics.NewRegisteredMeter("etai/req/receipts/out/traffic", nil)

	getHeadInPacketsMeter  = metrics.NewRegisteredMeter("etai/get/head/in/packets", nil)
	getHeadInTrafficMeter  = metrics.NewRegisteredMeter("etai/get/head/in/traffic", nil)
	getHeadOutPacketsMeter = metrics.NewRegisteredMeter("etai/get/head/out/packets", nil)
	getHeadOutTrafficMeter = metrics.NewRegisteredMeter("etai/get/head/out/traffic", nil)
	getBodyInPacketsMeter  = metrics.NewRegisteredMeter("etai/get/bodies/in/packets", nil)
	getBodyInTrafficMeter  = metrics.NewRegisteredMeter("etai/get/bodies/in/traffic", nil)
	getBodyOutPacketsMeter = metrics.NewRegisteredMeter("etai/get/bodies/out/packets", nil)
	getBodyOutTrafficMeter = metrics.NewRegisteredMeter("etai/get/bodies/out/traffic", nil)

	getNodeInfoInPacketsMeter  = metrics.NewRegisteredMeter("etai/get/nodeinfo/in/packets", nil)
	getNodeInfoInTrafficMeter  = metrics.NewRegisteredMeter("etai/get/nodeinfo/in/traffic", nil)
	getNodeInfoOutPacketsMeter = metrics.NewRegisteredMeter("etai/get/nodeinfo/out/packets", nil)
	getNodeInfoOutTrafficMeter = metrics.NewRegisteredMeter("etai/get/nodeinfo/out/traffic", nil)

	miscInPacketsMeter  = metrics.NewRegisteredMeter("etai/misc/in/packets", nil)
	miscInTrafficMeter  = metrics.NewRegisteredMeter("etai/misc/in/traffic", nil)
	miscOutPacketsMeter = metrics.NewRegisteredMeter("etai/misc/out/packets", nil)
	miscOutTrafficMeter = metrics.NewRegisteredMeter("etai/misc/out/traffic", nil)
)

// meteredMsgReadWriter is a wrapper around a p2p.MsgReadWriter, capable of
// accumulating the above defined metrics based on the data stream contents.
type meteredMsgReadWriter struct {
	p2p.MsgReadWriter     // Wrapped message stream to meter
	version           int // Protocol version to select correct meters
}

// newMeteredMsgWriter wraps a p2p MsgReadWriter with metering support. If the
// metrics system is disabled, this function returns the original object.
func newMeteredMsgWriter(rw p2p.MsgReadWriter) p2p.MsgReadWriter {
	if !metrics.Enabled {
		return rw
	}
	return &meteredMsgReadWriter{MsgReadWriter: rw}
}

// Init sets the protocol version used by the stream to know which meters to
// increment in case of overlapping message ids between protocol versions.
func (rw *meteredMsgReadWriter) Init(version int) {
	rw.version = version
}

func (rw *meteredMsgReadWriter) ReadMsg() (p2p.Msg, error) {
	// Read the message and short circuit in case of an error
	msg, err := rw.MsgReadWriter.ReadMsg()
	if err != nil {
		return msg, err
	}
	// Account for the data traffic
	packets, traffic := miscInPacketsMeter, miscInTrafficMeter
	switch {
	case msg.Code == FastBlockHeadersMsg:
		packets, traffic = reqFHeaderInPacketsMeter, reqFHeaderInTrafficMeter
	case msg.Code == SnailBlockHeadersMsg:
		packets, traffic = reqSHeaderInPacketsMeter, reqSHeaderInTrafficMeter
	case msg.Code == FastBlockBodiesMsg:
		packets, traffic = reqFBodyInPacketsMeter, reqFBodyInTrafficMeter
	case msg.Code == SnailBlockBodiesMsg:
		packets, traffic = reqSBodyInPacketsMeter, reqSBodyInTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateInPacketsMeter, reqStateInTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptInPacketsMeter, reqReceiptInTrafficMeter

	case msg.Code == NewFastBlockHashesMsg:
		packets, traffic = propFHashInPacketsMeter, propFHashInTrafficMeter
	case msg.Code == NewSnailBlockHashesMsg:
		packets, traffic = propSHashInPacketsMeter, propSHashInTrafficMeter
	case msg.Code == NewFastBlockMsg:
		packets, traffic = propFBlockInPacketsMeter, propFBlockInTrafficMeter
	case msg.Code == NewSnailBlockMsg:
		packets, traffic = propSBlockInPacketsMeter, propSBlockInTrafficMeter
	case msg.Code == TxMsg:
		packets, traffic = propTxnInPacketsMeter, propTxnInTrafficMeter
	case msg.Code == NewFruitMsg:
		packets, traffic = propFtnInPacketsMeter, propFtnInTrafficMeter
	case msg.Code == TbftNodeInfoMsg:
		packets, traffic = propNodeInfoInPacketsMeter, propNodeInfoInTrafficMeter
	case msg.Code == TbftNodeInfoHashMsg:
		packets, traffic = propNodeInfoHashInPacketsMeter, propNodeInfoHashInTrafficMeter
	case msg.Code == GetTbftNodeInfoMsg:
		packets, traffic = getNodeInfoInPacketsMeter, getNodeInfoInTrafficMeter
	case msg.Code == GetFastBlockHeadersMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadInTrafficMeter
	case msg.Code == GetFastBlockBodiesMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadInTrafficMeter
	case msg.Code == GetSnailBlockHeadersMsg:
		packets, traffic = getBodyInPacketsMeter, getBodyInTrafficMeter
	case msg.Code == GetSnailBlockBodiesMsg:
		packets, traffic = getBodyInPacketsMeter, getBodyInTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	switch {
	case msg.Code == FastBlockHeadersMsg:
		packets, traffic = reqFHeaderOutPacketsMeter, reqFHeaderOutTrafficMeter
	case msg.Code == SnailBlockHeadersMsg:
		packets, traffic = reqSHeaderOutPacketsMeter, reqSHeaderOutTrafficMeter
	case msg.Code == FastBlockBodiesMsg:
		packets, traffic = reqFBodyOutPacketsMeter, reqFBodyOutTrafficMeter
	case msg.Code == SnailBlockBodiesMsg:
		packets, traffic = reqSBodyOutPacketsMeter, reqSBodyOutTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateOutPacketsMeter, reqStateOutTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptOutPacketsMeter, reqReceiptOutTrafficMeter

	case msg.Code == NewFastBlockHashesMsg:
		packets, traffic = propFHashOutPacketsMeter, propFHashOutTrafficMeter
	case msg.Code == NewSnailBlockHashesMsg:
		packets, traffic = propSHashOutPacketsMeter, propSHashOutTrafficMeter
	case msg.Code == NewFastBlockMsg:
		packets, traffic = propFBlockOutPacketsMeter, propFBlockOutTrafficMeter
	case msg.Code == NewSnailBlockMsg:
		packets, traffic = propSBlockOutPacketsMeter, propSBlockOutTrafficMeter
	case msg.Code == TxMsg:
		packets, traffic = propTxnOutPacketsMeter, propTxnOutTrafficMeter
	case msg.Code == NewFruitMsg:
		packets, traffic = propFtnOutPacketsMeter, propFtnOutTrafficMeter
	case msg.Code == TbftNodeInfoMsg:
		packets, traffic = propNodeInfoOutPacketsMeter, propNodeInfoOutTrafficMeter
	case msg.Code == TbftNodeInfoHashMsg:
		packets, traffic = propNodeInfoHashOutPacketsMeter, propNodeInfoHashOutTrafficMeter
	case msg.Code == GetTbftNodeInfoMsg:
		packets, traffic = getNodeInfoOutPacketsMeter, getNodeInfoOutTrafficMeter
	case msg.Code == GetFastBlockHeadersMsg:
		packets, traffic = getHeadOutPacketsMeter, getHeadOutTrafficMeter
	case msg.Code == GetFastBlockBodiesMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadOutTrafficMeter
	case msg.Code == GetSnailBlockHeadersMsg:
		packets, traffic = getBodyOutPacketsMeter, getBodyOutTrafficMeter
	case msg.Code == GetSnailBlockBodiesMsg:
		packets, traffic = getBodyOutPacketsMeter, getBodyOutTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
